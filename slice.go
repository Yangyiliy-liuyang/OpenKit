package main

import "errors"

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
func Insert[T any](vals []T, idx int, val T) []T {
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

// Delete 指定下标位置删除,返回新切片数组、被删除元素、错误。
func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index > length {
		var t T
		//return nil, t, ErrIndexOutOfRange(length, index)
		return nil, t, errors.New("index不合法")
	}
	//src[index+1:] 表示从切片 src 的索引 index+1 开始到末尾的所有元素组成的切片。
	//然后使用 ... 将这个切片展开，作为函数 append 的参数,实现逐个添加到 append 函数的参数列表中
	res := src[index]
	src = append(src[:index], src[index+1:]...)
	return src, res, nil
}

func DelectV1(src []int, index int) []int {
	return []int{}
}

func test() {

}
