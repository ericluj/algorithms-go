package string

// Boyer-Moore字符串匹配算法（启发式地处理不匹配的字符）
type BoyerMoore struct {
	right []int
	pat   string
}

func NewBoyerMoore(pat string) *BoyerMoore {
	b := &BoyerMoore{
		pat: pat,
	}

	R := 256
	M := len(pat)
	b.right = make([]int, R)
	for c := 0; c < R; c++ {
		b.right[c] = -1 // 不包含在模式字符串中的字符的值为-1
	}
	for j := 0; j < M; j++ { // 包含在模式字符串中的字符的值为它在其中出现的最右位置
		b.right[pat[j]] = j
	}

	return b
}

func (b *BoyerMoore) Search(txt string) int {
	// 在txt中查找模式字符串
	M := len(b.pat)
	N := len(txt)
	skip := 0

	for i := 0; i <= N-M; i += skip {
		skip = 0
		for j := M - 1; j >= 0; j-- {
			if b.pat[j] != txt[i+j] { // 字符匹配失败，i右移继续匹配
				// txt[i+j]不在pat中，i右移到j+1处再匹配下一轮
				// txt[i+j]不在pat中，使用right数组来将模式字符串和文本对齐
				skip = j - b.right[txt[i+j]]
				// 如果skip无法增大i，那么模式字符串至少右移一个位置
				if skip < 1 {
					skip = 1
				}
				break
			}
		}
		if skip == 0 { // 走到这里说明pat从右向左走完了
			return i // 找到匹配
		}
	}
	return -1 // 未找到匹配
}
