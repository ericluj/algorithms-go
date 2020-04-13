package search

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBST(t *testing.T) {
	b := NewBST()
	for i := 20; i < 30; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	for i := 10; i < 20; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	fmt.Println(b.Get(Key("key" + strconv.Itoa(26))))
	fmt.Println(b.Min())
	fmt.Println(b.Max())
	fmt.Println(b.Floor(Key("key26")))
}
