package sort

import "github.com/ericluj/algorithms-go/lib"

// Selection 选择排序
func Selection(nums []int) {
	//先找到数组中最小的元素，其次将它和数组第一个元素交换位置
	//再次，在数组剩下元素找到最小的元素，将其与第二个元素交换位置
	//如此往复，直到数组排序结束

	l := len(nums)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		lib.Exch(nums, i, min)
	}
}
