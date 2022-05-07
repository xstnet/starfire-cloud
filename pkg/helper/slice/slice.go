package slice

import (
	"golang.org/x/exp/constraints"
)

func Unique[T constraints.Integer](val []T) []T {
	var tmp = make(map[T]byte, len(val))
	for _, v := range val {
		tmp[v] = 1
	}

	var res = make([]T, 0, len(tmp))
	for k := range tmp {
		res = append(res, k)
	}

	return res
}
