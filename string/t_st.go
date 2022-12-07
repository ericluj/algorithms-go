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

func (t *TST) Delete(key string) {
	t.Root = t.delete(t.Root, key, 0)
}

func (t *TST) delete(x *TNode, key string, d int) *TNode {
	if x == nil {
		return nil
	}

	c := key[d]
	if c < x.C { // 没找到，小于沿着左子树
		x.Left = t.delete(x.Left, key, d)
	} else if c > x.C { // 没找到，大于沿着右子树
		x.Right = t.delete(x.Right, key, d)
	} else if d < len(key)-1 { // 找到了但没有到尾部
		x.Mid = t.delete(x.Mid, key, d+1)
	} else { // 找到了且到尾部，设置val
		x.Val = nil
		// // 中子树存在，不能删除当前结点，直接返回
		// if x.Mid != nil {
		// 	return x
		// }
		// // 中子树不存在，可以删除
		// if x.Left == nil {
		// 	return x.Right
		// }
		// if x.Right == nil {
		// 	return x.Left
		// }
		// // 中子树不存在，左右结点存在时，使用有子树最左边结点替换x
		// // （这里与二叉树的delete原理相同）
		// tmp := x
		// x := t.min(tmp.Right)
		// x.Right = t.delMin(tmp.Right)
		// x.Left = tmp.Left
	}

	// 递归返回时删除没有中子树，并且当前结点值为nil的结点
	// 没有中子结点时，相当于是一颗二叉树了
	// 这里与二叉树的delete原理相同
	if x.Mid == nil && x.Val == nil {
		// 如果是单子结点
		if x.Left == nil {
			return x.Right
		}
		if x.Right == nil {
			return x.Left
		}
		// 如果是双子结点
		tmp := x
		x := t.min(tmp.Right)
		x.Right = t.delMin(tmp.Right)
		x.Left = tmp.Left
	}

	return x
}

func (t *TST) min(x *TNode) *TNode {
	if x.Left == nil {
		return x
	}
	return t.min(x.Left)
}

func (t *TST) delMin(x *TNode) *TNode {
	if x.Left == nil {
		return x.Right
	}
	x.Left = t.delMin(x.Left)
	return x
}
