package string

type Value interface{}

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
