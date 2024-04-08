package sqlx

import (
	"fmt"
	"sync/atomic"
)

// cas 原子操作
func Cas[T any](addr T, old, new int32) bool {
	switch v := addr.(type) {
	case *int32:
		return atomic.CompareAndSwapInt32(v, old, new)
	}
	return false
}

func testCas() {
	var value int32 = 10
	addr := &value

	// 尝试将value从10更改为20
	if Cas[*int32](addr, 10, 20) {
		fmt.Println("CAS succeeded, new value:", *addr)
	} else {
		fmt.Println("CAS failed, value:", *addr)
	}

	// 再次尝试将value从10更改为30（失败，因为value现在已经是20了）
	if Cas[*int32](addr, 10, 30) {
		fmt.Println("CAS succeeded, new value:", *addr)
	} else {
		fmt.Println("CAS failed, value:", *addr)
	}
}
