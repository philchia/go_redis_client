package redis_test

import (
	"log"
	"testing"

	"github.com/philchia/go_redis_driver/redis"
)

// func BenchmarkConnection(b *testing.B) {
// 	log.Println("redis proxy connection")
// 	for i := 0; i < b.N; i++ {
// 		conn, err := redis.Connect("127.0.0.1:6379", "112919147")
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		conn.Close()
// 	}
// }

func BenchmarkSetKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := redis.Connect("127.0.0.1:6379", "112919147")
		if err != nil {
			log.Println(err.Error())
		}
		conn.Exec("SET", "name", "chia")
		res, err := conn.Exec("GET", "name")
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(res.String())
		conn.Close()
	}
}
