package string

import "github.com/ericluj/algorithms-go/lib"

// 基于单词查找树的符号表
type TrieST struct {
	R    int   // 基数
	Root *Node // 单词查找树的根结点
}

func NewTrieST() *TrieST {
	R := 256
	return &TrieST{
		R:    R,
		Root: NewNode(R, nil),
	}
}

func (t *TrieST) Get(key string) Value {
	x := t.get(t.Root, key, 0)
	if x == nil {
		return nil
	}
	return x.Val
}

// 返回以x作为根结点的子单词查找树中与key相关联的值
func (t *TrieST) get(x *Node, key string, d int) *Node {
	if x == nil {
		return nil
	}
	// 已经到尾部了，返回子单词树本身
	if d == len(key) {
		return x
	}
	// 找到第d个字符所对应的子单词查找树
	c := key[d]
	return t.get(x.Next[c], key, d+1)
}

func (t *TrieST) Put(key string, val Value) {
	t.Root = t.put(t.Root, key, val, 0)
}

func (t *TrieST) put(x *Node, key string, val Value, d int) *Node {
	if x == nil {
		x = NewNode(t.R, nil)
	}
	// 已经到尾部了，修改子单词树本身的值
	if d == len(key) {
		x.Val = val
		return x
	}
	// 找到第d个字符所对应的子单词查找树
	c := key[d]
	x.Next[c] = t.put(x.Next[c], key, val, d+1)
	return x
}

func (t *TrieST) Keys() *lib.Queue[string] {
	return t.KeysPrefix("")
}

func (t *TrieST) KeysPrefix(pre string) *lib.Queue[string] {
	q := lib.NewQueue[string]()
	t.collect(t.get(t.Root, pre, 0), pre, q)
	return q
}

func (t *TrieST) collect(x *Node, pre string, q *lib.Queue[string]) {
	if x == nil {
		return
	}
	if x.Val != nil {
		q.Enqueue(pre)
	}
	for c := 0; c < t.R; c++ {
		t.collect(x.Next[c], pre+string(byte(c)), q)
	}
}

// 通配符匹配（不需要考虑长度超过模式字符串的键）
func (t *TrieST) KeysThatMatch(pat string) *lib.Queue[string] {
	q := lib.NewQueue[string]()
	t.collectPat(t.Root, "", pat, q)
	return q
}

func (t *TrieST) collectPat(x *Node, pre, pat string, q *lib.Queue[string]) {
	if x == nil {
		return
	}

	// 模式字符串已经匹配到尾部，结束递归
	if len(pre) == len(pat) {
		if x.Val != nil { // 如果匹配的单词存在则加入队列
			q.Enqueue(pre)
		}
		return
	}

	next := pat[len(pre)] // 按照递归的pre对应去拿模式字符串下一个字符
	for c := 0; c < t.R; c++ {
		if next == byte('.') || next == byte(c) { // 如果匹配到递归单词查找树
			t.collectPat(x.Next[c], pre+string(byte(c)), pat, q)
		}
	}
}

// 找到给定字符串的最长键前缀
func (t *TrieST) LongestPrefixOf(s string) string {
	length := t.search(t.Root, s, 0, 0)
	return s[:length]
}

func (t *TrieST) search(x *Node, s string, d, length int) int {
	if x == nil {
		return length
	}
	// 找到树中结束的单词，那么将它的length往下传
	// 如果找到更长的则替换length
	if x.Val != nil {
		length = d
	}
	// 给定字符串已经到尾部
	if d == len(s) {
		return length
	}
	c := s[d]
	return t.search(x.Next[c], s, d+1, length)
}

func (t *TrieST) Delete(key string) {
	t.delete(t.Root, key, 0)
}

func (t *TrieST) delete(x *Node, key string, d int) *Node {
	if x == nil {
		return nil
	}

	// 找到键对应的结点置为nil
	if d == len(key) {
		x.Val = nil
	} else { // 还没找到则继续寻找
		c := key[d]
		// 递归删除，如果删除子节点使父节点所有链接也为空，继续删除父节点
		x.Next[c] = t.delete(x.Next[c], key, d+1)
	}

	// 递归的操作中，如果结点x没有被置nil直接返回即可
	if x.Val != nil {
		return x
	}

	// 递归操作中，被置nil的结点x，检查它是否还有指向的单词子树
	// 如果有，结点x还有存在的必要
	// 如果没有，结点x不需要存在，置为nil
	for c := 0; c < t.R; c++ {
		if x.Next[c] != nil {
			return x
		}
	}
	return nil
}
