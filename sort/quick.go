package sort

import "github.com/ericluj/algorithms-go/lib"

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
	v := nums[lo]
	i := lo + 1
	j := hi
	for {
		//从左往右扫描，找到第一个大于等于a[lo]的元素
		for nums[i] < v {
			if i == hi {
				break
			}
			i++
		}
		//此时nums[i]应该是大于等于nums[lo]，或者为nums[hi](nums[hi]<nums[lo])

		//从右往左扫描，找到第一个小于等于a[lo]的元素
		for nums[j] > v {
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

// Quick3way 三向切分的快速排序
func Quick3way(nums []int) {
	quick3waySort(nums, 0, len(nums)-1)
}

func quick3waySort(nums []int, lo, hi int) {
	//nums[lo,lt-1]小于v
	//nums[gt+1,hi]大于v
	//nums[lt,i-1]等于v
	//nums[i,gt]未确定
	//不断递归，缩小gt-i的值

	if hi <= lo {
		return
	}

	v := nums[lo] //不能使用nums[lo]去做比较，因为元素交换，nums[lo]的值在改变
	lt := lo
	i := lo + 1
	gt := hi

	for i <= gt {
		if nums[i] < v {
			lib.Exch(nums, i, lt) //交换位置后nums[lt]<v，nums[i]==v
			lt++                  //lt++，此时nums[lt]==v，即上一步的nums[i]
			i++                   //取下一个数来进行比较
		} else if nums[i] > v {
			lib.Exch(nums, i, gt) //交换位置后nums[gt]>v，nums[i]的值未确定，所以不需要i++
			gt--                  //gt--，此时nums[gt]未确定
		} else if nums[i] == v {
			i++ //i++，取一个新的未确定是来比较
		}
	}

	quick3waySort(nums, lo, lt-1)
	quick3waySort(nums, gt+1, hi)
}
