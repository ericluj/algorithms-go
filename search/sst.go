package search

import (
	"fmt"
	"strings"
)

type SNode struct {
	k    Key
	v    Val
	next *SNode
}

func NewSNode(k Key, v Val, next *SNode) *SNode {
	return &SNode{
		k:    k,
		v:    v,
		next: next,
	}
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
	s.first = NewSNode(k, v, s.first)
}

func (s *SST) String() string {
	arr := make([]string, 0)
	for n := s.first; n != nil; n = n.next {
		arr = append(arr, fmt.Sprintf("%s:%v", n.k, n.v))
	}
	return strings.Join(arr, ",")
}
