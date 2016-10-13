# go_redis_client
go_redis_client is a redis client for golang

# How to

## Single command

```go
	opt := redis.Option{
		Auth: "password",
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

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
	opt := redis.Option{
		Auth: "password",
	}
	conn, err := redis.Connect("127.0.0.1:6379", &opt)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

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

```

#Todo

* Pub/Sub
* Connection pool
