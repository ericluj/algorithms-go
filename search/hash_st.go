package search

// 基于拉链法的散列表
type SeparateChainingHashST struct {
	N  int    // 键值对总数
	M  int    // 散列表的大小
	st []*SST // 存放链表对象的数组
}

func NewSeparateChainingHashST(m int) *SeparateChainingHashST {
	hastST := &SeparateChainingHashST{
		M:  m,
		st: make([]*SST, m),
	}
	for i := 0; i < m; i++ {
		hastST.st[i] = NewSST()
	}

	return hastST
}

func (s *SeparateChainingHashST) hash(k Key) int {
	sum := 0
	for i := 0; i < len(k); i++ {
		sum += int(k[i])
	}
	return sum % s.M
}

func (s *SeparateChainingHashST) Get(k Key) Val {
	return s.st[s.hash(k)].Get(k)
}

func (s *SeparateChainingHashST) Put(k Key, v Val) {
	s.st[s.hash(k)].Put(k, v)
}
