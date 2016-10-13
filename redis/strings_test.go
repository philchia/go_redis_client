package redis

import (
	"log"
	"testing"
)

func TestJoinStrings(t *testing.T) {
	strs := []string{"string1", "string2", "string3", "string4", "string5", "string6", "string7", "string1", "string2", "string3", "string4", "string5", "string6", "string7", "string2", "string3", "string4", "string5", "string6", "string7", "string1", "string2", "string3", "string4", "string5", "string6", "string7"}

	_ = joinStrings(strs...)
}

func TestBytes2String(t *testing.T) {
	bts := []byte("hello")
	str := bytes2str(bts)
	log.Println(str)
}

func TestString2Bytes(t *testing.T) {
	str := "hello"
	bts := str2bytes(str)
	log.Println(bts)
}
