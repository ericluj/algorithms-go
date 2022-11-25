package lib

import "fmt"

// 栈 LIFO 后进先出
type StackT interface {
	int | *DirectedEdge
}

type Stack[T StackT] struct {
	data []T
}

func NewStack[T StackT]() *Stack[T] {
	return &Stack[T]{
		data: make([]T, 0),
	}
}

func (s *Stack[T]) Push(x T) {
	s.data = append([]T{x}, s.data...)
}

func (s *Stack[T]) Pop() T {
	h := s.data[0]
	s.data = s.data[1:]
	return h
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v", s.data)
}

func (s *Stack[T]) Data() []T {
	return s.data
}
