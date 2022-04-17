package stack

import "testing"

func TestStack(t *testing.T) {
	s := New[int]()
	testInt := 5
	s.Push(testInt)

	if l := s.Len(); l != 1 {
		t.Errorf("l.Len() = %d, want %d", l, 1)
	}

	if el, err := s.Pop(); err != nil || (el != testInt) {
		t.Errorf("s.Pop() = %d, expected = %d, err = %s", el, testInt, err)
	}

	if v, err := s.Pop(); err == nil {
		t.Errorf("s.Pop() = %d, err = %s; expected empty stack error", v, err)
	}
}
