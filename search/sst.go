package search

import (
	"fmt"
	"strings"
)

// 顺序查找（基于无序链表）
type SNode struct {
	k    Key
	v    Val
	next *SNode
}

type SST struct {
	ST
	first *SNode
}

func NewSST() *SST {
	return &SST{
		first: nil,
	}
}

func (s *SST) Get(k Key) Val {
	for n := s.first; n != nil; n = n.next {
		if k.CompareTo(n.k) == 0 {
			return n.v
		}
	}
	return nil
}

func (s *SST) Put(k Key, v Val) {
	for n := s.first; n != nil; n = n.next {
		if k.CompareTo(n.k) == 0 {
			n.v = v
			return
		}
	}
	s.first = &SNode{
		k:    k,
		v:    v,
		next: s.first,
	}
}

func (s *SST) Delete(k Key) {
	// 如果第一个元素就是要删除的元素
	if k.CompareTo(s.first.k) == 0 {
		s.first = s.first.next
		return
	}

	pre := s.first
	for pre.next != nil {
		if k.CompareTo(pre.next.k) == 0 {
			pre.next = pre.next.next
			return
		}
		pre = pre.next
	}
}

func (s *SST) String() string {
	arr := make([]string, 0)
	for n := s.first; n != nil; n = n.next {
		arr = append(arr, fmt.Sprintf("%s:%v", n.k, n.v))
	}
	return strings.Join(arr, ",")
}
