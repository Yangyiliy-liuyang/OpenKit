package main

type Number interface {
	int | uint | int64
}

func Sum[T Number](vals []T) T {
	var res T
	for _, v := range vals {
		res = res + v
	}
	return res
}

func Max[T Number](vals []T) T {
	t := vals[0]
	for i := 0; i < len(vals); i++ {
		if t < vals[i] {
			t = vals[i]
		}
	}
	return t
}

func Min[T Number](vals []T) T {
	t := vals[0]
	for i := 0; i < len(vals); i++ {
		if t > vals[i] {
			t = vals[i]
		}
	}
	return t
}

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
