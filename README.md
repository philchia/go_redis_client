# go_redis_driver
go_redis_driver is a redis client for golang

## how to 

```go
con, err := redis.Connect("127.0.0.1:6379", "password")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	res, err := con.Exec("SET", "name", "phil")
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
```