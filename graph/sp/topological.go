package sp

import "github.com/ericluj/algorithms-go/lib"

type Topological struct {
	Order *lib.Stack[int] // 顶点的拓扑顺序
}

// 拓扑排序：给定一副有向图，将所有顶点排序，使得所有边均从排在前面的元素指向排在后面的元素（或者说明无法做到这一点）
func NewTopological(g *EdgeWeightedDigraph) *Topological {
	// 优先级限制下的调度问题等价于计算有向无环图中的所有顶点的拓扑排序
	// 一副有向无环图的拓扑排序即为所有顶点的逆后序排列
	t := &Topological{}
	cycleFinder := NewDirectedCycle(g)
	if !cycleFinder.HasCycle() {
		dfo := NewDepthFirstOrder(g)
		t.Order = dfo.ReversePost
	}
	return t
}
