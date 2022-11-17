package graph

import "github.com/ericluj/algorithms-go/lib"

// 有向图中基于深度优先搜索的顶点排序
type DepthFirstOrder struct {
	marked      []bool
	pre         *lib.Queue // 所有顶点的前序排列（dfs的调用顺序）
	post        *lib.Queue // 所有顶点的后序排列（顶点遍历完成的顺序）
	reversePost *lib.Stack // 所有顶点的逆后序排列（顶点遍历完成的顺序的逆序：最后完成的顶点最先出栈）
}

func NewDepthFirstOrder(g *Digraph) *DepthFirstOrder {
	dfo := &DepthFirstOrder{
		marked:      make([]bool, g.V),
		pre:         lib.NewQueue(),
		post:        lib.NewQueue(),
		reversePost: lib.NewStack(),
	}
	for v := 0; v < g.V; v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}
	return dfo
}

func (d *DepthFirstOrder) dfs(g *Digraph, v int) {
	d.pre.Enqueue(v)

	d.marked[v] = true
	for _, w := range g.Adj(v).Data() {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}

	d.post.Enqueue(v)
	d.reversePost.Push(v)
}
