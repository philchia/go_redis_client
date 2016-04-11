package redis_test

import (
	"log"
	"testing"

	redis1 "github.com/garyburd/redigo/redis"
	"github.com/philchia/go_redis_driver/redis"
)

func TestSetGetString(t *testing.T) {
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

func TestSetGetInt(t *testing.T) {

	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("SET", "age", 25)

	res, err := conn.Exec("GET", "age").Int()
	if err != nil {
		t.Fatal(err)
	}
	if res != 25 {
		t.Fatal("Get wrong age")
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

func BenchmarkRedigo(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	conn.Do("AUTH", "112919147")
	for i := 0; i < b.N; i++ {
		conn.Do("SET", "name", "chia")
		_, err := redis1.String(conn.Do("GET", "name"))
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkSetKey(b *testing.B) {
	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		conn.Exec("SET", "name", "chia")
		_, err := conn.Exec("GET", "name").String()
		if err != nil {
			b.Fail()
		}
	}
}
