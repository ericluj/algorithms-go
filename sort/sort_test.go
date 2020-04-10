package sort

import (
	"fmt"
	"testing"
)

var data = []int{5, 7, 3, 2, 0, 1, 6, 4, 9, 8}

// IsSorted 判断是否为有序数组
func IsSorted(t *testing.T, nums []int) {
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			t.Error("不是有序数组")
			return
		}
	}
	fmt.Println(nums)
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
