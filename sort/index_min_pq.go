package sort

// 关联索引的优先队列
type IndexMinPQ struct {
	pq   []int     // 基于堆的完全二叉树，存储于pq[1..N]中，pq[0]没有使用，i是按照堆排序
	N    int       // 队列元素数量
	qp   []int     // qp[i]表示i存储在pq中的索引，目的是为了快速找到i元素在堆中的位置
	keys []float64 // 权重
}

func NewIndexMinPQ(maxN int) *IndexMinPQ {
	p := &IndexMinPQ{
		pq:   make([]int, maxN+1),
		qp:   make([]int, maxN+1),
		keys: make([]float64, maxN+1),
	}
	for i := 0; i < maxN; i++ {
		p.qp[i] = -1
	}
	return p
}

func (p *IndexMinPQ) IsEmpty() bool {
	return p.N == 0
}

func (p *IndexMinPQ) Contains(i int) bool {
	return p.qp[i] != -1
}

// 插入元素：新元素加到数组末尾，上浮新元素到合适位置
func (p *IndexMinPQ) Insert(i int, key float64) {
	if p.Contains(i) {
		return
	}
	// 添加新元素到末尾
	p.N++
	p.pq[p.N] = i // i放入堆中
	p.qp[i] = p.N // qp[i]表示i在堆中的索引
	p.keys[i] = key

	// 恢复堆的有序性
	p.Swim(p.N)
}

// 删除最小元素：数组顶端删除最小元素，将数组最后一个元素放到顶端，下沉它到合适位置
func (p *IndexMinPQ) DelMin() int {
	min := p.pq[1] // 从根结点得到最小元素
	p.Exch(1, p.N) // 将最后一个元素放到顶端
	p.N--          // 删除元素（并没有真正删除，N控制数组边界）
	p.Sink(1)      // 恢复堆的有序性

	p.qp[min] = -1
	p.keys[min] = -1
	p.pq[p.N+1] = -1

	return min
}

func (p *IndexMinPQ) Change(i int, key float64) {
	if !p.Contains(i) {
		return
	}
	p.keys[i] = key
	p.Swim(p.qp[i])
	p.Sink(p.qp[i])
}

func (p *IndexMinPQ) Delete(i int) {
	if !p.Contains(i) {
		return
	}
	index := p.qp[i] // 找到i在堆中的位置
	p.Exch(index, p.N)
	p.N--
	p.Swim(index)
	p.Sink(index)

	p.qp[i] = -1
	p.keys[i] = -1
}

// 由下至上的堆有序化（上浮）
func (p *IndexMinPQ) Swim(k int) {
	// 如果结点k不是根结点（k==1为根结点）且结点k小于它的父结点k/2
	// 将结点k与父结点k/2交换位置
	for (k > 1) && (p.keys[p.pq[k]] < p.keys[p.pq[k/2]]) {
		p.Exch(k/2, k)
		k = k / 2
	}
}

// 由上至下的堆有序化（下沉）
func (p *IndexMinPQ) Sink(k int) {
	// 如果结点k不是叶子结点2*k<=p.N
	for 2*k <= p.N {
		j := 2 * k // 左子结点

		// 如果左右子结点存在，找到子结点中较小的结点
		// 当j==p.N说明没有右子结点
		if (j < p.N) && (p.keys[p.pq[j+1]] < p.keys[p.pq[j]]) {
			j++
		}
		// 如果结点k小于它的两个子结点，无需下沉
		if p.keys[p.pq[k]] < p.keys[p.pq[j]] {
			break
		}
		// 将结点k与子结点中较大的结点交换位置
		p.Exch(k, j)
		k = j
	}
}

func (p *IndexMinPQ) Exch(i, j int) {
	p.pq[i], p.pq[j] = p.pq[j], p.pq[i]
	p.qp[p.pq[i]] = i
	p.qp[p.pq[j]] = j
}
