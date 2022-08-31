package search

import "fmt"

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

func (bst *BST) Print() {
	silce := make([]Key, 0)
	print(&silce, bst.Root)
	fmt.Println(silce)
}

func print(silce *[]Key, node *Node) {
	if node.Left != nil {
		print(silce, node.Left)
	}
	*silce = append(*silce, node.Key)
	if node.Right != nil {
		print(silce, node.Right)
	}
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

// 获取小于等于Key的最大值
func (bst *BST) Floor(k Key) Key {
	node := floor(bst.Root, k)
	if node == nil {
		return ""
	}
	return node.Key
}

func floor(node *Node, k Key) *Node {
	if node == nil {
		return nil
	}
	cmp := k.CompareTo(node.Key)

	// 相等直接返回
	if cmp == 0 {
		return node
	}

	// 小于的话那么，最接近的值在左子树
	if cmp < 0 {
		return floor(node.Left, k)
	}

	// 大于的话要判断右子树有没有小于k的值
	tmp := floor(node.Right, k)
	// 如果有，那么就是它
	if tmp != nil {
		return tmp
	}
	// 如果没有，那么就是当前node本身
	return node
}

// 从小往大排名，返回排名第num的Key（从0开始排名）
func (bst *BST) Select(num int) Key {
	node := selectFuc(bst.Root, num)
	if node == nil {
		return ""
	}
	return node.Key
}

func selectFuc(node *Node, n int) *Node {
	if node == nil {
		return nil
	}
	t := size(node.Left) // 注意这里是左子树
	if t == n {          // 相等正好是第n名
		return node
	}
	if t > n { // 如果左子树的数量大于n，那么排名第n肯定在左子树中
		return selectFuc(node.Left, n)
	}
	if t < n { // 如果左子树的数量小于n，继续在右子树找到去掉（左子树数量+根结点1）后的排名
		return selectFuc(node.Right, n-t-1)
	}
	return nil
}

// 从小往大排名，返回排名第num的Key（从1开始排名）
func (bst *BST) SelectFromOne(n int) Key {
	node := selectFromOne(bst.Root, n)
	if node == nil {
		return ""
	}
	return node.Key
}

func selectFromOne(node *Node, n int) *Node {
	if node == nil {
		return nil
	}
	t := size(node.Left) + 1 // 注意这里是左子树数量+根结点1
	if t == n {
		return node
	}
	if t > n {
		return selectFromOne(node.Left, n)
	}
	if t < n {
		return selectFromOne(node.Right, n-t)
	}
	return nil
}

// 获取Key的排名（从0开始）
func (bst *BST) Rank(k Key) int {
	return rank(bst.Root, k)
}

func rank(node *Node, k Key) int {
	if node == nil {
		return -1
	}

	cmp := k.CompareTo(node.Key)
	if cmp == 0 {
		return size(node.Left)
	} else if cmp < 0 {
		return rank(node.Left, k)
	} else if cmp > 0 {
		return size(node.Left) + 1 + rank(node.Right, k)
	}

	//不会走到这里
	return -1
}

func (bst *BST) Min() *Node {
	return min(bst.Root)
}

func min(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return min(node.Left)
}

func (bst *BST) DeleteMin() {
	bst.Root = deleteMin(bst.Root)
}

func deleteMin(node *Node) *Node {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMin(node.Left)
	node.Num = size(node.Left) + size(node.Right) + 1
	return node
}

func (bst *BST) Delete(k Key) {
	bst.Root = delete(bst.Root, k)
}

func delete(node *Node, k Key) *Node {
	if node == nil {
		return nil
	}

	cmp := k.CompareTo(node.Key)
	if cmp < 0 { // 接着在左子树删除
		node.Left = delete(node.Left, k)
	} else if cmp > 0 { // 接着在右子树删除
		node.Right = delete(node.Right, k)
	} else if cmp == 0 { // 删除的就是当前结点
		// 如果是单子结点
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// 如果是双子结点
		// 创建新结点，替换原来的node
		new := min(node.Right)
		new.Right = deleteMin(node.Right)
		new.Left = node.Left
		node = new
	}
	node.Num = size(node.Left) + size(node.Right) + 1

	return node
}
