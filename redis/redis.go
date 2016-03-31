package redis

import "net"

// Conn represent a connection
type Conn interface {
	Exec(cmd string, args ...interface{}) (res Result)
	Close() error
	Pipline(cmd string, args ...interface{}) error
	Commit() (res Result)
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
		_, err := conn.Exec("AUTH", auth).String()
		if err != nil {
			return nil, err
		}
	}
	return conn, nil
}
