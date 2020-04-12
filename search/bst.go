package search

type Key string

func (k Key) CompareTo(key Key) int {
	if k < key {
		return -1
	} else if k > key {
		return 1
	} else {
		return 0
	}
}

type Val interface{}

type BST struct {
	Root *Node
}

func NewBST() BST {
	return BST{}
}

type Node struct {
	Key   Key   //键
	Val   Val   //值
	Left  *Node //左子结点
	Right *Node //右子结点
	Num   int   //以该结点为根的子树中的结点总数
}

func (n *Node) Size() int {
	if n == nil {
		return 0
	} else {
		return n.Num
	}
}

func (b *BST) Get(k Key) Val {
	return get(b.Root, k)
}

func get(n *Node, k Key) Val {
	if n == nil {
		return nil
	}
	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		return get(n.Left, k)
	} else if cmp > 0 {
		return get(n.Right, k)
	} else {
		return n.Val
	}
}

func (b *BST) Put(k Key, v Val) {
	b.Root = put(b.Root, k, v)
}

func put(n *Node, k Key, v Val) *Node {
	if n == nil {
		return &Node{Key: k, Val: v, Num: 1}
	}
	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		n.Left = put(n.Left, k, v)
	} else if cmp > 0 {
		n.Right = put(n.Right, k, v)
	} else {
		n.Val = v
	}
	n.Num = n.Left.Size() + n.Right.Size() + 1
	return n
}
