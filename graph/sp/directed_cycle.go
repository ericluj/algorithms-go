package sp

import "github.com/ericluj/algorithms-go/lib"

// 寻找有向环
type DirectedCycle struct {
	marked  []bool
	edgeTo  []*lib.DirectedEdge
	cycle   *lib.Stack[*lib.DirectedEdge] // 有向环中的所有顶点（如果存在）
	onStack []bool                        // 递归调用的栈上的所有顶点
}

func NewDirectedCycle(g *EdgeWeightedDigraph) *DirectedCycle {
	dc := &DirectedCycle{
		marked:  make([]bool, g.V),
		edgeTo:  make([]*lib.DirectedEdge, g.V),
		onStack: make([]bool, g.V),
	}
	for v := 0; v < g.V; v++ {
		if !dc.marked[v] {
			dc.dfs(g, v) // 以v为起点的所有可达路径
		}
	}
	return dc
}

// 找到一条边v->w且w已经存在于栈中，说明找到了一个环
// 因为栈表示一条w->v的有向路径（v是最后一个放进栈中的顶点，已经在栈中的任何一个顶点，一定都会指向v）
// 而v->w正好补全了这个环
// 简单一点想：dfs会沿着有向路径一直递归，所以肯定是从最起点延伸出一个指向路径本身才会有环；而不会是从另一个起点指向之前的递归路径
func (dc *DirectedCycle) dfs(g *EdgeWeightedDigraph, v int) {
	dc.onStack[v] = true
	dc.marked[v] = true

	for _, e := range g.Adj(v).Data() {
		w := e.To()
		if dc.HasCycle() { // 已经有环，结束
			return
		} else if !dc.marked[w] { // 没有递归过的顶点，继续递归
			// 通过数组构造指向关系树
			// 因为所有顶点只会放入一次，所以不需要清空edgeTo，只有第一个遍历到的指向关系会放入
			dc.edgeTo[w] = e
			dc.dfs(g, w)
		} else if dc.onStack[w] { // 遇到递归过的顶点，并且是本次递归中的，那么说明有一个环了
			dc.cycle = lib.NewStack[*lib.DirectedEdge]()
			f := e
			for f.From() != w {
				dc.cycle.Push(f)
				f = dc.edgeTo[f.From()]
			}
			dc.cycle.Push(f)
		}
	}
	dc.onStack[v] = false
}

func (dc *DirectedCycle) HasCycle() bool {
	return dc.cycle != nil
}

func (dc *DirectedCycle) Cycle() *lib.Stack[*lib.DirectedEdge] {
	return dc.cycle
}
