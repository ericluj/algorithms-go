package digraph

import (
	"github.com/ericluj/algorithms-go/lib"
)

// 有向图中基于深度优先搜索的顶点排序
type DepthFirstOrder struct {
	marked      []bool
	Pre         *lib.Queue[int] // 所有顶点的前序排列（dfs的调用顺序）
	Post        *lib.Queue[int] // 所有顶点的后序排列（顶点遍历完成的顺序）
	ReversePost *lib.Stack      // 所有顶点的逆后序排列（顶点遍历完成的顺序的逆序：最后完成的顶点最先出栈）
}

func NewDepthFirstOrder(g *Digraph) *DepthFirstOrder {
	dfo := &DepthFirstOrder{
		marked:      make([]bool, g.V),
		Pre:         lib.NewQueue[int](),
		Post:        lib.NewQueue[int](),
		ReversePost: lib.NewStack(),
	}
	for v := 0; v < g.V; v++ {
		if !dfo.marked[v] {
			dfo.dfs(g, v)
		}
	}
	return dfo
}

func (d *DepthFirstOrder) dfs(g *Digraph, v int) {
	d.Pre.Enqueue(v)

	d.marked[v] = true
	for _, w := range g.Adj(v).Data() {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}

	d.Post.Enqueue(v)
	d.ReversePost.Push(v)
}
