package utils

import (
	"encoding/json"
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

func StringMarshal(data any) string {
	body, _ := json.Marshal(data)
	return String(body)
}

func StringUnmarshal(data string, obj any) {
	_ = json.Unmarshal(Bytes(data), obj)
}
