package main

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
