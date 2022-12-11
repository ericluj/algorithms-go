package string

// Knuth-Morris-Pratt字符串查找算法
// DFA（确定有限状态自动机）
// 重启状态（当C匹配文本字符失败后，会继续使用重启位置处的模式字符和该文本进行匹配）
type KMP struct {
	pat string
	dfa [][]int
}

func NewKMP(pat string) *KMP {
	k := &KMP{
		pat: pat,
	}

	// 由模式字符串构造DFA
	R := 256
	M := len(pat)
	dfa := make([][]int, R)
	for i := 0; i < R; i++ {
		dfa[i] = make([]int, M)
	}

	// 定义一个最基础的DFA（状态0）
	dfa[pat[0]][0] = 1 // pat[0] + 状态0 => 状态1，列的其他位都是0

	x := 0 // 初始化重启状态

	// 定义状态0后的DFA（1到M-1）
	// 计算dfa[][j]
	for j := 1; j < M; j++ {
		// 复制重启状态对应的列
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x]
		}

		// 设置匹配成功情况下的值（成功时下一状态为j+1）
		dfa[pat[j]][j] = j + 1

		// 更新重启状态
		x = dfa[pat[j]][x]
	}

	k.dfa = dfa
	return k
}

func (k *KMP) Search(txt string) int {
	// 在txt上模拟DFA的运行
	M := len(k.pat)
	N := len(txt)
	i := 0
	j := 0

	// 字符的不断输入会让状态机的状态变化
	// 直到变为最终状态j=M
	// 或者跑完所有输入字符i=N
	for ; i < N && j < M; i++ {
		j = k.dfa[txt[i]][j]
	}

	if j == M { // 找到匹配（到达模式字符串的末尾）
		return i - M
	}

	return -1 // 未找到匹配
}
