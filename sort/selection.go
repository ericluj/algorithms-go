package sort

import "github.com/ericluj/algorithms-go/lib"

// 选择排序
// 对于长度为N的数组，选择排序需要大约(N^2)/2次比较和N次交换
func Selection(nums []int) {
	//先找到数组中最小的元素，其次将它和数组第一个元素交换位置
	//再次，在数组剩下元素找到最小的元素，将其与第二个元素交换位置
	//如此往复，直到数组排序结束

	length := len(nums)
	for i := 0; i < length; i++ {
		min := i // 本轮查找中最小数的index
		for n := i; n < length; n++ {
			if nums[n] < nums[min] {
				min = n
			}
		}
		// 将最小的数放在本轮查找首位
		lib.Exch(nums, i, min)
	}
}
