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
			return parseResponse(str)
		case "-":
			return parseError(str)
		case "*":
			return parseArr(scanner, str)
		case "$":
			return parseStr(scanner, str)
		case ":":
			return parseInt(str)
		default:
			return &redisResult{Res: errors.New("Unknow response")}
		}
	}
	return nil
}

func parseResponse(str string) Result {
	return &redisResult{Res: str[1:]}
}

func parseError(str string) Result {
	return &redisResult{Res: errors.New(str[1:])}
}

func parseStr(scanner *bufio.Scanner, str string) Result {
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

func parseArr(scanner *bufio.Scanner, str string) Result {
	length := parseLength(str)
	if length < 0 {
		return nil
	}
	res := make([]Result, length)
	for i := range res {
		res[i] = parseResult(scanner)

	}
	return &redisResult{Res: res}
}

func parseInt(str string) Result {
	a := str[1:]
	i, _ := strconv.Atoi(a)
	return &redisResult{Res: i}
}

func parseLength(s string) int {
	str := s[1:]
	i, _ := strconv.Atoi(str)
	return i
}
