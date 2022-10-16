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
	return false
}

func (uf *UF) Find(p int) int {
	return 0
}

func (uf *UF) Union(p, q int) {

}
