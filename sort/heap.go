package sort

// 堆排序
func Heap(nums []int) {
	N := len(nums)

	// 构造堆
	for k := N / 2; k >= 1; k-- {
		sink(nums, k, N)
	}

	// 将最大元素a[1]与a[N]交换，然后修复堆，将a[1]到a[N]的元素排序
	// 将exch()和less()实现中的索引减1即可得到与其他排序算法一致的实现
	for N > 1 {
		exch(nums, 1, N)
		N--
		sink(nums, 1, N)
	}
}

// 由上至下的堆有序化（下沉）
func sink(nums []int, k, n int) {
	// 如果结点k不是叶子结点2*k<=p.N
	for 2*k <= n {
		j := 2 * k // 左子结点

		// 如果左右子结点存在，找到子结点中较大的结点
		// 当j==p.N说明没有右子结点
		if (j < n) && less(nums, j, j+1) {
			j++
		}
		// 如果结点k大于它的两个子结点，无需下沉
		if less(nums, j, k) {
			break
		}
		// 将结点k与子结点中较大的结点交换位置
		exch(nums, k, j)
		k = j
	}
}

func less(nums []int, i, j int) bool {
	return nums[i-1] < nums[j-1]
}

func exch(nums []int, i, j int) {
	nums[i-1], nums[j-1] = nums[j-1], nums[i-1]
}
