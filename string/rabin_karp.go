package string

// Rabin-Karp指纹字符串查找算法
type RabinKarp struct {
	pat     string // 模式字符串（仅拉斯维加斯算法需要）
	patHash int64  // 模式字符串的散列值
	M       int    // 模式字符串的长度
	Q       int64  // 一个很大的素数
	R       int64  // 字母表的大小
	RM      int64  // R^(M-1) % Q
}

func NewRabinKarp(pat string) *RabinKarp {
	r := &RabinKarp{
		pat: pat,
		M:   len(pat),
		Q:   longRandomPrime(),
		R:   256,
		RM:  1,
	}
	for i := 1; i <= r.M-1; i++ { // 计算R^(M-1) % Q
		r.RM = (r.R * r.RM) % r.Q // 用于减去第一个数字时的计算
	}
	r.patHash = r.hash(pat)
	return r
}

func longRandomPrime() int64 {
	return 997
}

func (r *RabinKarp) hash(key string) int64 {
	// 计算key[0..M-1]的散列值
	var h int64
	for j := 0; j < r.M; j++ {
		h = (r.R*h + int64(key[j])) % r.Q
	}
	return h
}

func (r *RabinKarp) check(txt string, i int) bool {
	for j := 0; j < r.M; j++ {
		if r.pat[j] != txt[i+j] {
			return false
		}
	}
	return true
}

func (r *RabinKarp) Search(txt string) int {
	// 在文本中查找相等的散列值
	N := len(txt)
	txtHash := r.hash(txt)
	if r.patHash == txtHash && r.check(txt, 0) { // 一开始就匹配成功
		return 0
	}

	for i := r.M; i < N; i++ {
		// 减去第一个数字，加上最后一个数字，再次检查匹配
		// 额外加上一个Q来保证所有数均为正，这样取余操作才能得到预期的效果
		txtHash = (txtHash + r.Q - r.RM*int64(txt[i-r.M])%r.Q) % r.Q
		txtHash = (txtHash*r.R + int64(txt[i])) % r.Q
		if r.patHash == txtHash && r.check(txt, i-r.M+1) {
			return i - r.M + 1 // 找到匹配
		}
	}
	return -1 // 未找到匹配
}
