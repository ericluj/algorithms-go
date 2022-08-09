package sort

import (
	"github.com/ericluj/algorithms-go/lib"
)

// merge 原地归并
func merge(nums []int, lo, mid, hi int) {
	//将数组分为两个部分[lo,mid]和[mid+1,hi]，这里两个部分都是各自有序
	//依次从两各部分中拿出元素进行比较，较小的元素放入
	//若某一部分所有元素都取完，直接把另一部分剩下元素依次放入即可

	i := lo
	j := mid + 1

	//先复制一份数据
	b := make([]int, len(nums))
	copy(b, nums)

	for k := lo; k <= hi; k++ {
		if i > mid {
			nums[k] = b[j]
			j++
		} else if j > hi {
			nums[k] = b[i]
			i++
		} else if b[i] < b[j] {
			nums[k] = b[i]
			i++
		} else {
			nums[k] = b[j]
			j++
		}
	}
}

// Merge 自顶向下的归并排序
func Merge(nums []int) {
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := lo + (hi-lo)/2 //这种取mid的方式，可以避免溢出
	mergeSort(nums, lo, mid)
	mergeSort(nums, mid+1, hi)
	merge(nums, lo, mid, hi)
}

// MergeBu 自底向上的归并排序
func MergeBu(nums []int) {
	//其实是递归的反思路，不过代码量更少

	l := len(nums)
	for sz := 1; sz < l; sz += sz { //sz子数组大小
		for lo := 0; lo < l-sz; lo += sz + sz {
			merge(nums, lo, lo+sz-1, lib.Min(lo+sz-1+sz, l-1))
		}
	}
}
