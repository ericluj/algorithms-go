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

	// a.匹配失败
	// txt: A B A B A A A
	//				  i
	// pat: A B A B A C
	// 				  j
	// j = 5匹配失败，重新开始匹配
	// txt: A B A B A A A
	//				  i
	// pat: 	A B A B A C
	// 				  j
	// j = 2处为重启位置，匹配后的状态（重启状态）为3
	// 所以dfa[c][5] = dfa[c][3]

	// b.匹配成功
	// txt: A B A B A A A
	//				  i
	// pat: A B A B A C
	// 				  j
	// j = 5匹配成功，dfa[c][5] = j + 1 = 6

	// 总结ab: x为重启状态
	// 匹配失败 dfa[c][j] = dfa[c][x]
	// 匹配成功 dfa[c][j] = j + 1

	// 重启状态的初始值
	// 0的初始状态肯定为0
	// 1的初始状态为0
	// txt: A C A
	//		  i
	// pat: A B A B A C
	// 		  j
	// j = 5匹配失败，重新开始匹配
	// txt: A C A
	//		  i
	// pat:   A B A B A C
	// 		  j
	// j = -1处为重启位置，匹配后的状态（重启状态）为0

	x := 0 // 初始化重启状态

	// 定义状态0后的DFA（1到M-1）
	// 计算dfa[][j]
	for j := 1; j < M; j++ {
		// 不匹配时转换和重启状态x处相同
		// 复制重启状态对应的列
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x]
		}

		// 匹配时下一个状态为 j + 1
		dfa[pat[j]][j] = j + 1

		// 更新重启状态
		// 状态 j 和状态 j + 1 的重启状态存在递推关系
		// 假设状态 j 的重启状态为 x
		// 将 j 处的字符作为重启状态的输入
		// 得到的下一个值应该为 j + 1 的重启状态
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
