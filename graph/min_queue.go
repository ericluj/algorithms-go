package graph

type MinQueue struct {
	data []*Edge
}

func NewMinQueue() *MinQueue {
	return &MinQueue{
		data: make([]*Edge, 0),
	}
}

// 加入队列
func (q *MinQueue) Enqueue(x *Edge) {
	q.data = append(q.data, x)
}

func (q *MinQueue) DelMin() *Edge {
	if q.IsEmpty() {
		return nil
	}

	index := 0
	min := q.data[0]
	for k, v := range q.data {
		if v.compareTo(min) < 0 {
			index = k
			min = v
		}
	}

	if index == len(q.data) { // 最小值为尾部
		q.data = q.data[:len(q.data)-1]
	} else {
		q.data = append(q.data[0:index], q.data[index+1:]...)
	}

	return min
}

func (q *MinQueue) IsEmpty() bool {
	return len(q.data) == 0
}
