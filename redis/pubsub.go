package redis

import (
	"sync"
)

var (
	// This is a compiler check
	_ PubSubConn = (*pubSubConn)(nil)
)

// PubSubConn represent a pubsub connection
type PubSubConn interface {
	Subscribe(channels ...interface{}) error
	Unsubscribe(channels ...interface{}) error
	UnsubscribeAll() error
	PSubscribe(channels ...interface{}) error
	PUnsubscribe(channels ...interface{}) error
	PUnsubscribeAll() error
	Close() error
}

// MessageHandler is alias func type for handle published message
type MessageHandler func(Message, error)

// pubSubConn represent a connect for pub sub
type pubSubConn struct {
	handler MessageHandler
	conn    Conn
	closed  bool
	mutex   sync.Mutex
}

// NewPubSubConn create a new pubsub
func NewPubSubConn(conn Conn, handler MessageHandler) PubSubConn {
	pubsub := &pubSubConn{
		conn:    conn,
		handler: handler,
	}
	if handler == nil {
		return nil
	}
	go pubsub.listen()

	return pubsub
}

// listen to published message or subscription message
func (c *pubSubConn) listen() {
	for {
		msg, err := c.readMessage()
		c.handler(msg, err)
		if err != nil {
			return
		}
	}
}

func (c *pubSubConn) readMessage() (Message, error) {
	return c.conn.Read().Message()
}

// Subscribe subscribes to given channels
func (c *pubSubConn) Subscribe(channels ...interface{}) error {
	return c.conn.Send("SUBSCRIBE", channels...)
}

// Unsubscribe unsubscribe from given channels
func (c *pubSubConn) Unsubscribe(channels ...interface{}) error {
	return c.conn.Exec("UNSUBSCRIBE", channels...).Error()
}

// UnsubscribeAll unsubscribe from all subscribed channels
func (c *pubSubConn) UnsubscribeAll() error {
	c.handler = nil
	return c.conn.Exec("UNSUBSCRIBE").Error()
}

// PSubscribe subscribe to given channels with pattern match mode with given handler
func (c *pubSubConn) PSubscribe(channels ...interface{}) error {
	return c.conn.Send("PSUBSCRIBE", channels...)
}

// PUnsubscribe unsubscribe to given pattern match channels
func (c *pubSubConn) PUnsubscribe(channels ...interface{}) error {
	return c.conn.Exec("PUNSUBSCRIBE", channels...).Error()
}

// PUnsubscribeAll unsubscribe to all pattern match channels
func (c *pubSubConn) PUnsubscribeAll() error {
	return c.conn.Exec("PUNSUBSCRIBE").Error()
}

// Close close the PubSubConn connection
func (c *pubSubConn) Close() error {
	return c.conn.Close()
}
