package sp

import (
	"math"

	"github.com/ericluj/algorithms-go/lib"
)

// 基于队列的BellmanFord算法
// 队列中不出现重复的顶点
// 在某一轮中，改变了EdgeTo[]和DistTo[]的值的所有顶点都会在下一轮中处理
type BellmanFordSP struct {
	DistTo []float64                     // 从起点到某个顶点的路径长度
	EdgeTo []*lib.DirectedEdge           // 从起点到某个顶点的最后一条边
	OnQ    []bool                        // 该顶点是否存在于队列中
	Queue  *lib.Queue[int]               // 正在被放松的顶点
	Cost   int                           // RelaxV的调用次数
	Cycle  *lib.Stack[*lib.DirectedEdge] // EdgeTo中是否有负权重环
}

func NewBellmanFordSP(g *EdgeWeightedDigraph, s int) *BellmanFordSP {
	b := &BellmanFordSP{
		DistTo: make([]float64, g.V),
		EdgeTo: make([]*lib.DirectedEdge, g.V),
		OnQ:    make([]bool, g.V),
		Queue:  lib.NewQueue[int](),
	}
	for v := 0; v < g.V; v++ {
		b.DistTo[v] = math.MaxFloat64
	}
	b.DistTo[0] = 0.0
	b.Queue.Enqueue(s)
	b.OnQ[s] = true
	for !b.Queue.IsEmpty() && !b.HasNegativeCycle() {
		v := b.Queue.Dequeue()
		b.OnQ[v] = false
		b.RelaxV(g, v)
	}
	return b
}

// 顶点的松弛
func (b *BellmanFordSP) RelaxV(g *EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v).Data() {
		w := e.To()
		// 若边e引入小于之前最小的值，则e属于最短路径
		if b.DistTo[w] > b.DistTo[v]+e.Weight {
			b.DistTo[w] = b.DistTo[v] + e.Weight
			b.EdgeTo[w] = e
			if !b.OnQ[w] {
				b.Queue.Enqueue(w)
				b.OnQ[w] = true
			}
		}
		// 每调用V次RelaxV后检查负权重环
		b.Cost++
		if b.Cost%g.V == 0 {
			b.FindNegativeCycle()
		}
	}
}

// 将所有边放松V轮之后当且仅当队列非空时有向图才存在从起点可达的负权重环
func (b *BellmanFordSP) FindNegativeCycle() {
	V := len(b.EdgeTo)
	spt := NewEdgeWeightedDigraph(V)
	for v := 0; v < V; v++ {
		if b.EdgeTo[v] != nil {
			spt.AddEdge(b.EdgeTo[v])
		}
	}
	// 这里仅仅检测了是否有环，不检测负权重是否可以？
	finder := NewDirectedCycle(spt)
	b.Cycle = finder.Cycle()
}

func (b *BellmanFordSP) HasNegativeCycle() bool {
	return b.Cycle != nil
}

func (b *BellmanFordSP) HasPathTo(v int) bool {
	return b.DistTo[v] < math.MaxFloat64
}

func (b *BellmanFordSP) PathTo(v int) *lib.Stack[*lib.DirectedEdge] {
	path := lib.NewStack[*lib.DirectedEdge]()
	if !b.HasPathTo(v) {
		return path
	}

	for e := b.EdgeTo[v]; e != nil; e = b.EdgeTo[e.From()] {
		path.Push(e)
	}
	return path
}
