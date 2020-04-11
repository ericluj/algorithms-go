package sort

import (
	"fmt"
	"testing"
)

var data = []int{5, 7, 3, 2, 0, 1, 6, 4, 9, 8}

// IsSorted 判断是否为有序数组
func IsSorted(t *testing.T, nums []int) {
	defer fmt.Println(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			t.Error("不是有序数组")
			return
		}
	}
}

func TestSelection(t *testing.T) {
	Selection(data)
	IsSorted(t, data)
}

func TestInsertion(t *testing.T) {
	Insertion(data)
	IsSorted(t, data)
}

func TestShell(t *testing.T) {
	Shell(data)
	IsSorted(t, data)
}

func TestMerge(t *testing.T) {
	Merge(data)
	IsSorted(t, data)
}

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
