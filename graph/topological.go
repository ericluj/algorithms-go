package graph

import "github.com/ericluj/algorithms-go/lib"

type Topological struct {
	order *lib.Stack // 顶点的拓扑顺序
}

func NewTopological(g *Digraph) *Topological {
	// 优先级限制下的调度问题等价于计算有向无环图中的所有顶点的拓扑排序
	// 一副有向无环图的拓扑排序即为所有顶点的逆后序排列
	t := &Topological{}
	cycleFinder := NewDirectedCycle(g)
	if !cycleFinder.hasCycle() {
		dfo := NewDepthFirstOrder(g)
		t.order = dfo.reversePost
	}
	return t
}