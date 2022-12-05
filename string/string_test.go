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
