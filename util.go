package aocutils

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Signed | constraints.Float](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Delta[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return b - a
	}
	return a - b
}

func Min[T constraints.Ordered](t ...T) T {
	var least T
	if len(t) == 0 {
		return least
	}
	least = t[0]
	for _, l := range t {
		if l < least {
			least = l
		}
	}
	return least
}

func Max[T constraints.Ordered](t ...T) T {
	var most T
	if len(t) == 0 {
		return most
	}
	most = t[0]
	for _, l := range t {
		if l > most {
			most = l
		}
	}
	return most
}

// SliceMap - apply a function f to all elements of slice vals, returns the resulting slice.
func SliceMap[T any](vals []T, f func(T) T) []T {
	for i := range vals {
		vals[i] = f(vals[i])
	}
	return vals
}

func SliceMapMutate[T1 any, T2 any](vals []T1, f func(T1) T2) []T2 {
	retVals := make([]T2, len(vals))
	for i := range vals {
		retVals[i] = f(vals[i])
	}
	return retVals
}

// SliceMapNoModify - like SliceMap but does not modify the vals slice.
func SliceMapNoModify[T any](vals []T, f func(T) T) []T {
	tmp := make([]T, len(vals))
	copy(tmp, vals)
	return SliceMap[T](tmp, f)
}

func MapKeys[T comparable, Y any](m map[T]Y) []T {
	retVal := make([]T, 0, len(m))
	for k := range m {
		retVal = append(retVal, k)
	}
	return retVal
}
