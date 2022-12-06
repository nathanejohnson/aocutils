package aocutils

// Set is a trivial set implementation based on maps.
type Set[T comparable] struct {
	m map[T]bool
}

// MakeSet - factory method.  Returns a set containing the data.
func MakeSet[T comparable](data []T) *Set[T] {
	s := &Set[T]{
		m: make(map[T]bool),
	}
	return s.Add(data...)
}

// Add one or more elements to the set.  Returns a reference to the set.
func (s *Set[T]) Add(ts ...T) *Set[T] {
	for _, t := range ts {
		s.m[t] = true
	}
	return s
}

// Delete one or more elements from set.  Returns the set.
func (s *Set[T]) Delete(ts ...T) *Set[T] {
	for _, t := range ts {
		delete(s.m, t)
	}
	return s
}

// Len returns length of set.
func (s *Set[T]) Len() int {
	return len(s.m)
}

// InSet returns true if t is in the set, false otherwise.
func (s *Set[T]) InSet(t T) bool {
	return s.m[t]
}

// Map is a getter function for the internal map.
func (s *Set[T]) Map() map[T]bool {
	return s.m
}

// UnorderedSlice returns an unordered slice containing the elements in the set.
func (s *Set[T]) UnorderedSlice() []T {
	t := make([]T, 0, len(s.m))
	for k := range s.m {
		t = append(t, k)
	}
	return t
}
