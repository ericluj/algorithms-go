package sort

import (
	"testing"
)

var data = []int{5, 39, 6666, 9, 54, 777, 324}

// IsSorted 判断是否为有序数组
func IsSorted(t *testing.T, nums []int) {
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
