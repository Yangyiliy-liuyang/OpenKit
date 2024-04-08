package string

import "unsafe"

// BytesToString  []byte to string
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
