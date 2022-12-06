package aocutils

// Window is meant for sliding window loops where you want to maintain statistics about
// how many occurrences of a given value of type T are in a subset of the data.
type Window[T comparable] struct {
	m map[T]int
}

// NewWindow - Factory method for Window.  Initializes state with data, returns a Window
func NewWindow[T comparable](data []T) *Window[T] {
	m := make(map[T]int)
	for _, d := range data {
		m[d]++
	}
	return &Window[T]{
		m: m,
	}
}

// Slide moves the window down.  remove is the value being passed out of the window,
// add is the value passing into the window.
func (w *Window[T]) Slide(remove, add T) *Window[T] {
	if remove == add {
		return w
	}
	w.m[remove]--
	if w.m[remove] == 0 {
		delete(w.m, remove)
	}
	w.m[add]++
	return w
}

// Len - length of the data set inside of the window.
func (w *Window[T]) Len() int {
	return len(w.m)
}

// UnorderedSlice - this returns the data inside the internal map, unordered.
func (w *Window[T]) UnorderedSlice() []T {
	vals := make([]T, 0, len(w.m))
	for k := range w.m {
		vals = append(vals, k)
	}
	return vals
}

// Map - getter function for the internal map.
func (w *Window[T]) Map() map[T]int {
	return w.m
}
