package string

// 三向单词查找树的结点
type TNode struct {
	C     byte   // 字符
	Val   Value  // 和字符串相关联的值
	Left  *TNode // 左子三向单词查找树
	Mid   *TNode // 中子三向单词查找树
	Right *TNode // 右子三向单词查找树
}

func NewTNode() *TNode {
	return &TNode{}
}
