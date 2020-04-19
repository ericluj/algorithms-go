package search

import (
	"fmt"
	"strconv"
	"testing"
)

func TestBST(t *testing.T) {
	b := NewBST()
	for i := 111; i <= 120; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	for i := 101; i <= 110; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	fmt.Println(b.Root.Size())
	fmt.Println(b.Get(Key("key" + strconv.Itoa(116))))
	fmt.Println(b.Min())
	fmt.Println(b.Max())
	fmt.Println(b.Floor("key126"))
	fmt.Println(b.Select(6))
	b.Delete("key106")
	fmt.Println(b.Get("key106"))
	fmt.Println("--------")
	b.Print()
	fmt.Println("--------")
	b.PrintStructure()
}

func TestRBBST(t *testing.T) {
	b := NewRedBlackBST()
	for i := 111; i <= 120; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	for i := 101; i <= 110; i++ {
		b.Put(Key("key"+strconv.Itoa(i)), i)
	}
	fmt.Println(b.Root.Size())
	fmt.Println(b.Get("key106"))
	fmt.Println("--------")
}
