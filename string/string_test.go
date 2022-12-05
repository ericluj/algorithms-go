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
