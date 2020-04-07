package sort

// Selection 选择排序
func Selection(nums []int) {
	//先找到数组中最小的元素，其次将它和数组第一个元素交换位置
	//再次，在数组剩下元素找到最小的元素，将其与第二个元素交换位置
	//如此往复，直到数组排序结束
	l := len(nums)
	for i := 0; i < l; i++ {
		min := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[i] {
				min = j
			}
		}
		Exch(nums, i, min)
	}
}

// Exch 交换数组两个元素位置
func Exch(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}
