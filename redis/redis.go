package redis

import (
	"bufio"
	"net"
	"time"
)

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
		if conn.Exec("AUTH", conn.Conf.Auth).Error() != nil {
			return nil, err
		}
	}
	return conn, nil
}
