package graph

import "github.com/ericluj/algorithms-go/lib"

// 寻找有向环
type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	cycle   *lib.Stack // 有向环中的所有顶点（如果存在）
	onStack []bool     // 递归调用的栈上的所有顶点
}

func NewDirectedCycle(g *Digraph) *DirectedCycle {
	dc := &DirectedCycle{
		marked:  make([]bool, g.V),
		edgeTo:  make([]int, g.V),
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
func (dc *DirectedCycle) dfs(g *Digraph, v int) {
	dc.onStack[v] = true
	dc.marked[v] = true

	for _, w := range g.Adj(v).Data() {
		if dc.hasCycle() { // 已经有环，结束
			return
		} else if !dc.marked[w] { // 没有递归过的顶点，继续递归
			// 通过数组构造指向关系树
			// 因为所有顶点只会放入一次，所以不需要清空edgeTo，只有第一个遍历到的指向关系会放入
			dc.edgeTo[w] = v
			dc.dfs(g, w)
		} else if dc.onStack[w] { // 遇到递归过的顶点，并且是本次递归中的，那么说明有一个环了
			dc.cycle = lib.NewStack()
			for x := v; x != w; x = dc.edgeTo[x] {
				dc.cycle.Push(x)
			}
			dc.cycle.Push(w)
			dc.cycle.Push(v)
		}
	}
	dc.onStack[v] = false
}

func (dc *DirectedCycle) hasCycle() bool {
	return dc.cycle != nil
}

func (dc *DirectedCycle) Cycle() *lib.Stack {
	return dc.cycle
}
