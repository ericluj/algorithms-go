package search

// RedBlackBST 红黑二叉查找树
type RedBlackBST struct {
	Root *RBNode
}

const (
	Red   = true
	Black = false
)

type RBNode struct {
	Key   Key     // 键
	Val   Val     // 值
	Num   int     // 以该结点为根的子树中的结点总数（包括自己）
	Left  *RBNode // 左子结点
	Right *RBNode // 右子结点
	Color bool    // （父结点）指向当前结点的链接的颜色
}

func (n *RBNode) IsRed() bool {
	// 约定空结点为黑色
	if n == nil {
		return Black
	}
	return n.Color
}

func RBSize(node *RBNode) int {
	if node == nil {
		return 0
	}

	return node.Num
}

// 左旋（向右倾斜的红链接变为向左倾斜的红链接）
// n的颜色可红可黑，right为红，left为黑
// 左旋目的使left为红，right为黑
// 左旋前后遵循二叉树左小右大原则
func rotateLeft(n *RBNode) *RBNode {
	t := n.Right                                 // right变为新的n（临时保存为t）
	n.Right = t.Left                             // n的right本来为t，现在变为t的左结点（n~t之间的数）
	t.Left = n                                   // t的left本来为（n~t之间的数），现在变为n
	t.Color = n.Color                            // t继承了n的颜色（必须先继承颜色，才能设置n颜色为红，与下一行对调会出错）
	n.Color = Red                                // n变为红
	t.Num = n.Num                                // t继承了n的Num
	n.Num = 1 + RBSize(n.Left) + RBSize(n.Right) // 重新计算n的Num
	return t
}

// 右旋
func rotateRight(n *RBNode) *RBNode {
	t := n.Left
	n.Left = t.Right
	t.Right = n
	t.Color = n.Color
	n.Color = Red
	t.Num = n.Num
	n.Num = 1 + RBSize(n.Left) + RBSize(n.Right)
	return t
}

// 颜色转换（使中结点颜色变红，相当于将它送入了父结点）
// 两个子结点颜色由红变黑，父结点颜色由黑变红
func flipColors(n *RBNode) {
	n.Color = Red
	n.Left.Color = Black
	n.Right.Color = Black
}

func (r *RedBlackBST) Put(k Key, v Val) {
	// 查找key，找到更新值，否则为它新建一个结点
	r.Root = rbPut(r.Root, k, v)
	// 如果发生颜色转换，根结点可能变为红色，按照定义空链接为黑色，所以根结点肯定为黑色
	r.Root.Color = Black
}

func rbPut(n *RBNode, k Key, v Val) *RBNode {
	// 如果一直没有找到，递归到空链接，那么新建结点，和父结点red链接相连
	if n == nil {
		return &RBNode{Key: k, Val: v, Num: 1, Color: Red}
	}

	// 引入了新结点或者更新了值
	cmp := k.CompareTo(n.Key)
	if cmp < 0 { // 小于，沿着左子树一直找（找不到的话会在left新建结点）
		n.Left = rbPut(n.Left, k, v)
	} else if cmp > 0 { // 大于，沿着右子树一直找（找不到的话会在right新建结点）
		n.Right = rbPut(n.Right, k, v)
	} else if cmp == 0 { // 找到了更新值
		n.Val = v
	}

	// 判断当前结点本身是否需要有旋转或者颜色转换操作

	// 递归操作中，未引入新结点前树一直是符合红黑树定义
	// a.red链接均为左链接；
	// b.没有任何结点同时和两条red链接相连；
	// c.该树是完美黑色平衡，即任意空链接到根结点路径上black链接数量相同

	// 如果2结点插入新结点：
	//  a.left插入red结点
	//	 无需操作
	//  b.right插入red结点
	//	 左旋

	// 如果3结点插入新结点：
	//  a.middle插入red结点
	//	 左旋->右旋->颜色转换
	//  b.left插入red结点
	//	 右旋->颜色转换
	//  c.right插入red结点
	//	 颜色转换

	// 结合上述插入情况，我们只需要依次判断操作后的结点是否需要(左旋，右旋，颜色转换）即可
	// 步骤1：是否需要左旋（right红，left黑）
	if n.Right.IsRed() && !n.Left.IsRed() {
		n = rotateLeft(n)
	}
	// 步骤2：是否需要右旋（left红，left.left红）
	// 经历了步骤1，right肯定是黑，但可能导致两个红链接相连
	// 短路与，不会出现n.Left.Left的panic
	if n.Left.IsRed() && n.Left.Left.IsRed() {
		n = rotateRight(n)
	}
	// 步骤3：是否需要颜色转换（left红，right红）
	if n.Left.IsRed() && n.Right.IsRed() {
		flipColors(n)
	}

	n.Num = 1 + RBSize(n.Left) + RBSize(n.Right)
	return n
}
