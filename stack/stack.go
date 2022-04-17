package stack

import (
	"errors"
)

type Stack[T any] struct {
	storage []T
}

// Constructore and initialization
func (s *Stack[T]) Init(defaultCapasity ...int) *Stack[T] {
	defaultCapasity_ := 0
	if len(defaultCapasity) > 0 {
		defaultCapasity_ = defaultCapasity[0]
	}
	s.storage = make([]T, 0, defaultCapasity_)

	return s
}

func New[T any]() *Stack[T] {
	return new(Stack[T]).Init()
}

func (s *Stack[T]) Len() int {
	return len(s.storage)
}

func (s *Stack[T]) Push(el T) {
	s.storage = append(s.storage, el)
}

func (s *Stack[T]) Pop() (el T, err error) {
	i := len(s.storage) - 1
	if i < 0 {
		return *new(T), errors.New("pop() from empty stack")
	}
	r := s.storage[i]
	s.storage = s.storage[0:i]

	return r, nil
}

// peak()
func (s *Stack[T]) Peak() T {
	return s.storage[len(s.storage)-1]
}
