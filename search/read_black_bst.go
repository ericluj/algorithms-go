package search

// RedBlackBST 红黑二叉查找树
type RedBlackBST struct {
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
