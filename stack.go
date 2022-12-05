package aocutils

type Stack[T any] struct {
	vals []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(t ...T) {
	s.vals = append(s.vals, t...)
}

func (s *Stack[T]) Pop() T {
	var t T
	if len(s.vals) > 0 {
		last := len(s.vals) - 1
		t = s.vals[last]
		// fmt.Printf("len before: %d ", len(s.vals))
		s.vals = s.vals[:last]
		// fmt.Printf("len after: %d\n", len(s.vals))
	}
	return t
}

func (s *Stack[T]) PopN(n int) []T {
	end := len(s.vals) - n
	if end < 0 {
		end = 0
	}
	vals := s.vals[end:]
	s.vals = s.vals[:end]
	return vals
}

func (s *Stack[T]) Peek(index int) T {
	var t T
	if index < len(s.vals) {
		return s.vals[index]
	}
	return t
}

func (s *Stack[T]) Shift() T {
	var t T
	if len(s.vals) > 0 {
		t = s.vals[0]
		s.vals = s.vals[1:]
	}
	return t
}

func (s *Stack[T]) ShiftN(n int) []T {
	newBeg := n
	if len(s.vals) < newBeg {
		newBeg = len(s.vals)
	}
	vals := s.vals[:newBeg]
	s.vals = s.vals[newBeg:]
	return vals
}

func (s *Stack[T]) Unshift(t ...T) {
	tl := len(t)
	newVals := make([]T, len(s.vals)+tl)
	copy(newVals, t)
	copy(newVals[tl:], s.vals)
	s.vals = newVals
}

func (s *Stack[T]) Len() int {
	return len(s.vals)
}

func (s *Stack[T]) Clone() *Stack[T] {
	vals := make([]T, len(s.vals))
	copy(vals, s.vals)
	return &Stack[T]{vals: vals}
}
