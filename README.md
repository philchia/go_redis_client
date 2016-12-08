# go_redis_client

go_redis_client is a redis client for Go

[![Golang](https://img.shields.io/badge/Language-Go-green.svg?style=flat)](https://golang.org)
[![Build Status](https://travis-ci.org/philchia/go_redis_client.svg?branch=master)](https://travis-ci.org/philchia/go_redis_client)
[![Coverage Status](https://coveralls.io/repos/github/philchia/go_redis_client/badge.svg?branch=dev)](https://coveralls.io/github/philchia/go_redis_client?branch=dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/philchia/go_redis_client)](https://goreportcard.com/report/github.com/philchia/go_redis_client)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/philchia/go_redis_client/redis?status.svg)](https://godoc.org/github.com/philchia/go_redis_client/redis)

## Warnning

go_redis_client is under heavy development, if you want to use it in your project, **vendor** it!

## How to

### Single command

```go
    opt := redis.Option{
        Auth: "password",
    }
    conn, err := redis.Connect("127.0.0.1", "6379", &opt)
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

### Pipline

```go
    opt := redis.Option{
        Auth: "password",
    }
    conn, err := redis.Connect("127.0.0.1", "6379", &opt)
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

    strs, err := con.Exec("GET", "name", "gender").Strings()
    if err != nil {
        log.Fatal(err)
    }
	log.Println(strs)

```

## Todo

* Pub/Sub
* Connection pool

## License

go_redis_client code is published under MIT license