package search

import (
	"fmt"
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
	st.Delete(Key("L"))
	fmt.Println(st)
	st.Delete(Key("A"))
	fmt.Println(st)
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
	input := "SEARCHX"
	st := NewBST()
	for i, a := range input {
		st.Put(Key(a), i)
	}
	st.Print()
	assert.Equal(t, 1, st.Get("E"))
	assert.Equal(t, Key("E"), st.Floor("G"))
	assert.Equal(t, Key("H"), st.Select(3))
	assert.Equal(t, Key("E"), st.SelectFromOne(3))
	assert.Equal(t, 3, st.Rank(Key("H")))
	st.DeleteMin()
	st.Print()
	assert.Equal(t, Key("C"), st.Select(0))
	assert.Equal(t, Key("C"), st.Min().Key)
	st.Delete(Key("E"))
	st.Print()
	assert.Equal(t, []Key{Key("H"), Key("R"), Key("S")}, st.Keys(Key("H"), Key("S")))
}

// 红黑二叉查找树
func TestRedBlackBST(t *testing.T) {
	input := "SEARCHX"
	st := NewRedBlackBST()
	for i, a := range input {
		st.Put(Key(a), i)
	}
	st.Print()
	assert.Equal(t, 1, st.Get("E"))
	st.DeleteMin()
	st.Print()
	st.DeleteMax()
	st.Print()
}

// 基于拉链法的散列表
func TestChainHashST(t *testing.T) {
	input := "SEARCHX"
	st := NewSeparateChainingHashST(997)
	for i, a := range input {
		st.Put(Key(a), i)
	}
	fmt.Println(st)
	for i, a := range input {
		assert.Equal(t, i, st.Get(Key(a)))
	}
	st.Delete(Key("R"))
	fmt.Println(st)
}

// 基于线性探测法的散列表
func TestLineHashST(t *testing.T) {
	input := "SEARCHX"
	st := NewLinearProbingHashST(10)
	for i, a := range input {
		st.Put(Key(a), i)
	}
	fmt.Println(st)
	for i, a := range input {
		assert.Equal(t, i, st.Get(Key(a)))
	}
	st.Delete(Key("E"))
	fmt.Println(st)
}
