package undigraph

// 深度优先搜索找出图中所有连通分量
type CC struct {
	marked []bool
	Id     []int
	Count  int
}

func NewCC(g *Graph) *CC {
	cc := &CC{
		marked: make([]bool, g.V),
		Id:     make([]int, g.V),
	}
	for s := 0; s < g.V; s++ {
		if !cc.marked[s] {
			cc.dfs(g, s)
			cc.Count++
		}
	}
	return cc
}

func (cc *CC) dfs(g *Graph, v int) {
	cc.marked[v] = true
	cc.Id[v] = cc.Count
	for _, w := range g.Adj(v).Data() {
		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}

func (cc *CC) Connected(v, w int) bool {
	return cc.Id[v] == cc.Id[w]
}
