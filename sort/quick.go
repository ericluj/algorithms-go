package sort

import "github.com/ericluj/algorithms-go/lib"

// 快速排序
func Quick(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, lo, hi int) {
	// 快速排序是当两个子数组有序时，整个数组也就自然有序了
	// 关键地方在于切分，切分需要使满足三个条件:
	// 1.对于某个j，nums[j]已经排定
	// 2.nums[lo]到nums[j-1]所有元素都小于等于nums[j]
	// 3.nums[j+1]到nums[hi]所有元素都大于等于nums[j]

	if hi <= lo {
		return
	}
	j := partition(nums, lo, hi) // 切分
	quickSort(nums, lo, j-1)     // 将左半部分nums[lo..j-1]排序
	quickSort(nums, j+1, hi)     // 将右半部分nums[j+1..hi]排序
}

// 快速排序的切分
// 1.随意取nums[lo]为切分元素
// 2.从左向右扫描直到一个大于等于它的元素
// 3.从右向左扫描直到一个小于等于它的元素
// 4.交换找到的两个元素
// 5.继续2.3.4操作，可以保证左指针i左侧都小于等于切分元素，右指针j右侧都大于等于切分元素
// 6.当两个指针相遇，交换切分元素nums[lo]和左子数组最右侧元素nums[j]并返回j即可
func partition(nums []int, lo, hi int) int {
	p := nums[lo] // 切分元素
	i := lo       // 左指针
	j := hi + 1   // 右指针
	for {
		// 从左往右扫描，找到第一个大于等于p的元素
		for {
			i++
			if nums[i] >= p {
				break
			}

			// 找到了最后一个元素，说明全部小于p
			// j=hi 切分元素与最后一个元素交换位置
			if i == hi {
				break
			}
		}

		// 从右往左扫描，找到第一个小于等于p的元素
		for {
			j--
			if nums[j] <= p {
				break
			}

			// 找到了第一个元素，说明全部大于p
			// j=lo 切分元素位置没有变化
			// （这里的代码是冗余的，因为走到nums[lo]上面就break掉了）
			if j == lo {
				break
			}
		}

		// 能走到这里说明肯定不是全小于或者全大于
		// 所以i和j的位置在移动之前有nums[i]>=p nums[j]<=p
		// 可能有两种情况会触发break：
		//  1.i从左侧逼近j
		//	  nums[j]>=p 所以i移动到j的位置肯定会停下 i==j
		//	2.j从右侧逼近i
		//	  因为i先移动，所以肯定是i先移动到nums[i]>=p的位置
		//	  j从右侧逼近，nums[i]==p，j移动到i停下，
		//		         nums[i]>p,j移动到i-1停下，nums[i-1]<p
		// 因为切分元素在数组最左边，所以与其交换的元素必须小于等于它，
		// 所以return j
		if i >= j {
			break
		}

		// 交换i,j元素，保证i左侧元素小于等于p，j右侧元素大于等于p
		lib.Exch(nums, i, j)
	}
	lib.Exch(nums, lo, j)
	return j
}

// 三向切分的快速排序
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
			i++                   // 取下一个数来进行比较
		} else if nums[i] > v {
			lib.Exch(nums, i, gt) // 交换位置后nums[gt]>v，nums[i]的值未确定，所以不需要i++
			gt--                  // gt--，此时nums[gt]未确定
		} else if nums[i] == v {
			i++ // i++，取一个新的未确定是来比较
		}
	}

	quick3waySort(nums, lo, lt-1)
	quick3waySort(nums, gt+1, hi)
}
