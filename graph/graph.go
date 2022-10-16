package graph

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	V   int    // 顶点数目
	E   int    // 边的数目
	adj []*Bag // 邻接表
}

func NewGraph(v int) *Graph {
	g := &Graph{
		V:   v,
		E:   0,
		adj: make([]*Bag, v),
	}
	for i := 0; i < v; i++ {
		g.adj[i] = NewBag()
	}
	return g
}

func NewGraphByFile(fileName string) *Graph {

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

	g := NewGraph(v)
	// 添加边
	for i := 0; (i < e) && scanner.Scan(); i++ {
		arr := strings.Split(scanner.Text(), " ")
		v, _ := strconv.Atoi(arr[0])
		w, _ := strconv.Atoi(arr[1])
		g.AddEdge(v, w)
	}

	return g
}

func (g *Graph) GetV() int {
	return g.V
}

func (g *Graph) GetE() int {
	return g.E
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.E++
}

func (g *Graph) Adj(v int) *Bag {
	return g.adj[v]
}

func (g *Graph) String() string {
	var res string
	for key, val := range g.adj {
		res += fmt.Sprintf("%d:%v\n", key, val)
	}
	return res
}
