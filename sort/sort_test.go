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

// 快速排序
func TestQuick(t *testing.T) {
	Quick(data)
	IsSorted(t, data)
}

// 三向切分的快速排序
func TestQuick3way(t *testing.T) {
	Quick3way(data)
	IsSorted(t, data)
}

// 基于堆的优先队列
func TestMaxPQ(t *testing.T) {
	pq := NewMaxPQ(len(data))
	for _, item := range data {
		pq.Insert(item)
	}

	res := make([]int, len(data))
	i := len(data) - 1
	for !pq.IsEmpty() {
		res[i] = pq.DelMax()
		i--
	}
	IsSorted(t, res)
}

// 关联索引的优先队列
func TestIndexMinPQ(t *testing.T) {
	pq := NewIndexMinPQ(10)
	pq.Insert(1, 3.7)
	pq.Insert(1, 5.6)
	pq.Insert(2, 7.5)
	pq.Insert(3, 4.4)
	pq.Insert(4, 12.3)
	pq.Insert(5, 6.6)
	pq.Insert(6, 9.8)
	pq.Insert(7, 10.1)
	pq.Insert(8, 15.8)
	pq.Insert(9, 1.1)

	res := make([]int, 0)
	for !pq.IsEmpty() {
		i := pq.DelMin()
		res = append(res, i)
	}
	fmt.Println(res)
}

// 堆排序
func TestHeap(t *testing.T) {
	Heap(data)
	IsSorted(t, data)
}
