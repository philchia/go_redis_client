package redis

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func parseResults(s string, n int) Result {
	result := new(redisResult)
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)

	if n == 1 {
		return parseResult(scanner)
	}
	var lines = make([]Result, n)
	for i := range lines {
		lines[i] = parseResult(scanner)
	}
	result.Res = lines
	return result
}

func parseResult(scanner *bufio.Scanner) Result {
	if scanner.Scan() {
		str := scanner.Text()
		switch str[:1] {
		case "+":
			return parseResponse(scanner)
		case "-":
			return parseError(scanner)
		case "*":
			return parseArr(scanner)
		case "$":
			return parseStr(scanner)
		case ":":
			return parseInt(scanner)
		default:
			return &redisResult{Res: errors.New("Unknow response")}
		}
	}
	return nil
}

func parseResponse(scanner *bufio.Scanner) Result {
	return &redisResult{Res: scanner.Text()[1:]}
}

func parseError(scanner *bufio.Scanner) Result {
	return &redisResult{Res: errors.New(scanner.Text()[1:])}
}

func parseStr(scanner *bufio.Scanner) Result {
	str := scanner.Text()
	length := parseLength(str)
	if length < 0 {
		return nil
	}
	if scanner.Scan() {
		nxStr := scanner.Text()
		return &redisResult{Res: nxStr[:length]}
	}
	return nil
}

func parseArr(scanner *bufio.Scanner) Result {
	str := scanner.Text()
	length := parseLength(str)
	if length < 0 {
		return nil
	}
	res := make([]Result, length)
	for i := range res {
		if scanner.Scan() {
			res[i] = parseResult(scanner)
		} else {
			break
		}
	}
	return &redisResult{Res: res}
}

func parseInt(scanner *bufio.Scanner) Result {
	str := scanner.Text()[1:]
	i, _ := strconv.Atoi(str)
	return &redisResult{Res: i}
}

func parseLength(s string) int {
	str := s[1:]
	i, _ := strconv.Atoi(str)
	return i
}
