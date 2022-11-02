package lib

import "fmt"

// 栈 LIFO 后进先出
type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(x int) {
	s.data = append([]int{x}, s.data...)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	h := s.data[0]
	s.data = s.data[1:]
	return h
}

func (q *Stack) IsEmpty() bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.data)
}