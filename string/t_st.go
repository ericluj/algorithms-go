package string

// 基于三向单词查找树的符号表
type TST struct {
	Root *TNode
}

func NewTST() *TST {
	return &TST{
		Root: nil,
	}
}

func (t *TST) Get(key string) Value {
	x := t.get(t.Root, key, 0)
	if x == nil {
		return nil
	}
	return x.Val
}

func (t *TST) get(x *TNode, key string, d int) *TNode {
	if x == nil {
		return nil
	}

	c := key[d]
	if c < x.C { // 没找到，小于沿着左子树找
		return t.get(x.Left, key, d)
	} else if c > x.C { // 没找到，大于沿着右子树找
		return t.get(x.Right, key, d)
	} else if d < len(key)-1 { // 找到了但没有到尾部，沿着中子树找下一个字符
		return t.get(x.Mid, key, d+1)
	} else { // 找到了且到尾部，返回结点
		return x
	}
}

func (t *TST) Put(key string, val Value) {
	t.Root = t.put(t.Root, key, val, 0)
}

func (t *TST) put(x *TNode, key string, val Value, d int) *TNode {
	c := key[d]
	if x == nil {
		x = NewTNode()
		x.C = c
	}

	if c < x.C { // 没找到，小于沿着左子树put
		x.Left = t.put(x.Left, key, val, d)
	} else if c > x.C { // 没找到，大于沿着右子树put
		x.Right = t.put(x.Right, key, val, d)
	} else if d < len(key)-1 { // 找到了但没有到尾部，沿着中子树put下一个字符
		x.Mid = t.put(x.Mid, key, val, d+1)
	} else { // 找到了且到尾部，设置val
		x.Val = val
	}

	return x
}
