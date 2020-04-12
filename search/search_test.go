package search

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBST(t *testing.T) {
	b := NewBST()
	for i := 0; i < 100; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	fmt.Println(b.Get(Key("key" + strconv.Itoa(66))))
}
