package search

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSST(t *testing.T) {
	input := "SEARCHEXAMPLE"
	sst := NewSST()
	for i, a := range input {
		sst.Put(Key(a), i)
		fmt.Println(sst)
	}

	assert.Equal(t, 12, sst.Get("E"))
}

func TestBinaryST(t *testing.T) {
	fmt.Println(Rank(10, 0, 6))
}

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
