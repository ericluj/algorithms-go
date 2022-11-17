package sort

import "github.com/ericluj/algorithms-go/lib"

// 插入排序
// 对部分有序数据很有效
func Insertion(nums []int) {
	//按索引位置依次移动数组中元素
	//不停与前一位比较，若比其小，则交换位置
	//与选择排序不同地方在于，遍历过程中，左侧所有元素的最终位置还不确定
	//适用于部分有序的数组

	length := len(nums)
	for i := 1; i < length; i++ { // 第1个元素是有序的
		for n := i; n > 0 && nums[n] < nums[n-1]; n-- {
			lib.Exch(nums, n, n-1)
		}
	}
}
