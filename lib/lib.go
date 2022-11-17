package lib

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// Exch 交换数组两个元素位置
func Exch(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
