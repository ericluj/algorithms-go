package search

import "fmt"

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

// BST 二叉查找树
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
	//在以n为根结点的子树中查找并返回k对应的值
	//若找不到返回nil

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
	//如果k存在于以n为根结点的子树中则更新它的值
	//否则将以k,v为键值对的新结点插入到该子树中

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

func (b *BST) Min() Key {
	return min(b.Root).Key
}

func min(n *Node) *Node {
	if n.Left == nil {
		return n
	}
	return min(n.Left)
}

func (b *BST) Max() Key {
	return max(b.Root).Key
}

func max(n *Node) *Node {
	if n.Right == nil {
		return n
	}
	return max(n.Right)
}

// Floor 向下取整获取key
func (b *BST) Floor(k Key) Key {
	n := floor(b.Root, k)
	if n == nil {
		return ""
	}
	return n.Key
}

func floor(n *Node, k Key) *Node {
	if n == nil {
		return nil
	}
	cmp := k.CompareTo(n.Key)
	if cmp == 0 {
		return n
	} else if cmp < 0 {
		return floor(n.Left, k)
	}
	t := floor(n.Right, k)
	if t != nil {
		return t
	} else {
		return n
	}
}

// Select 寻找排名为k的Key(即树中正好有k个小于它的键)
func (b *BST) Select(k int) Key {
	return selectF(b.Root, k).Key
}

func selectF(n *Node, k int) *Node {
	if n == nil {
		return nil
	}
	t := n.Left.Size()
	if t > k {
		return selectF(n.Left, k)
	} else if t < k {
		return selectF(n.Right, k-t-1)
	} else {
		return n
	}
}

// DeleteMin ...
func (b *BST) DeleteMin() {
	b.Root = deleteMin(b.Root)
}

func deleteMin(n *Node) *Node {
	if n.Left == nil {
		return n.Right
	}
	n.Left = deleteMin(n.Left)
	n.Num = n.Left.Size() + n.Right.Size() + 1
	return n
}

// Delete ...
func (b *BST) Delete(k Key) {
	b.Root = deleteF(b.Root, k)
}

func deleteF(n *Node, k Key) *Node {
	if n == nil {
		return nil
	}
	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		n.Left = deleteF(n.Left, k)
	} else if cmp > 0 {
		n.Right = deleteF(n.Right, k)
	} else {
		if n.Right == nil {
			return n.Left
		}
		if n.Left == nil {
			return n.Right
		}
		t := n
		n = min(t.Right)
		n.Right = deleteMin(n.Right)
		n.Left = t.Left
	}
	n.Num = n.Left.Size() + n.Right.Size() + 1
	return n
}

func (b *BST) Print() {
	print(b.Root)
}

func print(n *Node) {
	if n == nil {
		return
	}
	print(n.Left)
	fmt.Println(n.Key)
	print(n.Right)
}
