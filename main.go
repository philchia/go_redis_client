package main

import (
	"log"

	"github.com/philchia/go_redis_driver/redis"
)

func main() {
	con, err := redis.Connect("127.0.0.1:6379", "password")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	res, err := con.Exec("SET", "name", "name").String()
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(res)
	res, err = con.Exec("GET", "name").String()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(res)
}
