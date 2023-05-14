package utils

import (
	"reflect"
	"unsafe"
)

func Bytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
