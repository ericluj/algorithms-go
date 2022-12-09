package string

// 暴力子字符串查找
type Violence struct {
}

func NewViolence() *Violence {
	return &Violence{}
}

func (v *Violence) Search(pat, txt string) int {
	M := len(pat)
	N := len(txt)
	for i := 0; i <= N-M; i++ { // 如果索引>n-m，最后剩的字符串小于m个肯定不会匹配成功
		j := 0
		for j < M {
			if txt[i+j] != pat[j] {
				break
			}
			j++
			if j == M { // 到了pat的尾部，找到匹配
				return i
			}
		}
	}

	return -1 // 未找到匹配
}

// 显式回退（暴力子字符匹配算法的另一种实现）
func (v *Violence) Search2(pat, txt string) int {
	M := len(pat)
	N := len(txt)
	i := 0
	j := 0

	for ; i < N && j < M; i++ {
		if txt[i] == pat[j] {
			j++
		} else {
			i -= j // 显式回退到本轮匹配开头，然后在后面i++到下一个字符，进行下一轮
			j = 0
		}
	}

	if j == M { // 找到匹配（到达模式字符串的末尾）
		return i - M
	}

	return -1 // 未找到匹配
}
