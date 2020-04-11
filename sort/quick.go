package sort

import "algorithms-go/lib"

// Quick 快速排序
func Quick(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, lo, hi int) {
	//快速排序是当两个子数组有序时，整个数组也就自然有序了
	//关键地方在于切分，切分需要使满足三个条件:
	//1.对于某个j，nums[j]已经排定
	//2.nums[lo]到nums[j-1]所有元素都不大于nums[j]
	//3.nums[j+1]到nums[hi]所有元素都不小于nums[j]

	if hi <= lo {
		return
	}
	j := partition(nums, lo, hi)
	quickSort(nums, lo, j-1) //这里是j-1，因为nums[j]已经排定
	quickSort(nums, j+1, hi)
}

func partition(nums []int, lo, hi int) int {

	i := lo + 1
	j := hi
	for {
		//从左往右扫描，找到第一个大于等于a[lo]的元素
		for nums[i] < nums[lo] {
			if i == hi {
				break
			}
			i++
		}
		//此时nums[i]应该是大于等于nums[lo]，或者为nums[hi](nums[hi]<nums[lo])

		//从右往左扫描，找到第一个小于等于a[lo]的元素
		for nums[j] > nums[lo] {
			if j == lo {
				break
			}
			j--
		}
		//此时nums[j]应该是小于等于nums[lo]，或者为nums[lo]

		if i >= j { //如果左右指针相遇，退出循环
			break
		}

		//交换i,j元素，保证i左侧元素不大于nums[lo]，j右侧元素不小于nums[lo]
		lib.Exch(nums, i, j)
	}
	lib.Exch(nums, lo, j)
	return j
}
