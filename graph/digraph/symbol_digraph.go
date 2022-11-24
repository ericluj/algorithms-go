package digraph

import (
	"bufio"
	"os"
	"strings"
)

// 有向符号图
type SymbolDigraph struct {
	st   map[string]int // 符号名 -> 索引
	keys []string       // 索引 -> 符号名
	g    *Digraph       // 有向图
}

// 从文件中读入一幅有向符号图
func NewSymbolDigraphByFile(fileName, separator string) *SymbolDigraph {
	f, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f.Close()

	f2, err := os.Open(fileName)
	if err != nil {
		return nil
	}
	defer f2.Close()

	sg := &SymbolDigraph{
		st: make(map[string]int),
	}

	// 第一遍 构造索引
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), separator)
		for i := 0; i < len(arr); i++ {
			if _, ok := sg.st[arr[i]]; !ok {
				sg.st[arr[i]] = len(sg.st)
			}
		}
	}

	sg.keys = make([]string, len(sg.st))
	for k, v := range sg.st {
		sg.keys[v] = k
	}

	// 第二遍 构造图
	sg.g = NewDigraph(len(sg.st))
	scanner2 := bufio.NewScanner(f2)
	for scanner2.Scan() {
		arr := strings.Split(scanner2.Text(), separator)
		v := sg.st[arr[0]]
		for i := 1; i < len(arr); i++ {
			sg.g.AddEdge(v, sg.st[arr[i]])
		}
	}

	return sg
}

func (sg *SymbolDigraph) Contains(s string) bool {
	if _, ok := sg.st[s]; ok {
		return true
	}
	return false
}

func (sg *SymbolDigraph) Index(s string) int {
	return sg.st[s]
}

func (sg *SymbolDigraph) Name(v int) string {
	return sg.keys[v]
}

func (sg *SymbolDigraph) G() *Digraph {
	return sg.g
}
