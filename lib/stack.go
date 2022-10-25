package lib

import "fmt"

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

func (s *Stack) String() string {
	return fmt.Sprintf("%v", s.data)
}
