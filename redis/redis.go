package redis

import (
	"bufio"
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
	Error() error
	OK() bool
	PONG() bool
	String() (string, error)
	Strings() ([]string, error)
	StringMap() (map[string]string, error)
	Int() (int, error)
	Float64() (float64, error)
	Results() ([]Result, error)
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
		BW:   bufio.NewWriter(tcpConn),
		BR:   bufio.NewReader(tcpConn),
		Conf: option,
		Crlf: []byte{'\r', '\n'},
	}
	if option != nil && len(option.Auth) > 0 {
		conn.Exec("AUTH", option.Auth)
	}
	return conn, nil
}
