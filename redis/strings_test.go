package redis

import (
	"log"
	"strings"
	"testing"
)

func TestJoinStrings(t *testing.T) {
	strs := []string{"string1", "string2", "string3", "string4", "string5", "string6", "string7", "string1", "string2", "string3", "string4", "string5", "string6", "string7", "string2", "string3", "string4", "string5", "string6", "string7", "string1", "string2", "string3", "string4", "string5", "string6", "string7"}

	r := joinStrings(strs...)

	s := strings.Join(strs, "")
	if s != r {
		t.Fail()
	}
}

func TestBytes2String(t *testing.T) {
	bts := []byte("hello")
	str := bytes2str(bts)
	if str != "hello" {
		t.Fail()
	}
}

func TestString2Bytes(t *testing.T) {
	str := "hello"
	bts := str2bytes(str)
	log.Println(bts)
	if string(bts) != str {
		t.Fail()
	}
}
