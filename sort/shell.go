package sort

// Shell 希尔排序
func Shell(nums []int) {
	l := len(nums)
	h := 1
	for h < (l / 3) {
		h = 3*h + 1
	}

	for h >= 1 {
		//将数组变为h有序
		for i := h; i < l; i++ {
			for j := i; j >= h && nums[j] < nums[j-h]; j -= h {
				Exch(nums, j, j-h)
			}
		}

		h = h / 3
	}
}
