package string

import (
	"github.com/ericluj/algorithms-go/graph/digraph"
	"github.com/ericluj/algorithms-go/lib"
)

// 正则表达式的模式匹配（grep)
type NFA struct {
	re string           // 匹配转换
	G  *digraph.Digraph // epsilon转换
	M  int              // 状态数量
}

func NewNFA(re string) *NFA {
	// 根据给定的正则表达式构造NFA
	ops := lib.NewStack[int]()
	M := len(re)
	G := digraph.NewDigraph(M + 1)

	for i := 0; i < M; i++ {
		lp := i

		// 或表达式：从'('指向第'|'后的第一个字符；从'|'指向')'
		if re[i] == '(' || re[i] == '|' {
			ops.Push(i)
		} else if re[i] == ')' {
			or := ops.Pop()
			if re[or] == '|' {
				lp = ops.Pop()
				G.AddEdge(lp, or+1)
				G.AddEdge(or, i)
			} else {
				lp = or
			}
		}

		// 闭包表达式
		if i < M-1 && re[i+1] == '*' { // 查看下一个字符
			G.AddEdge(lp, i+1)
			G.AddEdge(i+1, lp)
		}

		// 与下一个字符的连接
		if re[i] == '(' || re[i] == '*' || re[i] == ')' {
			G.AddEdge(i, i+1)
		}
	}
	return &NFA{
		re: re,
		G:  G,
		M:  M,
	}
}

func (n *NFA) Recognizes(txt string) bool {
	pc := lib.NewBag[int]()
	dfs := digraph.NewDirectedDFS(n.G, 0)
	for v := 0; v < n.G.V; v++ {
		if dfs.Marked[v] {
			pc.Add(v)
		}
	}

	for i := 0; i < len(txt); i++ {
		// 计算txt[i+1]可能到达的所有NFA状态
		match := lib.NewBag[int]()
		for _, v := range pc.Data() {
			if v < n.M {
				if n.re[v] == txt[i] || n.re[v] == '.' {
					match.Add(v + 1)
				}
			}
		}
		pc := lib.NewBag[int]()
		dfs := digraph.NewDirectedDFSSources(n.G, match.Data())
		for v := 0; v < n.G.V; v++ {
			if dfs.Marked[v] {
				pc.Add(v)
			}
		}
	}

	for _, v := range pc.Data() {
		if v == n.M {
			return true
		}
	}

	return false
}
