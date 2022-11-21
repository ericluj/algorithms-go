package graph

import "fmt"

// 加权边
type Edge struct {
	v      int     // 顶点之一
	w      int     // 另一个顶点
	weight float64 // 边的权重
}

func NewEdge(v, w int, weight float64) *Edge {
	return &Edge{
		v:      v,
		w:      w,
		weight: weight,
	}
}

func (e *Edge) eigher() int {
	return e.v
}

func (e *Edge) other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	return e.v
}

func (e *Edge) compareTo(that *Edge) int {
	if e.weight < that.weight {
		return -1
	} else if e.weight > that.weight {
		return 1
	} else {
		return 0
	}
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.weight)
}
