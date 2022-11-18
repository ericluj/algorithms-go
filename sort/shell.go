package sort

import "github.com/ericluj/algorithms-go/lib"

// 希尔排序
func Shell(nums []int) {
	//插入排序的改进版本
	//先对间隔h的逻辑数组进行插入排序，如此可以将元素移动到很远的地方，形成h有序数组
	//然后h不断递减至1，完成数组排序

	l := len(nums)
	h := 1

	//确定一个合适的间隔初始值(1, 4, 13, 40, 121, 364, 1093, ...)
	for h < (l / 3) {
		h = 3*h + 1
	}

	for h >= 1 {
		//将数组变为h有序
		for i := h; i < l; i++ { // h字数组的第一个元素是有序的
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				lib.Exch(nums, j, j-h)
			}
		}

		h = h / 3
	}
}
