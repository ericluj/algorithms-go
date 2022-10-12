package graph

type Bag struct {
	data []int
}

func NewBag() *Bag {
	return &Bag{
		data: make([]int, 0),
	}
}

func (b *Bag) Add(v int) {
	b.data = append([]int{v}, b.data...)
}
