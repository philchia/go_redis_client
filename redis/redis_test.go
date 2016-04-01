package redis_test

import (
	"log"
	"testing"

	"github.com/philchia/go_redis_driver/redis"
)

// func TestSetGetKey(t *testing.T) {
// 	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	conn.Exec("SET", "name", "chia")
// 	res, err := conn.Exec("GET", "name").String()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(res)
// }

// func TestSetGetHash(t *testing.T) {
// 	conn, err := redis.Connect("127.0.0.1:6379", "112919147")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	defer conn.Close()

// 	conn.Pipline("HSET", "profile", "name", "chia")
// 	conn.Pipline("HSET", "profile", "age", "12")
// 	conn.Pipline("HSET", "profile", "gender", "male")
// 	responses, err := conn.Commit().Array()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, resp := range responses {
// 		res, _ := resp.String()
// 		t.Log(res)
// 	}

// 	conn.Pipline("HGET", "profile", "name")
// 	conn.Pipline("HGET", "profile", "age")
// 	strArr, err := conn.Commit().Array()
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	for _, resp := range strArr {
// 		res, _ := resp.String()
// 		t.Log(res)
// 	}
// }

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
