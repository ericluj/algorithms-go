package base

type UF struct {
	id    []int // 分量id（以触点作为索引）
	count int   // 分量数量
}

func NewUF(n int) *UF {
	uf := &UF{
		id:    make([]int, n),
		count: n,
	}
	for i := 0; i < n; i++ {
		uf.id[i] = i
	}
	return uf
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

// quick-find
func (uf *UF) Find(p int) int {
	return uf.id[p]
}

// quick-find
// 将p和q归并到相同的分量中
func (uf *UF) Union(p, q int) {
	pID := uf.Find(p)
	qID := uf.Find(q)

	// p和q已经在相同的分量中
	if pID == qID {
		return
	}

	// 将p的分量重命名为q的名称
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}

	uf.count--
}

// quick-union
func (uf *UF) FindU(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

// quick-union
func (uf *UF) UnionU(p, q int) {
	pRoot := uf.FindU(p)
	qRoot := uf.FindU(q)

	if pRoot == qRoot {
		return
	}

	uf.id[pRoot] = qRoot
	uf.count--
}
