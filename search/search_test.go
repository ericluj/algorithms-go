package search

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 顺序查找（基于无序链表）
func TestSST(t *testing.T) {
	input := "SEARCHEXAMPLE"
	st := NewSST()
	for i, a := range input {
		st.Put(Key(a), i)
		fmt.Println(st)
	}

	assert.Equal(t, 12, st.Get("E"))
}

// 二分查找（基于有序数组）
func TestBinaryST(t *testing.T) {
	input := "SEARCHEXAMPLE"
	st := NewBinanryST()
	for i, a := range input {
		st.Put(Key(a), i)
		fmt.Println(st)
	}

	assert.Equal(t, 12, st.Get("E"))
}

// 二叉查找树
func TestBST(t *testing.T) {
	input := "SEARCHEXAMPLE"
	st := NewBST()
	for i, a := range input {
		st.Put(Key(a), i)
	}
	st.Print()
	assert.Equal(t, 12, st.Get("E"))
	assert.Equal(t, Key("E"), st.Floor("G"))
	assert.Equal(t, Key("H"), st.Select(3))
	assert.Equal(t, Key("E"), st.SelectFromOne(3))
	assert.Equal(t, 3, st.Rank(Key("H")))
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
