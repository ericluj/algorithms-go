package digraph

// 计算强连通分量的Kosaraju算法
type KosarajuSCC struct {
	marked []bool // 已访问过的顶点
	Id     []int  // 强连通分量的标识符
	Count  int    // 强连通分量的数量
}

func NewKosarajuSCC(g *Digraph) *KosarajuSCC {
	scc := &KosarajuSCC{
		marked: make([]bool, g.V),
		Id:     make([]int, g.V),
	}
	// 反向图的逆后序排列
	order := NewDepthFirstOrder(g.Reverse())
	// 按照上一步的顶点顺序进行dfs
	for _, s := range order.ReversePost.Data() {
		if !scc.marked[s] {
			scc.dfs(g, s)
			scc.Count++
		}
	}
	return scc
}

func (scc *KosarajuSCC) dfs(g *Digraph, v int) {
	scc.marked[v] = true
	scc.Id[v] = scc.Count
	for _, w := range g.Adj(v).Data() {
		if !scc.marked[w] {
			scc.dfs(g, w)
		}
	}
}

func (scc *KosarajuSCC) StronglyConnected(v, w int) bool {
	return scc.Id[v] == scc.Id[w]
}
