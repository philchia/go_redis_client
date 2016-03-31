# go_redis_driver
go_redis_driver is a redis client for golang

## how to 

```go
	con, err := redis.Connect("127.0.0.1:6379", "112919147")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	res, err := con.Exec("SET", "name", "your name").String()
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
```