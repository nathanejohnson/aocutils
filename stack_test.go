package aocutils

import (
	"reflect"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	s := NewStack[string]()
	s.vals = append(s.vals, "a", "b", "c")
	v := s.Pop()
	if v != "c" {
		t.Errorf("unexpected value %s", v)
	}
}

func TestStack_PopN(t *testing.T) {
	s := NewStack[string]()
	var expected []string = nil
	vals := s.PopN(3)
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
	s.vals = append(s.vals, "a", "b", "c", "d", "e")
	vals = s.PopN(3)
	expected = []string{"a", "b"}
	if !reflect.DeepEqual(s.vals, expected) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
	expected = []string{"c", "d", "e"}
	if !reflect.DeepEqual(vals, expected) {
		t.Errorf("expected %#v got %#v", expected, vals)
	}

	vals = s.PopN(3)
	expected = []string{}
	if !reflect.DeepEqual(s.vals, expected) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
	expected = []string{"a", "b"}
	if !reflect.DeepEqual(vals, expected) {
		t.Errorf("expected %#v got %#v", expected, vals)
	}
}

func TestStack_Push(t *testing.T) {
	s := NewStack[string]()
	s.Push("a", "b", "c")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
}

func TestStack_Unshift(t *testing.T) {
	s := NewStack[string]()
	s.vals = append(s.vals, "a", "b", "c")
	s.Unshift("e", "f", "g")
	expected := []string{"e", "f", "g", "a", "b", "c"}
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
}
func TestStack_ShiftN(t *testing.T) {
	s := NewStack[string]()
	var expected []string = nil
	vals := s.ShiftN(3)
	if !reflect.DeepEqual(expected, vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
	s.vals = []string{"e", "f", "g", "a", "b", "c"}
	vals = s.ShiftN(3)
	expected = []string{"a", "b", "c"}
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}

	expected = []string{"e", "f", "g"}
	if !reflect.DeepEqual(expected, vals) {
		t.Errorf("expected %#v got %#v", expected, vals)
	}
}

func TestStack_Poke(t *testing.T) {
	s := NewStack[string]()
	s.vals = []string{"a", "b", "c"}
	s.Poke(1, "Z", "Y")
	expected := []string{"a", "Z", "Y", "b", "c"}
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
	s.Poke(0, "X")
	expected = append([]string{"X"}, expected...)
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
	s.Poke(-1, "W")
	expected = []string{"X", "a", "Z", "Y", "b", "W", "c"}
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
}

func TestStack_Peek(t *testing.T) {
	s := NewStack[string]()
	s.vals = []string{"a", "b", "c"}
	var expected string
	val := s.Peek(-4)
	if !reflect.DeepEqual(expected, val) {
		t.Errorf("expected %#v got %#v", expected, val)
	}
	expected = "c"
	val = s.Peek(-1)
	if !reflect.DeepEqual(expected, val) {
		t.Errorf("expected %#v got %#v", expected, val)
	}
}

func TestStack_Set(t *testing.T) {
	s := NewStack[string]()
	s.vals = []string{"a", "b", "c"}
	s.Set(-1, "d")
	expected := []string{"a", "b", "d"}
	if !reflect.DeepEqual(expected, s.vals) {
		t.Errorf("expected %#v got %#v", expected, s.vals)
	}
}
