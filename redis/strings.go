package redis

import "unsafe"

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func joinStrings(strs ...string) string {
	ln := 0
	for _, str := range strs {
		ln += len(str)
	}
	bts := make([]byte, ln)
	ln = 0
	for _, str := range strs {
		ln += copy(bts[ln:], str)
	}

	return string(bts)
}
