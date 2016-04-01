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

	res, err := con.Exec("SET", "name", "name").String()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("response is", res)

	res, err = con.Exec("SET", "age", "12").String()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("response is", res)

	result, err := con.Exec("GET", "name").String()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("result is", result)

	err = con.Pipline("GET", "age")
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = con.Pipline("GET", "name")
	if err != nil {
		log.Println(err.Error())
		return
	}

	arr, err := con.Commit().StringArray()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println("arr is", arr)
}
