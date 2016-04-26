package redis

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
)

func parseResults(scanner *bufio.Scanner, n int) Result {
	scanner.Split(bufio.ScanLines)
	if n == 1 {
		result := redisResult{
			Value: parseResult(scanner),
		}
		return result
	}
	var lines = make([]Result, n)
	for i := range lines {
		lines[i] = redisResult{
			Value: parseResult(scanner),
		}
	}
	result := redisResult{
		Value: lines,
	}
	return result
}

func parseResult(scanner *bufio.Scanner) interface{} {
	if scanner.Scan() {
		str := scanner.Text()
		fmt.Println("scan text", str)
		switch str[:1] {
		case "+":
			return parseResponse(str)
		case "-":
			return parseError(str)
		case "*":
			return parseArr(scanner, str)
		case "$":
			length := parseLength(str)
			return parseStr(scanner, length)
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

func parseStr(scanner *bufio.Scanner, length int) string {
	if length < 0 {
		return ""
	}
	if scanner.Scan() {
		nxStr := scanner.Text()
		fmt.Println("Scan str", nxStr)
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
