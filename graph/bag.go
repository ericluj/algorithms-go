package graph

// Bag数据结构，使用了范型
type BagT interface {
	int | *Edge
}

type Bag[T BagT] struct {
	data []T
}

func NewBag[T BagT]() *Bag[T] {
	return &Bag[T]{
		data: make([]T, 0),
	}
}

func (b *Bag[T]) Add(v T) {
	b.data = append([]T{v}, b.data...)
}

func (b *Bag[T]) Len() int {
	return len(b.data)
}

func (b *Bag[T]) Data() []T {
	return b.data
}
