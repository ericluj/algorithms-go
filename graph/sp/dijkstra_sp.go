package sp

import (
	"math"

	"github.com/ericluj/algorithms-go/lib"
)

// 最短路径的Dijkstra算法
type DijkstraSP struct {
	EdgeTo []*lib.DirectedEdge
	DistTo []float64
}

func NewDijkstraSP(g *EdgeWeightedDigraph, s int) *DijkstraSP {
	d := &DijkstraSP{
		EdgeTo: make([]*lib.DirectedEdge, g.V),
		DistTo: make([]float64, g.V),
	}

	// 初始每个顶点给最大值
	for v := 0; v < g.V; v++ {
		d.DistTo[v] = math.MaxFloat64
	}
	// 起点给0
	d.DistTo[s] = 0.0

	return d
}

// 边的松弛
func (d *DijkstraSP) RelaxEdge(e *lib.DirectedEdge) {
	v := e.From()
	w := e.To()
	if d.DistTo[w] > d.DistTo[v]+e.Weight {
		d.DistTo[w] = d.DistTo[v] + e.Weight
		d.EdgeTo[w] = e
	}
}

// 顶点的松弛
func (d *DijkstraSP) RelaxV(g *EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v).Data() {
		w := e.To()
		if d.DistTo[w] > d.DistTo[v]+e.Weight {
			d.DistTo[w] = d.DistTo[v] + e.Weight
			d.EdgeTo[w] = e
		}
	}
}
