package redis

import (
	"bufio"
	"errors"
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	// CMDExec represent the exec command
	CMDExec = "EXEC"
)

var (
	// here is a compiler check
	_ Conn = (*connection)(nil)
)

var (
	// ErrBufferFull ...
	ErrBufferFull = errors.New("response too long")
	// ErrResponseFormat ...
	ErrResponseFormat = errors.New("incorrect response format")
	// OK represent response "OK"
	OK = &redisResult{Value: "OK"}
	// PONG represent response "PONG"
	PONG = &redisResult{Value: "PONG"}
)

// Conn represent a connection
type Conn interface {
	Send(cmd string, args ...interface{}) error
	Exec(cmd string, args ...interface{}) (res Result)
	Close() error
	Pipline(cmd string, args ...interface{}) error
	Commit() (res Result)
	Read() Result
}

// connection ...
type connection struct {
	Con       net.Conn
	BW        *bufio.Writer
	BR        *bufio.Reader
	QueueSize int
	Conf      *Option
	Mx        sync.Mutex

	Draft [64]byte
	Crlf  []byte
}

// Close close the connection to redis server
func (c *connection) Close() error {
	return c.Con.Close()
}

// Send will only write cmd to server
func (c *connection) Send(cmd string, args ...interface{}) error {
	err := c.writeCmd(cmd, args...)
	if err != nil {
		return err
	}
	return c.flush()
}

// Exec do a single command
func (c *connection) Exec(cmd string, args ...interface{}) Result {
	c.Mx.Lock()
	defer c.Mx.Unlock()

	c.QueueSize++

	if err := c.writeCmd(cmd, args...); err != nil {
		return &redisResult{Value: err}
	}

	if err := c.flush(); err != nil {
		return &redisResult{Value: err}
	}
	return c.Read()
}

//Pipline cache all the command
func (c *connection) Pipline(cmd string, args ...interface{}) error {
	c.Mx.Lock()
	defer c.Mx.Unlock()
	if c.QueueSize == 0 {
		if err := c.writeCmd("MULTI"); err != nil {
			return err
		}
		c.QueueSize++
	}
	c.QueueSize++
	return c.writeCmd(cmd, args...)
}

func (c *connection) Commit() Result {
	return c.Exec("EXEC")
}

func (c *connection) flush() error {
	if c.Conf != nil && c.Conf.WriteTimeout > 0 {
		c.Con.SetWriteDeadline(time.Now().Add(c.Conf.WriteTimeout))
	}
	return c.BW.Flush()
}

func (c *connection) Read() Result {
	size := c.QueueSize
	c.QueueSize = 0
	if c.Conf != nil && c.Conf.ReadTimeout > 0 {
		c.Con.SetReadDeadline(time.Now().Add(c.Conf.WriteTimeout))
	}

	if size > 1 {
		res := make([]Result, size)
		for i := range res {
			res[i] = c.readReply()
		}
		return &redisResult{Value: res}
	}

	return c.readReply()
}

func (c *connection) readReply() Result {
	bts, err := c.readLine()
	if err != nil {
		return &redisResult{Value: err}
	}

	switch bts[0] {

	case '+':
		switch {
		case len(bts) == 3 && bts[1] == 'O' && bts[2] == 'K':
			return OK
		case len(bts) == 5 && bts[1] == 'P' && bts[2] == 'O' && bts[3] == 'O' && bts[4] == 'G':
			return PONG
		default:
			return &redisResult{Value: bytes2str(bts[1:])}
		}

	case '-':
		return &redisResult{Value: errors.New(bytes2str(bts[1:]))}

	case ':':
		res, err := parseInt(bts[1:])
		if err != nil {
			return &redisResult{Value: err}
		}
		return &redisResult{Value: res}

	case '$':
		line, err := c.readLine()
		if err != nil {
			return &redisResult{Value: err}
		}
		return &redisResult{Value: bytes2str(line)}
	case '*':
		count, err := parseInt(bts[1:])
		if err != nil {
			return &redisResult{Value: err}
		}
		res := make([]Result, count)
		for i := range res {
			res[i] = c.readReply()
		}
		return &redisResult{Value: res}
	}
	return &redisResult{Value: errors.New("unexpected response line")}
}

// parseInt parses an integer reply.
func parseInt(p []byte) (int64, error) {
	if len(p) == 0 {
		return 0, errors.New("malformed integer")
	}

	var negate bool
	if p[0] == '-' {
		negate = true
		p = p[1:]
		if len(p) == 0 {
			return 0, errors.New("malformed integer")
		}
	}

	var n int64
	for _, b := range p {
		n *= 10
		if b < '0' || b > '9' {
			return 0, errors.New("illegal bytes in length")
		}
		n += int64(b - '0')
	}

	if negate {
		n = -n
	}
	return n, nil
}

func (c *connection) readLine() ([]byte, error) {
	bts, err := c.BR.ReadSlice('\n')
	if err == bufio.ErrBufferFull {
		return nil, ErrBufferFull
	}
	if err != nil {
		return bts, err
	}

	i := len(bts) - 2
	if i < 0 || bts[i] != '\r' {
		return nil, ErrResponseFormat
	}
	return bts[:i], err
}

func (c *connection) writeCmd(cmd string, args ...interface{}) (err error) {
	c.writeLen('*', len(args)+1)
	c.writeString(cmd)
	for _, arg := range args {
		err = c.writeAny(arg)
		if err != nil {
			return err
		}
	}
	return
}

func (c *connection) writeAny(cmd interface{}) error {
	switch val := cmd.(type) {
	case string:
		return c.writeString(val)
	case int8:
		return c.writeInt64(int64(val))
	case int16:
		return c.writeInt64(int64(val))
	case int32:
		return c.writeInt64(int64(val))
	case int64:
		return c.writeInt64(val)
	case int:
		return c.writeInt64(int64(val))
	case []byte:
		return c.writeBytes(val)
	case float32:
		return c.writeFloat64(float64(val))
	case float64:
		return c.writeFloat64(val)
	case bool:
		return c.writeBool(val)
	case nil:
		return c.writeNil()
	default:
		return errors.New("unknow type")
	}
}

func (c *connection) writeLen(pre byte, length int) error {
	i := cap(c.Draft)
	c.Draft[i-1] = '\n'
	c.Draft[i-2] = '\r'
	i -= 3
	for {
		n := length%10 + '0'
		c.Draft[i] = byte(n)
		length = length / 10
		i--
		if length == 0 {
			c.Draft[i] = pre
			break
		}
	}
	_, err := c.BW.Write(c.Draft[i:])
	return err
}

func (c *connection) writeCrlf() error {
	_, err := c.BW.Write(c.Crlf)
	return err
}

func (c *connection) writeString(s string) error {
	c.writeLen('$', len(s))
	c.BW.WriteString(s)
	return c.writeCrlf()
}

func (c *connection) writeBytes(bts []byte) error {
	c.writeLen('$', len(bts))
	c.BW.Write(bts)
	return c.writeCrlf()
}

func (c *connection) writeInt64(i int64) error {
	return c.writeBytes(strconv.AppendInt(c.Draft[:0], i, 10))
}

func (c *connection) writeFloat64(f float64) error {
	return c.writeBytes(strconv.AppendFloat(c.Draft[:0], f, 'g', -1, 64))
}

func (c *connection) writeBool(b bool) error {
	if b {
		return c.writeInt64(1)
	}
	return c.writeInt64(0)
}

func (c *connection) writeNil() error {
	return c.writeString("")
}
