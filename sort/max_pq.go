package sort

// 基于堆的优先队列
// 当一棵二叉树的每个结点都大于等于它的两个字结点时，称为堆有序
type MaxPQ struct {
	pq []int // 基于堆的完全二叉树
	N  int   // 堆的大小
}

func NewMaxPQ(maxN int) *MaxPQ {
	return &MaxPQ{
		pq: make([]int, maxN),
	}
}

func (p *MaxPQ) IsEmpty() bool {
	return p.N == 0
}

func (p *MaxPQ) Insert(k int) {

}

func (p *MaxPQ) Swim(k int) {

}

func (p *MaxPQ) Sink(k int) {

}
