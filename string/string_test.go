package string

import (
	"fmt"
	"testing"
)

// 低位优先的字符串排序
func TestLSD(t *testing.T) {
	data := []string{
		"4PGC938",
		"2IYE230",
		"3CI0720",
		"1ICK750",
		"1OHV845",
		"4JZY524",
		"1ICK750",
		"3CI0720",
		"1OHV845",
		"1OHV845",
		"2RLA629",
		"2RLA629",
		"3ATW723",
	}
	lsd := NewLSD()
	lsd.Sort(data, 7)
	fmt.Println(data)
}

// 高位优先的字符串排序
func TestMSD(t *testing.T) {
	dt := []string{
		"she",
		"sells",
		"seashells",
		"by",
		"the",
		"sea",
		"shore",
		"the",
		"shells",
		"she",
		"sells",
		"are",
		"surely",
		"seashells",
	}
	data := append(dt, dt...)
	msd := NewMSD()
	msd.Sort(data)
	fmt.Println(data)
}

// 三向字符串快速排序
func TestQuick3string(t *testing.T) {
	data := []string{
		"she",
		"sells",
		"seashells",
		"by",
		"the",
		"sea",
		"shore",
		"the",
		"shells",
		"she",
		"sells",
		"are",
		"surely",
		"seashells",
	}
	q := NewQuick3string()
	q.Sort(data)
	fmt.Println(data)
}

// 基于单词查找树的符号表
func TestTrieST(t *testing.T) {
	st := NewTrieST()
	// Put
	st.Put("shells", 1)
	st.Put("she", 2)
	st.Put("shr", 3)
	// Get
	fmt.Println(st.Get("sh"))
	fmt.Println(st.Get("she"))
	fmt.Println(st.Get("she1"))
	// 前缀匹配
	fmt.Println(st.KeysPrefix("she"))
	// 获取所有
	fmt.Println(st.Keys())
	// 通配符匹配
	fmt.Println(st.KeysThatMatch("shw"))
	fmt.Println(st.KeysThatMatch("shr"))
	fmt.Println(st.KeysThatMatch("sh."))
	// 最长键前缀
	fmt.Println(st.LongestPrefixOf("shellsw"))
	// Delete
	st.Delete("she")
	fmt.Println(st.Get("she"))
	fmt.Println(st.Get("shells"))
}

// 基于三向单词查找树的符号表
func TestTST(t *testing.T) {
	st := NewTST()
	// Put
	st.Put("shells", 1)
	st.Put("she", 2)
	st.Put("shr", 3)
	// Get
	fmt.Println(st.Get("sh"))
	fmt.Println(st.Get("she"))
	fmt.Println(st.Get("she1"))
}

// 暴力子字符串查找
func TestViolence(t *testing.T) {
	v := NewViolence()
	fmt.Println(v.Search("se", "teesell"))
	fmt.Println(v.Search("rt", "teesell"))
	fmt.Println(v.Search2("se", "teesell"))
	fmt.Println(v.Search2("rt", "teesell"))
}
