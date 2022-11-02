package graph

// 深度优先搜索找出图中所有连通分量
type CC struct {
	marked []bool
	id     []int
	count  int
}

func NewCC(g *Graph) *CC {
	cc := &CC{
		marked: make([]bool, g.V),
		id:     make([]int, g.V),
	}
	for s := 0; s < g.V; s++ {
		if !cc.marked[s] {
			cc.dfs(g, s)
			cc.count++
		}
	}
	return cc
}

func (cc *CC) dfs(g *Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, w := range g.Adj(v).Data() {
		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}

func (cc *CC) connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}
