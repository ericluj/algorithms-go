package sp

import (
	"math"

	"github.com/ericluj/algorithms-go/lib"
)

// 无环加权有向图的最长路径算法
type AcyclicLP struct {
	EdgeTo []*lib.DirectedEdge // 最长路径树
	DistTo []float64           // 从顶点s到v的最短距离
}

func NewAcyclicLP(g *EdgeWeightedDigraph, s int) *AcyclicLP {
	d := &AcyclicLP{
		EdgeTo: make([]*lib.DirectedEdge, g.V),
		DistTo: make([]float64, g.V),
	}

	// 初始每个顶点给最小值，这样第一条边一定大于会被放入
	for v := 0; v < g.V; v++ {
		d.DistTo[v] = math.SmallestNonzeroFloat64
	}
	// 起点给0
	d.DistTo[s] = 0.0

	top := NewTopological(g)
	for _, v := range top.Order.Data() {
		// 这里没有加判断，所以起点实际上固定为拓扑排序的第一个值
		d.RelaxV(g, v)
	}

	return d
}

// 顶点的松弛
func (d *AcyclicLP) RelaxV(g *EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v).Data() {
		w := e.To()
		// 若边e引入大于之前最大的值，则e属于最长路径
		if d.DistTo[w] < d.DistTo[v]+e.Weight {
			d.DistTo[w] = d.DistTo[v] + e.Weight
			d.EdgeTo[w] = e
		}
	}
}

func (d *AcyclicLP) HasPathTo(v int) bool {
	return d.DistTo[v] < math.MaxFloat64
}

func (d *AcyclicLP) PathTo(v int) *lib.Stack[*lib.DirectedEdge] {
	path := lib.NewStack[*lib.DirectedEdge]()
	if !d.HasPathTo(v) {
		return path
	}

	for e := d.EdgeTo[v]; e != nil; e = d.EdgeTo[e.From()] {
		path.Push(e)
	}
	return path
}
