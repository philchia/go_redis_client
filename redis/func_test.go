package redis

import (
	"testing"

	redis1 "github.com/garyburd/redigo/redis"
)

func dial() (Conn, error) {
	return Connect("127.0.0.1", "6379")
}

func TestSetGetString(t *testing.T) {
	conn, err := dial()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("SET", "name", "chia")

	res, err := conn.Exec("GET", "name").String()
	if err != nil {
		t.Fatal(err)
	}

	if res != "chia" {
		t.Fatal("Get wrong name")
	}
}

func TestSetGetInt(t *testing.T) {
	conn, err := dial()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	conn.Exec("SET", "age", 25)

	res, err := conn.Exec("GET", "age").Int()
	if err != nil {
		t.Fatal(err)
	}
	if res != 25 {
		t.Fatal("Get wrong age")
	}

}

func TestArr(t *testing.T) {
	conn, err := dial()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Exec("HSET", "Profile", "name", "phil").Int()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("HSET", "Profile", "age", "12").Int()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("HGETALL", "Profile").Strings()
	if err != nil {
		t.Fatal(err)
	}

}

func TestMap(t *testing.T) {
	conn, err := dial()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Exec("HSET", "Profile", "name", "phil").Int()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("HSET", "Profile", "age", "12").Int()
	if err != nil {
		t.Fatal(err)
	}

	_, err = conn.Exec("HGETALL", "Profile").StringMap()
	if err != nil {
		t.Fatal(err)
	}

}

func TestPing(t *testing.T) {
	conn, err := dial()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	if !conn.Exec("PING").PONG() {
		t.Fail()
	}
}

func TestPipline(t *testing.T) {
	conn, err := dial()
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	err = conn.Pipline("GET", "name")
	if err != nil {
		t.Fatal(err)
	}

	err = conn.Pipline("GET", "age")
	if err != nil {
		t.Fatal(err)
	}

	if conn.Commit().Error() != nil {
		t.Fail()
	}

}

func BenchmarkRedigoGetKey(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	defer conn.Close()
	conn.Do("SET", "name", "chia")

	for i := 0; i < b.N; i++ {
		redis1.String(conn.Do("GET", "name"))
	}
}

func BenchmarkGetKey(b *testing.B) {
	conn, err := dial()
	if err != nil {
		b.Fatalf("error while connection %v", err)
	}
	defer conn.Close()
	conn.Exec("SET", "name", "phil")

	for i := 0; i < b.N; i++ {
		conn.Exec("GET", "name").String()
	}
}

func BenchmarkRedigoGetIntKey(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	defer conn.Close()
	conn.Do("SET", 1, "one")

	for i := 0; i < b.N; i++ {
		redis1.Int(conn.Do("GET", 1))
	}
}

func BenchmarkGetIntKey(b *testing.B) {
	conn, err := dial()
	if err != nil {
		b.Fatalf("error while connection %v", err)
	}
	defer conn.Close()

	conn.Exec("SET", 1, "one")

	for i := 0; i < b.N; i++ {
		conn.Exec("GET", 1).Int()
	}
}

func BenchmarkRedigoSetKey(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		conn.Do("SET", "name", "chia")
	}
}

func BenchmarkSetKey(b *testing.B) {
	conn, err := dial()
	if err != nil {
		b.Fatalf("error while connection %v", err)
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		conn.Exec("SET", "name", "chia")
	}
}

func BenchmarkRedigoPing(b *testing.B) {
	conn, err := redis1.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		b.Fail()
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		conn.Do("PING")
	}
}

func BenchmarkPing(b *testing.B) {

	conn, err := dial()
	if err != nil {
		b.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		conn.Exec("PING")
	}
}
