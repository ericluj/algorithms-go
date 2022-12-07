package string

// 暴力子字符串查找
type Violence struct {
}

func NewViolence() *Violence {
	return &Violence{}
}

func (v *Violence) Search(pat, txt string) int {
	m := len(pat)
	n := len(txt)
	for i := 0; i <= n-m; i++ { // 如果索引>n-m，最后剩的字符串小于m个肯定不会匹配成功
		j := 0
		for j < m {
			if txt[i+j] != pat[j] {
				break
			}
			j++
			if j == m { // 到了pat的尾部，找到匹配
				return i
			}
		}
	}

	return -1 // 未找到匹配
}

// 显式回退（暴力子字符匹配算法的另一种实现）
func (v *Violence) Search2(pat, txt string) int {
	m := len(pat)
	n := len(txt)
	i := 0
	j := 0

	for i < n && j < m {
		if txt[i] == pat[j] {
			j++
		} else {
			i -= j // 显式回退到本轮匹配开头，然后在后面i++到下一个字符，进行下一轮
			j = 0
		}

		i++
	}

	if j == m {
		return i - m
	}

	return -1 // 未找到匹配
}
