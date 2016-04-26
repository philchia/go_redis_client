package redis_test

import (
	"log"
	"testing"

	redis1 "github.com/garyburd/redigo/redis"
	"github.com/philchia/go_redis_client/redis"
)

func TestSetGetString(t *testing.T) {
	opt := redis.Option{
		Auth:         "112919147",
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
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
	log.Println("test set get string succeed")
}

func TestSetGetInt(t *testing.T) {

	opt := redis.Option{
		Auth:         "112919147",
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
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
	log.Println("test set get int succeed")

}

func TestArr(t *testing.T) {
	opt := redis.Option{
		Auth:         "112919147",
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
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

	res, err := conn.Exec("HGETALL", "Profile").StringArray()
	if err != nil {
		t.Fatal(err)
	}

	log.Println(res)
	log.Println("test set get map succeed")
}

func TestMap(t *testing.T) {
	opt := redis.Option{
		Auth:         "112919147",
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
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

	log.Println(res)
	log.Println("test set get map succeed")

}

func BenchmarkRedigo(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	defer conn.Close()
	conn.Do("AUTH", "112919147")
	b.ResetTimer()
	conn.Do("SET", "name", "chia")
	for i := 0; i < b.N; i++ {
		_, err := redis1.String(conn.Do("GET", "name"))
		if err != nil {
			b.Fail()
		}
	}
}

func BenchmarkRedigoPing(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	defer conn.Close()
	conn.Do("AUTH", "112919147")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conn.Do("PING")
	}
}

func BenchmarkPing(b *testing.B) {
	opt := redis.Option{
		Auth: "112919147",
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		conn.Exec("PING")
	}
}

func BenchmarkSetKey(b *testing.B) {
	opt := redis.Option{
		Auth:         "112919147",
		ReadTimeout:  1000000,
		WriteTimeout: 1000000,
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
	if err != nil {
		b.Fatalf("error while connection %v", err)
	}
	defer conn.Close()

	b.ResetTimer()
	conn.Exec("SET", "name", "chia")

	for i := 0; i < b.N; i++ {
		_, err := conn.Exec("GET", "name").String()
		if err != nil {
			b.Fatal(err)
		}
	}
}
