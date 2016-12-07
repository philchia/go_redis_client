package redis

import (
	"reflect"
	"testing"
)

func Test_str2bytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			"case1",
			args{
				"hello",
			},
			[]byte("hello"),
		},
		{
			"case2",
			args{
				"hello, world!",
			},
			[]byte("hello, world!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := str2bytes(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("str2bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bytes2str(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{
				[]byte("hello"),
			},
			("hello"),
		},
		{
			"case2",
			args{
				[]byte("hello, world!"),
			},
			"hello, world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bytes2str(tt.args.b); got != tt.want {
				t.Errorf("bytes2str() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinStrings(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"case1",
			args{
				[]string{"hello,", " ", "world!"},
			},
			"hello, world!",
		},
		{
			"case2",
			args{
				[]string{"hello", " ", "wo", "rl", "d", "!"},
			},
			"hello world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinStrings(tt.args.strs...); got != tt.want {
				t.Errorf("joinStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
