package sort

import "github.com/ericluj/algorithms-go/lib"

// 三向切分的快速排序
// 标准快速排序在重复值较多的时候会导致效率下降
// 三向切分将数组分为小于、等于、大于三部分，递归中不再处理等于部分
func Quick3way(nums []int) {
	quick3waySort(nums, 0, len(nums)-1)
}

func quick3waySort(nums []int, lo, hi int) {
	// nums[lo,lt-1]小于v
	// nums[gt+1,hi]大于v
	// nums[lt,i-1]等于v
	// nums[i,gt]未确定
	// 不断递归，缩小gt-i的值

	if hi <= lo {
		return
	}

	v := nums[lo] // 不能使用nums[lo]去做比较，因为元素交换，nums[lo]的值在改变
	lt := lo
	i := lo + 1
	gt := hi

	for i <= gt {
		if nums[i] < v {
			lib.Exch(nums, i, lt) // 交换位置后nums[lt]<v，nums[i]==v
			lt++                  // lt++，此时nums[lt]==v，即上一步的nums[i]
			i++                   // 因为交换后肯定是nums[i]==v，所以取下一个数，减少一次比较
		} else if nums[i] > v {
			lib.Exch(nums, i, gt) // 交换位置后nums[gt]>v，nums[i]的值未确定，所以再次比较交换得到的值
			gt--                  // gt--，此时nums[gt]未确定
		} else if nums[i] == v {
			i++ // i++，取一个新的未确定是来比较
		}
	}

	quick3waySort(nums, lo, lt-1)
	quick3waySort(nums, gt+1, hi)
}
