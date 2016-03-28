package redis

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"unicode/utf8"
)

// connection ...
type connection struct {
	Con net.Conn
	Cmd bytes.Buffer
}

// Close close the connection to redis server
func (c *connection) Close() error {
	return c.Con.Close()
}

// Exec do a single command
func (c *connection) Exec(cmd string, args ...interface{}) (res Result, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); ok {
				res = nil
			}
		}
	}()

	c.writeCmd(cmd, args...)
	c.Flush()
	return c.read()
}

//Pip pipline the command
func (c *connection) Pip(cmd string, args ...interface{}) error {
	return c.writeCmd(cmd, args...)
}

func (c *connection) Commit() (res Result, err error) {
	err = c.Flush()
	if err != nil {
		return nil, err
	}
	return c.read()
}

//Flush all command
func (c *connection) Flush() error {
	defer c.clear()

	_, err := c.Con.Write(c.Cmd.Bytes())

	if err != nil {
		return err
	}
	return nil
}

func (c *connection) read() (Result, error) {
	buf := make([]byte, 512)
	n, err := c.Con.Read(buf)
	if err != nil {
		return nil, err
	}
	tmp := buf[:n]

	switch string(tmp[:1]) {
	case "+":
		res := parseResponse(string(tmp[1:]))
		result := &redisResult{
			Res: res,
		}
		return result, nil

	case "-":
		err := parseError(string(tmp[1:]))
		return nil, err

	case "$":
		str := parseSingleLineString(string(tmp[1:]))
		result := &redisResult{
			Res: str,
		}
		return result, nil

	case "*":
		arr := parseArr(string(tmp[1:]))
		result := &redisResult{
			Res: arr,
		}
		return result, nil

	case ":":
		num := parseInt(string(tmp[1:]))
		result := &redisResult{
			Res: num,
		}
		return result, nil

	}
	return nil, nil
}

func parseResponse(s string) string {
	return strings.TrimRight(s, "\r\n")
}

func parseSingleLineString(s string) string {
	return strings.Split(s, "\r\n")[1]
}

func parseInt(s string) int {
	str := strings.TrimRight(s, "\r\n")
	i, _ := strconv.Atoi(str)
	return i
}

func parseArr(s string) []string {
	return strings.Split(s, "\r\n")[1:]
}

func parseError(s string) error {
	return errors.New(strings.TrimRight(s, "\r\n"))
}

func (c *connection) writeCmd(cmd string, args ...interface{}) (err error) {
	c.writeLength(len(args) + 1)
	c.writeString(cmd)
	for _, arg := range args {
		switch arg.(type) {
		case string:
			c.writeString(arg.(string))
		case int32:
			c.writeInt32(arg.(int32))
		case int64:
			c.writeInt64(arg.(int64))
		case []byte:
			c.writeBytes(arg.([]byte))
		case float32:
			c.writeFloat32(arg.(float32))
		case float64:
			c.writeFloat64(arg.(float64))
		default:
			panic(errors.New("unknow type"))
		}
	}
	return
}

func (c *connection) writeLength(length int) {
	str := fmt.Sprintf("*%d\r\n", length)
	_, err := c.Cmd.WriteString(str)
	if err != nil {
		panic(err)
	}
}

func (c *connection) writeString(s string) {
	str := fmt.Sprintf("$%d\r\n%s\r\n", utf8.RuneCountInString(s), s)

	_, err := c.Cmd.WriteString(str)
	if err != nil {
		panic(err)
	}
}

func (c *connection) writeBytes(bts []byte) {
	str := fmt.Sprintf("$%d\r\n%s\r\n", utf8.RuneCount(bts), bts)
	_, err := c.Cmd.WriteString(str)
	if err != nil {
		panic(err)
	}
}

func (c *connection) writeInt64(i int64) {
	str := fmt.Sprint(i)
	c.writeString(str)
}

func (c *connection) writeFloat64(f float64) {
	str := fmt.Sprint(f)
	c.writeString(str)
}

func (c *connection) writeInt32(i int32) {
	str := fmt.Sprint(i)
	c.writeString(str)
}

func (c *connection) writeFloat32(f float32) {
	str := fmt.Sprint(f)
	c.writeString(str)
}

func (c *connection) clear() {
	c.Cmd.Reset()
}
