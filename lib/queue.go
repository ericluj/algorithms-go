package lib

import "fmt"

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]int, 0),
	}
}

// 加入队列
func (q *Queue) Enqueue(x int) {
	q.data = append(q.data, x)
}

// 移除队列头
func (q *Queue) Dequeue(x int) int {
	if q.IsEmpty() {
		return 0
	}

	h := q.data[0]
	q.data = q.data[1:]
	return h
}

func (q *Queue) IsEmpty() bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}

func (q *Queue) String() string {
	return fmt.Sprintf("%v", q.data)
}
