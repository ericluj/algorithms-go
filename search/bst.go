package search

// BST 二叉查找树
type BST struct {
	Root *Node
}

type Node struct {
	Key   Key   //键
	Val   Val   //值
	Num   int   //以该结点为根的子树中的结点总数（包括自己）
	Left  *Node //左子结点
	Right *Node //右子结点
}

func NewBST() *BST {
	return &BST{}
}

func size(node *Node) int {
	if node == nil {
		return 0
	}

	return node.Num
}

func (bst *BST) Get(k Key) Val {
	return get(bst.Root, k)
}

func get(node *Node, k Key) Val {
	// 找不到返回nil
	if node == nil {
		return nil
	}

	cmp := k.CompareTo(node.Key)
	if cmp < 0 {
		return get(node.Left, k)
	} else if cmp > 0 {
		return get(node.Right, k)
	}

	// 相等直接返回
	return node.Val
}

func (bst *BST) Put(k Key, v Val) {
	bst.Root = put(bst.Root, k, v)
}

func put(node *Node, k Key, v Val) *Node {
	// 没有找到那么返回一个新结点，这个新结点会被连接到最后一个查找的结点上
	if node == nil {
		return &Node{Key: k, Val: v, Num: 1}
	}

	cmp := k.CompareTo(node.Key)
	if cmp < 0 {
		// 没有找到返回新结点，新结点连接在left
		node.Left = put(node.Left, k, v)
	} else if cmp > 0 {
		// 没有找到返回新结点，新结点连接在right
		node.Right = put(node.Right, k, v)
	} else {
		// 找到了结点，更新值
		node.Val = v
	}

	// 插入新结点或者更新了值，重新计算num
	node.Num = size(node.Left) + size(node.Right) + 1
	return node
}
