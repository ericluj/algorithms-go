package string

type Value interface{}

// 单词查找树的结点
type Node struct {
	Val  Value
	Next []*Node
}

func NewNode(R int, v Value) *Node {
	return &Node{
		Val:  v,
		Next: make([]*Node, R),
	}
}
