package redis

import (
	"log"
	"sync"
	"testing"
	"time"
)

func TestSubscribe(t *testing.T) {
	var wg sync.WaitGroup
	c, err := dial()
	if err != nil {
		t.Error(err)
	}

	psc := NewPubSubConn(c, func(msg Message, err error) {

		if err != nil {
			log.Println("=====================", err)
		} else {
			log.Println("message", msg)
		}
	})
	defer psc.Close()

	err = psc.Subscribe("name")
	if err != nil {
		t.Error(err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		con, err := dial()
		if err != nil {
			t.Error(err)
		}
		defer con.Close()

		con.Exec("PUBLISH", "name", "phil")
	}()
	wg.Wait()
	time.Sleep(time.Second * 2)
}
