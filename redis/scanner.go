package redis

import (
	"bufio"
	"errors"
	"strconv"
)

func parseResults(scanner *bufio.Scanner, n int) Result {
	scanner.Split(bufio.ScanLines)
	if n == 1 {
		return parseResult(scanner)
	}
	var lines = make([]Result, n)
	for i := range lines {
		lines[i] = parseResult(scanner)
	}
	return lines[n-1]
}

func parseResult(scanner *bufio.Scanner) interface {
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
			return &redisResult{Value: errors.New("Unknow response")}
		}
	}
	return nil
}

func parseResponse(str string) string {
	return str[1:]
}

func parseError(str string) error {
	return errors.New(str[1:])
}

func parseStr(scanner *bufio.Scanner, str string) string {
	length := parseLength(str)
	if length < 0 {
		return ""
	}
	if scanner.Scan() {
		nxStr := scanner.Text()
		return nxStr[:length]
	}
	return ""
}

func parseArr(scanner *bufio.Scanner, str string) []interface{} {
	length := parseLength(str)
	if length < 0 {
		return nil
	}
	res := make([]interface{}, length)
	for i := range res {
		res[i] = parseResult(scanner)

	}
	return res
}

func parseInt(str string) int {
	a := str[1:]
	i, _ := strconv.Atoi(a)
	return i
}

func parseLength(s string) int {
	str := s[1:]
	i, _ := strconv.Atoi(str)
	return i
}
