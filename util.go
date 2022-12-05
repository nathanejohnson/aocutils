package aocutils

// SliceMap - apply a function f to all elements of slice vals, returns the resulting slice.
func SliceMap[T any](vals []T, f func(T) T) []T {
	for i := range vals {
		vals[i] = f(vals[i])
	}
	return vals
}
