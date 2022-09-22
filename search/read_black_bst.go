package search

// RedBlackBST 红黑二叉查找树
type RedBlackBST struct {
	Root *RBNode
}

func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
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

func IsRed(n *RBNode) bool {
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
// 左旋前后变化的只有结点和它的子树
func rotateLeft(n *RBNode) *RBNode {
	// 需要返回一个变化后的树，我们用t来表示

	// 构造根结点
	// 根结点由n变为n.Right
	t := n.Right

	// 构造根结点的左子结点
	// 左子结点为n且需要将（n~t之间的数）设为n的右子结点
	n.Right = t.Left
	t.Left = n

	// 构造根节点的右子结点
	// 不需要操作，因为根结点的右子结点就是t.Right

	// 根结点 t
	// 根结点左子结点 t.Letf => n
	// 根结点右子结点 t.Right

	// 设置根结点颜色
	// 继承了n的颜色（必须先继承颜色，才能设置n颜色为红，与下一行对调会出错）
	t.Color = n.Color

	// 设置根结点的左子结点的颜色
	// n变为红
	n.Color = Red

	// 设置根结点的右子结点的颜色
	// 不需要操作，右子结点中的关系没有变化过

	// 设置根结点Num
	// 继承了n的Num
	t.Num = n.Num

	// 设置根结点的左子结点的Num
	// 重新计算n的Num
	n.Num = 1 + RBSize(n.Left) + RBSize(n.Right)

	// 设置根结点的右子结点的Num
	// 不需要操作，右子结点中的关系没有变化过

	// 返回变化后的树
	return t
}

// 右旋（参考左旋）
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
	// 如果发生颜色转换，根结点可能变为红色，但它其实是一个2-结点，所以给黑色
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
	if IsRed(n.Right) && !IsRed(n.Left) {
		n = rotateLeft(n)
	}
	// 步骤2：是否需要右旋（left红，left.left红）
	// 经历了步骤1，right肯定是黑，但可能导致两个红链接相连
	// 短路与，不会出现n.Left.Left的panic
	if IsRed(n.Left) && IsRed(n.Left.Left) {
		n = rotateRight(n)
	}
	// 步骤3：是否需要颜色转换（left红，right红）
	if IsRed(n.Left) && IsRed(n.Right) {
		flipColors(n)
	}

	n.Num = 1 + RBSize(n.Left) + RBSize(n.Right)
	return n
}

func (r *RedBlackBST) Get(k Key) Val {
	return rbGet(r.Root, k)
}

func rbGet(n *RBNode, k Key) Val {
	if n == nil {
		return nil
	}

	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		return rbGet(n.Left, k)
	} else if cmp > 0 {
		return rbGet(n.Right, k)
	}

	// 相等直接返回
	return n.Val
}

// 删除结点颜色转换（中结点颜色变黑，相当于将其送入子结点）
func delFlipColors(n *RBNode) {
	n.Color = Black
	n.Left.Color = Red
	n.Right.Color = Red
}

// n的左子结点n.Left是2-结点，把它变化为
func moveRedLeft(n *RBNode) *RBNode {

	delFlipColors(n)

	if IsRed(n.Right.Left) {
		n.Right = rotateRight(n.Right)
		n = rotateLeft(n)
	}

	return n
}

func (r *RedBlackBST) DeleteMin() {
	// 根结点是2-结点，且它的左右子结点都是2-结点，将三个结点变成一个4-结点

	// 维护隐形不变量
	if !IsRed(r.Root.Left) && !IsRed(r.Root.Right) {
		r.Root.Color = Red
	}

	r.Root = rbDeleteMin(r.Root)
	if r.Root != nil {
		r.Root.Color = Black
	}
}

// 隐形不变量：n为红或者n.left为红
func rbDeleteMin(n *RBNode) *RBNode {
	// 删除的结点必须是红链接（这样才不影响平衡）

	// 当前结点最小，删除它
	if n.Left == nil {
		return nil
	}

	// 存在左子结点
	// n.Left是红色，可以直接删除
	// n.Left是黑色，但是存在红色的n.Left.Left（n.Left不需要被删除，删除的可能是n.Left.Left）
	if !IsRed(n.Left) && !IsRed(n.Left.Left) {
		n = moveRedLeft(n)
	}

	n.Left = rbDeleteMin(n.Left)

	return balance(n)
}

func balance(n *RBNode) *RBNode {
	if IsRed(n.Right) && !IsRed(n.Left) {
		n = rotateLeft(n)
	}
	if IsRed(n.Left) && IsRed(n.Left.Left) {
		n = rotateRight(n)
	}
	if IsRed(n.Left) && IsRed(n.Right) {
		delFlipColors(n)
	}

	n.Num = 1 + RBSize(n.Left) + RBSize(n.Right)
	return n
}
