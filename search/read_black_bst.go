package search

const (
	Red   = true
	Black = false
)

// RedBlackBST 红黑树
type RedBlackBST struct {
	Root *RBNode
}

// NewRedBlackBST ...
func NewRedBlackBST() RedBlackBST {
	return RedBlackBST{}
}

// RBNode ...
type RBNode struct {
	Key   Key     //键
	Val   Val     //值
	Left  *RBNode //左子树
	Right *RBNode //右子树
	Num   int     //以该结点为根的子树中的结点总数
	Color bool    //其父结点指向它的链接的颜色
}

// IsRed ...
func (n *RBNode) IsRed() bool {
	if n == nil {
		return false
	}
	if n.Color {
		return Red
	} else {
		return Black
	}
}

// Size ...
func (n *RBNode) Size() int {
	if n == nil {
		return 0
	} else {
		return n.Num
	}
}

// 左旋 (右红链接会用到)
func rotateLeft(h *RBNode) *RBNode {
	x := h.Right                               //新的根结点是以前的右子结点
	h.Right = x.Left                           //介于之间的子结点需要交换
	x.Left = h                                 //新的根结点左子变为以前的根
	x.Color = h.Color                          //根结点颜色与之前一直
	h.Color = Red                              //左旋后左链接为红
	x.Num = h.Num                              //根结点下的子树结点树不会变
	h.Num = h.Left.Size() + h.Right.Size() + 1 //左结点子树num 正常形式计算即可
	return x
}

// 右旋 (左红链接会用到)
func rotateRight(h *RBNode) *RBNode {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = h.Color
	h.Color = Red
	x.Num = h.Num
	h.Num = h.Left.Size() + h.Right.Size() + 1
	return x
}

// 颜色转换 (父结点颜色变红，左右红链接颜色变黑)
func flipColors(h *RBNode) {
	h.Color = Red
	h.Left.Color = Black
	h.Right.Color = Black
}

// Put ...
func (b *RedBlackBST) Put(k Key, v Val) {
	b.Root = putRB(b.Root, k, v)
}

func putRB(h *RBNode, k Key, v Val) *RBNode {
	if h == nil { //标准的插入操作，和父结点用红链接相连
		return &RBNode{Key: k, Val: v, Num: 1, Color: Red}
	}
	cmp := k.CompareTo(h.Key)
	if cmp < 0 {
		h.Left = putRB(h.Left, k, v)
	} else if cmp > 0 {
		h.Right = putRB(h.Right, k, v)
	} else {
		h.Val = v
	}

	//如果右红链接 左旋
	if h.Right.IsRed() && !h.Left.IsRed() {
		h = rotateLeft(h)
	}
	//如果左红链接+左红链接 右旋
	if h.Left.IsRed() && h.Left.Left.IsRed() {
		h = rotateRight(h)
	}
	//如果左红链接+右红链接 转换颜色
	if h.Left.IsRed() && h.Right.IsRed() {
		flipColors(h)
	}
	h.Num = h.Left.Size() + h.Right.Size() + 1
	return h
}

func (b *RedBlackBST) Get(k Key) Val {
	return getRB(b.Root, k)
}

func getRB(n *RBNode, k Key) Val {
	//在以n为根结点的子树中查找并返回k对应的值
	//若找不到返回nil

	if n == nil {
		return nil
	}
	cmp := k.CompareTo(n.Key)
	if cmp < 0 {
		return getRB(n.Left, k)
	} else if cmp > 0 {
		return getRB(n.Right, k)
	} else {
		return n.Val
	}
}
