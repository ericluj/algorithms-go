package sort

import (
	"github.com/ericluj/algorithms-go/lib"
)

// 原地归并的抽象方法
func merge(nums []int, lo, mid, hi int) {
	// 将数组分为两个部分[lo,mid]和[mid+1,hi]，这里两个部分都是各自有序
	// 依次从两各部分中拿出元素进行比较，较小的元素放入
	// 若某一部分所有元素都取完，直接把另一部分剩下元素依次放入即可

	leftIndex := lo
	rightIndex := mid + 1

	// 先复制一份数据来读取，改动原来的数组
	c := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		c[i] = nums[i]
	}

	for k := lo; k <= hi; k++ {
		if leftIndex > mid { // 左半边全部取完，取右半边
			nums[k] = c[rightIndex]
			rightIndex++
		} else if rightIndex > hi { // 右半边全部取完，取左半边
			nums[k] = c[leftIndex]
			leftIndex++
		} else if c[leftIndex] < c[rightIndex] { // 左半边当前元素小于右半边当前元素取左
			nums[k] = c[leftIndex]
			leftIndex++
		} else { // 右半边元素小于等于左半边元素取右
			nums[k] = c[rightIndex]
			rightIndex++
		}
	}
}

// 自顶向下的归并排序
func Merge(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, lo, hi int) {
	if lo >= hi { // 此时数组只有一个元素，是有序的
		return
	}
	mid := lo + (hi-lo)/2      // 这种取mid的方式，可以避免溢出
	mergeSort(nums, lo, mid)   // 将左边边排序
	mergeSort(nums, mid+1, hi) // 将右半边排序
	merge(nums, lo, mid, hi)   // 归并结果
}

// 自底向上的归并排序
func MergeBu(nums []int) {
	// 其实是递归的反思路，不过代码量更少

	l := len(nums)
	for sz := 1; sz < l; sz += sz { // sz子数组大小
		for lo := 0; lo < l-sz; lo += sz + sz {
			// mid属于左半边
			// 左半边起点 lo	左半边终点 lo+sz-1
			// 右半边起点 lo+sz 右半边终点 lo+sz+sz-1
			merge(nums, lo, lo+sz-1, lib.Min(lo+sz+sz-1, l-1))
		}
	}
}
