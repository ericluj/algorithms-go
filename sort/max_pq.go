package sort

import "github.com/ericluj/algorithms-go/lib"

// 基于堆的优先队列
// 当一棵二叉树的每个结点都大于等于它的两个子结点时，称为堆有序
// 数组表示二叉堆：数据中按照层级存储（索引为0的位置不使用）
// 位置k的结点的父结点的位置为k/2，它的两个子结点位置为2k和2k+1
type MaxPQ struct {
	pq []int // 基于堆的完全二叉树
	N  int   // 存储于pq[1..N]中，pq[0]没有使用
}

func NewMaxPQ(maxN int) *MaxPQ {
	return &MaxPQ{
		pq: make([]int, maxN+1),
	}
}

func (p *MaxPQ) IsEmpty() bool {
	return p.N == 0
}

// 插入元素：新元素加到数组末尾，上浮新元素到合适位置
func (p *MaxPQ) Insert(v int) {
	// 添加新元素到末尾
	p.N++
	p.pq[p.N] = v
	// 恢复堆的有序性
	p.Swim(p.N)
}

// 删除最大元素：数组顶端删除最大元素，将数组最后一个元素放到顶端，下沉它到合适位置
func (p *MaxPQ) DelMax() int {
	max := p.pq[1]         // 从根结点得到最大元素
	lib.Exch(p.pq, 1, p.N) // 将最后一个元素放到顶端
	p.N--                  // 删除元素（并没有真正删除，N控制数组边界）
	p.Sink(1)              // 恢复堆的有序性
	return max
}

// 由下至上的堆有序化（上浮）
func (p *MaxPQ) Swim(k int) {
	// 如果结点k不是根结点（k==1为根结点）且结点k大于它的父结点k/2
	// 将结点k与父结点k/2交换位置
	for (k > 1) && (p.pq[k] > p.pq[k/2]) {
		lib.Exch(p.pq, k/2, k)
		k = k / 2
	}
}

// 由上至下的堆有序化（下沉）
func (p *MaxPQ) Sink(k int) {
	// 如果结点k不是叶子结点2*k<=p.N
	for 2*k <= p.N {
		j := 2 * k // 左子结点

		// 如果左右子结点存在，找到子结点中较大的结点
		// 当j==p.N说明没有右子结点
		if (j < p.N) && (p.pq[j+1] > p.pq[j]) {
			j++
		}
		// 如果结点k大于它的两个子结点，无需下沉
		if p.pq[k] >= p.pq[j] {
			break
		}
		// 将结点k与子结点中较大的结点交换位置
		lib.Exch(p.pq, k, j)
		k = j
	}
}
