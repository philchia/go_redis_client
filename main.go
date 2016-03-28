package main

import (
	"log"

	"github.com/philchia/go_redis_driver/redis"
)

func main() {
	con, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	res, err := con.Exec("SET", "name", "Zhai Fei")
	if err != nil {
		log.Println(err.Error())
		return
	}
	str, _ := res.String()
	log.Println(str)
	res, err = con.Exec("GET", "name")
	if err != nil {
		log.Println(err.Error())
		return
	}
	str, _ = res.String()
	log.Println(str)
}
