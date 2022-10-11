package search

import "fmt"

// 基于线性探测法的散列表
type LinearProbingHashST struct {
	N    int   // 键值对总数
	M    int   // 线性探测表的大小
	Keys []Key // 键
	Vals []Val // 值
}

func NewLinearProbingHashST(m int) *LinearProbingHashST {
	hastST := &LinearProbingHashST{
		M:    m,
		Keys: make([]Key, m),
		Vals: make([]Val, m),
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
	for i := s.hash(k); s.Keys[i] != ""; i = (i + 1) % s.M {
		if k.CompareTo(s.Keys[i]) == 0 {
			return s.Vals[i]
		}
	}

	return nil
}

func (s *LinearProbingHashST) Put(k Key, v Val) {
	// 线性探测表大小调整
	if s.N >= (s.M / 2) {
		s.Resize(s.M * 2)
	}

	var i int
	for i = s.hash(k); s.Keys[i] != ""; i = (i + 1) % s.M {
		if k.CompareTo(s.Keys[i]) == 0 {
			s.Vals[i] = v
			return
		}
	}

	s.Keys[i] = k
	s.Vals[i] = v
	s.N++
}

func (s *LinearProbingHashST) Resize(m int) {
	t := NewLinearProbingHashST(m)
	for i := 0; i < s.M; i++ {
		if s.Keys[i] != "" {
			t.Put(s.Keys[i], s.Vals[i])
		}
	}
	s.Keys = t.Keys
	s.Vals = t.Vals
	s.M = t.M
}

func (s *LinearProbingHashST) String() string {
	str := ""
	for i := 0; i < s.M; i++ {
		if s.Keys[i] != "" {
			str += fmt.Sprintf("%s:%v ", s.Keys[i], s.Vals[i])
		} else {
			str += "- "
		}
	}
	return str
}

func (s *LinearProbingHashST) Contains(k Key) bool {
	for i := s.hash(k); s.Keys[i] != ""; i = (i + 1) % s.M {
		if k.CompareTo(s.Keys[i]) == 0 {
			return true
		}
	}

	return false
}

func (s *LinearProbingHashST) Delete(k Key) {
	// 不存在不需要删除
	if !s.Contains(k) {
		return
	}

	i := s.hash(k)
	for k.CompareTo(s.Keys[i]) != 0 {
		i = (i + 1) % s.M
	}
	s.Keys[i] = ""
	s.Vals[i] = nil

	// 被删除键右侧的所有键重新插入散列表
	i = (i + 1) % s.M
	for s.Keys[i] != "" {
		kToDo := s.Keys[i]
		vToDo := s.Vals[i]
		s.Keys[i] = ""
		s.Vals[i] = nil
		s.N--
		s.Put(kToDo, vToDo)
		i = (i + 1) % s.M
	}

	s.N--
	if (s.N > 0) && (s.N == s.M/8) {
		s.Resize(s.M / 2)
	}
}
