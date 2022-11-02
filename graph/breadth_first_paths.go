package graph

import "github.com/ericluj/algorithms-go/lib"

// 广度优先搜索查找图中路径
type BreadthFirstPaths struct {
	marked []bool // 这个顶点上调用过dfs了吗？
	edgeTo []int  // 从起点到一个顶点的已知路径上的最后一个顶点
	s      int    // 起点
}

func NewBreadthFirstPaths(g *Graph, s int) *BreadthFirstPaths {
	b := &BreadthFirstPaths{
		marked: make([]bool, g.V),
		edgeTo: make([]int, g.V),
		s:      s,
	}
	b.bfs(g, s)
	return b
}

func (b *BreadthFirstPaths) bfs(g *Graph, s int) {
	queue := lib.NewQueue()
	b.marked[s] = true // 标记起点
	queue.Enqueue(s)   // 将它加入队列
	for !queue.IsEmpty() {
		v := queue.Dequeue() // 从队列中删去下一顶点
		for _, w := range g.Adj(v).Data() {
			if !b.marked[w] { // 对于每个未被标记的下一顶点
				b.edgeTo[w] = v    // 保存最短路径的最后一条边
				b.marked[w] = true // 标记它，因为最短路径已知
				queue.Enqueue(w)   // 并将它添加的队列中
			}
		}
	}
}

func (b *BreadthFirstPaths) hasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstPaths) pathTo(v int) *lib.Stack {
	path := lib.NewStack()
	if !b.hasPathTo(v) {
		return path
	}

	for x := v; x != b.s; x = b.edgeTo[x] {
		path.Push(x)
	}
	path.Push(b.s)
	return path
}
