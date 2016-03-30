package redis

import "net"

// Conn represent a connection
type Conn interface {
	Exec(cmd string, args ...interface{}) (res Result, err error)
	Close() error
	Pipline(cmd string, args ...interface{}) error
	Commit() (res Result, err error)
}

// Connect generate a new Redis struct pointer
func Connect(addr string, auth string) (Conn, error) {
	tcpConn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn := &connection{
		Con: tcpConn,
	}
	if len(auth) > 0 {
		_, err := conn.Exec("AUTH", auth)
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}
