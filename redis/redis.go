package redis

import (
	"bytes"
	"net"
	"time"
)

// Conn represent a connection
type Conn interface {
	Exec(cmd string, args ...interface{}) (res Result)
	Close() error
	Pipline(cmd string, args ...interface{}) error
	Commit() (res Result)
}

// Result interface
type Result interface {
	String() (string, error)
	StringArray() ([]string, error)
	StringMap() (map[string]string, error)
	Int() (int, error)
	Int32() (int32, error)
	Int64() (int64, error)
	Float32() (float32, error)
	Float64() (float64, error)
	Array() ([]Result, error)
	Bool() (bool, error)
}

// Option handle the password and time out configuration
type Option struct {
	Auth         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Connect generate a new Redis struct pointer
func Connect(addr string, option *Option) (Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	tcpConn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	conn := &connection{
		Con:  tcpConn,
		Cmd:  new(bytes.Buffer),
		Conf: option,
	}
	if option != nil && len(option.Auth) > 0 {
		_, err := conn.Exec("AUTH", option.Auth).String()
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}
