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
func Connect(host, port string, options ...*Option) (Conn, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
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
		Crlf: []byte{'\r', '\n'},
	}
	if len(options) > 0 {
		conn.Conf = options[0]
	}
	if conn.Conf != nil && len(conn.Conf.Auth) > 0 {
		conn.Exec("AUTH", conn.Conf.Auth)
	}
	return conn, nil
}
