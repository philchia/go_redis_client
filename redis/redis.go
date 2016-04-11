package redis

import "net"

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
	Array() ([]Result, error)
	Bool() (bool, error)
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
