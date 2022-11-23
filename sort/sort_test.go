package sort

import (
	"fmt"
	"testing"
)

var data = []int{5, 7, 3, 2, 0, 1, 6, 4, 9, 8}

// 判断是否为有序数组
func IsSorted(t *testing.T, nums []int) {
	defer fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			t.Error("不是有序数组")
			return
		}
	}
}

// 选择排序
func TestSelection(t *testing.T) {
	Selection(data)
	IsSorted(t, data)
}

// 插入排序
func TestInsertion(t *testing.T) {
	Insertion(data)
	IsSorted(t, data)
}

// 插入排序优化版
func TestInsertionPro(t *testing.T) {
	InsertionPro(data)
	IsSorted(t, data)
}

func BenchmarkInsertion(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Insertion(data)
	}
}

func BenchmarkInsertionPro(b *testing.B) {
	for n := 0; n < b.N; n++ {
		InsertionPro(data)
	}
}

// 希尔排序
func TestShell(t *testing.T) {
	Shell(data)
	IsSorted(t, data)
}

// 自顶向下的归并排序
func TestMerge(t *testing.T) {
	Merge(data)
	IsSorted(t, data)
}

// 自底向上的归并排序
func TestMergeBu(t *testing.T) {
	MergeBu(data)
	IsSorted(t, data)
}

func TestQuick(t *testing.T) {
	Quick(data)
	IsSorted(t, data)
}

func TestQuick3way(t *testing.T) {
	Quick3way(data)
	IsSorted(t, data)
}
