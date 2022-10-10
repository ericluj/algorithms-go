package search

// 基于线性探测法的散列表
type LinearProbingHashST struct {
	N    int    // 键值对总数
	M    int    // 线性探测表的大小
	Keys []*Key // 键
	Vals []*Val // 值
}

func NewLinearProbingHashST(m int) *LinearProbingHashST {
	hastST := &LinearProbingHashST{
		M:    m,
		Keys: make([]*Key, m),
		Vals: make([]*Val, m),
	}

	return hastST
}

func (s *LinearProbingHashST) hash(k Key) int {
	sum := 0
	for i := 0; i < len(k); i++ {
		sum += int(k[i])
	}
	return sum % s.M
}

func (s *LinearProbingHashST) Get(k Key) Val {
	for i := s.hash(k); s.Keys[i] != nil; i = (i + 1) % s.M {
		if k.CompareTo(*s.Keys[i]) == 0 {
			return s.Vals[i]
		}
	}

	return nil
}
