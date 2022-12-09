package string

// Knuth-Morris-Pratt字符串查找算法
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
	// 定义一个最基础的DFA
	dfa[pat[0]][0] = 1 // 模式字符串的第一个字符+初始状态0 => 状态1
	x := 0             // 初始化重启状态
	for j := 1; j < M; j++ {
		// 计算dfa[][j]
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x] // 复制匹配失败情况下的值
		}
		dfa[pat[j]][j] = j + 1 // 设置匹配成功情况下的值
		x = dfa[pat[j]][x]     // 更新重启状态
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
	// 直到变为最终状态M
	// 或者跑完所有输入字符
	for ; i < N && j < M; i++ {
		j = k.dfa[txt[i]][j]
	}

	if j == M { // 找到匹配（到达模式字符串的末尾）
		return i - M
	}

	return -1 // 未找到匹配
}
