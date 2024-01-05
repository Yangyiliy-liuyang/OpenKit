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
