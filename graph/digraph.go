package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 有向图
type Digraph struct {
	V   int         // 顶点数目
	E   int         // 边的数目
	adj []*Bag[int] // 邻接表
}

// 创建一个含有v个顶点但不含有边的有向图
func NewDigraph(v int) *Digraph {
	g := &Digraph{
		V:   v,
		E:   0,
		adj: make([]*Bag[int], v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = NewBag[int]()
	}
	return g
}

// 从文件中读入一幅有向图
func NewDigraphByFile(fileName string) *Digraph {

	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// 不做错误处理，认为提供的文件一定符合格式
	scanner.Scan()
	v, _ := strconv.Atoi(scanner.Text()) // 顶点数目
	scanner.Scan()
	e, _ := strconv.Atoi(scanner.Text()) // 边的数目

	g := NewDigraph(v)
	// 添加有向边
	for i := 0; (i < e) && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(arr[0])
		w, _ := strconv.Atoi(arr[1])
		g.AddEdge(v, w)
	}

	return g
}

// 向有向图中添加一条边 v->w
func (g *Digraph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.E++
}

// 由v指出的边所连接的所有顶点
func (g *Digraph) Adj(v int) *Bag[int] {
	return g.adj[v]
}

// 该图的反向图
func (g *Digraph) Reverse() *Digraph {
	r := NewDigraph(g.V)
	for v := 0; v < g.V; v++ {
		for _, w := range g.Adj(v).Data() {
			r.AddEdge(w, v)
		}
	}
	return r
}

// 对象的字符串表示
func (g *Digraph) String() string {
	var res string
	for key, val := range g.adj {
		res += fmt.Sprintf("%d:%v\n", key, val)
	}
	return res
}
