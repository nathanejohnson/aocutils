package aocutils

// Stack implements a stack using generics.
type Stack[T any] struct {
	vals []T
}

// NewStack is a trivial factory method.
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

// Push - pushone or more values onto the stack.  Returns a reference to the stack.
func (s *Stack[T]) Push(t ...T) *Stack[T] {
	s.vals = append(s.vals, t...)
	return s
}

// Pop - pop the last value from the stack and return it.  If the
// stack is empty, return the zero value of T
func (s *Stack[T]) Pop() T {
	var t T
	last := len(s.vals) - 1
	if last >= 0 {
		t = s.vals[last]
		s.vals = s.vals[:last]
	}
	return t
}

// PopN - pop up to n values off the stack and return them.
// If there are less values in the stack than n, the
// length of the return will be the number of values in the stack.
func (s *Stack[T]) PopN(n int) []T {
	l := len(s.vals)
	end := l - n
	if end < 0 {
		end = 0
	}
	vals := s.vals[end:]
	s.vals = s.vals[:end]
	return vals
}

func (s *Stack[T]) ind(i int) int {
	if i < 0 {
		i = len(s.vals) + i
	}
	return i
}

// Peek gives us the value at index i and returns it.  If
// index i does not exist it returns the zero value of type T.
// negatives count from the right hand side, so -1 gives the
// last element.
func (s *Stack[T]) Peek(i int) T {
	var t T
	l := len(s.vals)
	i = s.ind(i)
	if i < l && i >= 0 {
		return s.vals[i]
	}
	return t
}

// Poke - poke values in at index i.  If i is invalid, returns nil.
// Otherwise inserts t into stack at index i, returning a reference to
// the stack.  Negative indexes count from the right, however -1
// will insert one before the last element of the stack.  To
// insert at the end, either use Push() or passive a positive
// integer equal to the length of the stack.
func (s *Stack[T]) Poke(i int, t ...T) *Stack[T] {
	i = s.ind(i)
	l := len(s.vals)
	if l < i || i < 0 {
		return nil
	}
	if i == 0 {
		return s.Unshift(t...)
	}
	if i == l {
		return s.Push(t...)
	}
	s.vals = append(s.vals[:i], append(t, s.vals[i:]...)...)
	return s
}

// Set assigns value t at index i.  negative values of i are allowed, counting from right.
// If i is invalid, return value is nil, otherwise stack is returned.
func (s *Stack[T]) Set(i int, t T) *Stack[T] {
	i = s.ind(i)
	l := len(s.vals)
	if l < i || i < 0 {
		return nil
	}
	s.vals[i] = t
	return s
}

// Shift removes the front most element from the stack and returns it.
// If the stack is empty, returns the zero value of type T.
func (s *Stack[T]) Shift() T {
	var t T
	if len(s.vals) > 0 {
		t = s.vals[0]
		s.vals = s.vals[1:]
	}
	return t
}

// ShiftN removes up to n elements from the front of the stack and returns them.
// If n is larger than the size of the internal slice, the length of the return
// will be smaller than n.
func (s *Stack[T]) ShiftN(n int) []T {
	l := len(s.vals)

	newBeg := n
	if l < newBeg {
		newBeg = l
	}
	vals := s.vals[:newBeg]
	s.vals = s.vals[newBeg:]
	return vals
}

// Unshift moves one or more elements t onto the front of the stack.  Returns
// a reference to the stack.
func (s *Stack[T]) Unshift(t ...T) *Stack[T] {
	s.vals = append(t, s.vals...)
	return s
}

// Len returns the length of the internal slice.
func (s *Stack[T]) Len() int {
	return len(s.vals)
}

// Clone returns a new stack with a fresh copy of the internal slice
func (s *Stack[T]) Clone() *Stack[T] {
	vals := make([]T, len(s.vals))
	copy(vals, s.vals)
	return &Stack[T]{vals: vals}
}

// Vals - getter function for the values in the stack.
func (s *Stack[T]) Vals() []T {
	return s.vals
}

// SetVals sets the values of the internal slice and returns
// a reference to the stack s.  No copy is made, just an assignment.
func (s *Stack[T]) SetVals(vals []T) *Stack[T] {
	s.vals = vals
	return s
}
