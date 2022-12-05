package string

// 低位优先的字符串排序
type LSD struct{}

func NewLSD() *LSD {
	return &LSD{}
}

func (lsd *LSD) Sort(a []string, w int) {
	// 通过前w个字符将a[]排序
	N := len(a)
	R := 256
	aux := make([]string, N)

	for d := w - 1; d >= 0; d-- {
		// 根据第d个字符用键索引计数法排序

		// 计算出现频率（count[0]不使用）
		count := make([]int, R+1)
		for i := 0; i < N; i++ {
			count[a[i][d]+1]++
		}

		// 将频率转换为索引
		for i := 0; i < R; i++ {
			count[i+1] += count[i]
		}

		// 将元素分类
		for i := 0; i < N; i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}

		// 回写
		for i := 0; i < N; i++ {
			a[i] = aux[i]
		}
	}
}
