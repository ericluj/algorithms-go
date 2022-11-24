package lib

import "fmt"

// 队列 FIFO 先进先出
type QueueT interface {
	int | *Edge
}

type Queue[T QueueT] struct {
	data []T
}

func NewQueue[T QueueT]() *Queue[T] {
	return &Queue[T]{
		data: make([]T, 0),
	}
}

// 加入队列
func (q *Queue[T]) Enqueue(x T) {
	q.data = append(q.data, x)
}

// 移除队列头
func (q *Queue[T]) Dequeue() T {
	h := q.data[0]
	q.data = q.data[1:]
	return h
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.data)
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", q.data)
}
