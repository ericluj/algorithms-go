package graph

// 计算强连通分量的Kosaraju算法
type KosarajuSCC struct {
	marked []bool // 已访问过的顶点
	id     []int  // 强连通分量的标识符
	count  int    // 强连通分量的数量
}

func NewKosarajuSCC(g *Digraph) *KosarajuSCC {
	scc := &KosarajuSCC{
		marked: make([]bool, g.V),
		id:     make([]int, g.V),
	}
	// 反向图的逆后序排列
	order := NewDepthFirstOrder(g.Reverse())
	for _, s := range order.reversePost.Data() {
		if !scc.marked[s] {
			scc.dfs(g, s)
			scc.count++
		}
	}
	return scc
}

func (scc *KosarajuSCC) dfs(g *Digraph, v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count
	for _, w := range g.Adj(v).Data() {
		if !scc.marked[w] {
			scc.dfs(g, w)
		}
	}
}

func (scc *KosarajuSCC) stronglyConnected(v, w int) bool {
	return scc.id[v] == scc.id[w]
}
