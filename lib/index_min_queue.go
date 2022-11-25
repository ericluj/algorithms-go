package lib

// 这里用map来方便实现，书上是用两个数组
type IndexMinQueue struct {
	data map[int]float64
}

func NewIndexMinQueue() *IndexMinQueue {
	return &IndexMinQueue{
		data: make(map[int]float64),
	}
}

func (q *IndexMinQueue) Insert(k int, x float64) {
	q.data[k] = x
}

func (q *IndexMinQueue) Change(k int, x float64) {
	q.data[k] = x
}

func (q *IndexMinQueue) Contains(k int) bool {
	_, ok := q.data[k]
	return ok
}

func (q *IndexMinQueue) DelMin() int {
	var (
		minK int
		min  float64
	)
	for k, v := range q.data {
		minK = k
		min = v
		break
	}
	for k, v := range q.data {
		if v < min {
			minK = k
			min = v
		}
	}
	delete(q.data, minK)
	return minK
}

func (q *IndexMinQueue) IsEmpty() bool {
	return len(q.data) == 0
}
