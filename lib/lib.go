package lib

func Min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// 交换int数组两个元素位置
func Exch(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 交换string数组两个元素位置
func ExchS(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}
