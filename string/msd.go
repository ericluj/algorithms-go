package string

import "github.com/ericluj/algorithms-go/lib"

// 高位优先的字符串排序
type MSD struct {
	R   int      // 基数
	M   int      // 小数组的切换阀值
	aux []string // 数组分类的辅助数组
}

func NewMSD() *MSD {
	msd := &MSD{
		R: 256,
		M: 15,
	}
	return msd
}

// 获取s中的第d个字符
// 如果字符串读取完了，那么返回-1
// 保证计算频率时，所有字符被检查完的字符串子数组排在所有子数组前面
func (msd *MSD) CharAt(s string, d int) int {
	if d < len(s) {
		return int(s[d])
	}
	return -1
}

func (msd *MSD) Sort(a []string) {
	N := len(a)
	msd.aux = make([]string, N)
	msd.sort(a, 0, N-1, 0)
}

func (msd *MSD) sort(a []string, lo, hi, d int) {
	// 以第d个字符为键将a[lo]至a[hi]排序
	if hi <= lo+msd.M {
		insertion(a, lo, hi, d)
		return
	}

	// 计算频率（count[0]不使用，count[1]表示字符串的结尾）
	count := make([]int, msd.R+2)
	for i := lo; i <= hi; i++ {
		count[msd.CharAt(a[i], d)+2]++
	}

	// 将频率转换为索引
	for i := 0; i < msd.R+1; i++ {
		count[i+1] += count[i]
	}

	// 数据分类
	for i := lo; i <= hi; i++ {
		msd.aux[count[msd.CharAt(a[i], d)+1]] = a[i]
		count[msd.CharAt(a[i], d)+1]++
	}

	// 回写
	for i := lo; i <= hi; i++ {
		a[i] = msd.aux[i-lo]
	}

	// 递归的以每个字符为键进行排序
	// 排序的是每个子数组
	for i := 0; i < msd.R; i++ {
		msd.sort(a, lo+count[i], lo+count[i+1]-1, d+1)
	}
}

// 根据第d个字符进行插入排序
func insertion(a []string, lo, hi, d int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(a[j], a[j-1], d); j-- {
			exch(a, j, j-1)
		}
	}
}

func less(v string, w string, d int) bool {
	for i := d; i < lib.Min(len(v), len(w)); i++ {
		if v[i] < w[i] {
			return true
		}
		if v[i] > w[i] {
			return false
		}
	}
	return len(v) < len(w)
}

func exch(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}
