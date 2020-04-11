package sort

import "algorithms-go/lib"

// Insertion 插入排序
func Insertion(nums []int) {
	//按索引位置依次移动数组中元素
	//不停与前一位比较，若比其小，则交换位置
	//与选择排序不同地方在于，遍历过程中，左侧所有元素的最终位置还不确定
	//适用于部分有序的数组

	l := len(nums)
	for i := 1; i < l; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			lib.Exch(nums, j, j-1)
		}
	}
}
