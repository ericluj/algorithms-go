package sp

import (
	"math"

	"github.com/ericluj/algorithms-go/lib"
)

// 最短路径的Dijkstra算法
type DijkstraSP struct {
	EdgeTo []*lib.DirectedEdge
	DistTo []float64 // 从顶点s到v的距离，如果不存在则路径为无穷大
	pq     *lib.IndexMinQueue
}

func NewDijkstraSP(g *EdgeWeightedDigraph, s int) *DijkstraSP {
	d := &DijkstraSP{
		EdgeTo: make([]*lib.DirectedEdge, g.V),
		DistTo: make([]float64, g.V),
		pq:     lib.NewIndexMinQueue(),
	}

	// 初始每个顶点给最大值
	for v := 0; v < g.V; v++ {
		d.DistTo[v] = math.MaxFloat64
	}
	// 起点给0
	d.DistTo[s] = 0.0
	// 起点放入队列
	d.pq.Insert(s, 0.0)

	for !d.pq.IsEmpty() {
		d.RelaxV(g, d.pq.DelMin())
	}

	return d
}

// 顶点的松弛
func (d *DijkstraSP) RelaxV(g *EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v).Data() {
		w := e.To()
		if d.DistTo[w] > d.DistTo[v]+e.Weight {
			d.DistTo[w] = d.DistTo[v] + e.Weight
			d.EdgeTo[w] = e

			if d.pq.Contains(w) {
				d.pq.Change(w, d.DistTo[w])
			} else {
				d.pq.Insert(w, d.DistTo[w])
			}
		}
	}
}

func (d *DijkstraSP) HasPathTo(v int) bool {
	return d.DistTo[v] < math.MaxFloat64
}

func (d *DijkstraSP) PathTo(v int) *lib.Stack[*lib.DirectedEdge] {
	path := lib.NewStack[*lib.DirectedEdge]()
	if !d.HasPathTo(v) {
		return path
	}

	for e := d.EdgeTo[v]; e != nil; e = d.EdgeTo[e.From()] {
		path.Push(e)
	}
	return path
}
