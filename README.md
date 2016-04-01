# go_redis_driver
go_redis_driver is a redis client for golang

# How to

## Single command

```go
	con, err := redis.Connect("127.0.0.1:6379", "password")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	res, err := con.Exec("SET", "name", "your name").String()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
	
	res, err = con.Exec("GET", "name").String()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
```

## Pipline

```go
	con, err := redis.Connect("127.0.0.1:6379", "password")
	if err != nil {
		log.Fatal(err)
	}
	defer con.Close()

	err = con.Pipline("SET", "name", "your name")
	if err != nil {
		log.Fatal(err)
	}
	
	err = con.Pipline("SET", "gender", "female")
	if err != nil {
		log.Fatal(err)
	}
	
	_, err := con.Commit()
	if err != nil {
		log.Fatal(err)
	}
	
	res, err = con.Exec("GET", "name").String()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(res)
```

#Todo

* Pub/Sub
* Connection pool
