package redis_test

import (
	"log"
	"testing"

	"github.com/philchia/go_redis_driver/redis"
)

func TestSetGetKey(t *testing.T) {
	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("SET", "name", "chia")
	res, err := conn.Exec("GET", "name").String()
	if err != nil {
		t.Fatal(err)
	}
	if res != "chia" {
		t.Fatal("Get wrong name")
	}
}

func TestMap(t *testing.T) {
	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Exec("HSET", "Profile", "name", "phil").Int()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("HSET", "Profile", "age", "12").Int()
	if err != nil {
		t.Fatal(err)
	}

	res, err := conn.Exec("HGETALL", "Profile").StringMap()
	if err != nil {
		t.Fatal(err)
	}
	if res["name"] != "phil" {
		t.Fatal("Get wrong name")
	}
	if res["age"] != "12" {
		t.Fatal("Get wrong age")
	}
	log.Println(res)
}

func BenchmarkSetKey(b *testing.B) {
	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		conn.Exec("SET", "name", "chia")
		res, err := conn.Exec("GET", "name").String()
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(res)
	}
}
