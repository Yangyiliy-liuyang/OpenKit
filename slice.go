package main

// Find 指定过滤器查找
func Find[T any](vals []T, filter func(t T) bool) T {
	for _, v := range vals {
		if filter(v) {
			return v
		}
	}
	var t T
	return t
}

// FindAll 指定过滤器查找所有
func FindAll[T any](vals []T, filter func(t T) bool) []T {
	var result []T
	for _, v := range vals {
		if filter(v) {
			result = append(result, v)
		}
	}
	return result
}

// Insert 指定下标位置插入
func Insert[T any](idx int, val T, vals []T) []T {
	if idx < 0 || idx > len(vals) {
		panic("idx不合法")
	}
	//扩容
	vals = append(vals, val)
	for i := len(vals) - 1; i > idx; i-- {
		if i-1 > 0 {
			vals[i] = vals[i-1]
		}
	}
	vals[idx] = val
	return vals
}

/*
作业：实现切片的删除操作
实现删除切片特定下标元素的方法。

要求一：能够实现删除操作就可以。
要求二：考虑使用比较高性能的实现。
要求三：改造为泛型方法
要求四：支持缩容，并旦设计缩容机制。
*/
func DelectV1(src []int, index int) []int {
	return []int{}
}

func Delete[Src any](src []Src, index int) []Src {
	src = append(src[:index], src[index+1:]...)
	return src
}
