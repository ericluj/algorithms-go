package string

import "github.com/ericluj/algorithms-go/lib"

// 三向字符串快速排序
// 仅在中间子数组中的下一个字符继续递归排序（因为键的首字母都与切分字符相等）
type Quick3string struct{}

func NewQuick3string() *Quick3string {
	return &Quick3string{}
}

func (q *Quick3string) CharAt(s string, d int) int {
	if d < len(s) {
		return int(s[d])
	}
	return -1
}

func (q *Quick3string) Sort(a []string) {
	q.sort(a, 0, len(a)-1, 0)
}

func (q *Quick3string) sort(a []string, lo, hi, d int) {
	// nums[lo,lt-1]小于v
	// nums[gt+1,hi]大于v
	// nums[lt,i-1]等于v
	// nums[i,gt]未确定
	// 不断递归，缩小gt-i的值

	if hi <= lo {
		return
	}

	v := q.CharAt(a[lo], d)
	lt := lo
	i := lo + 1
	gt := hi

	for i <= gt {
		t := q.CharAt(a[i], d)
		if t < v {
			lib.ExchS(a, i, lt) // 交换位置后nums[lt]<v，nums[i]==v
			lt++                // lt++，此时nums[lt]==v，即上一步的nums[i]
			i++                 // 因为交换后肯定是nums[i]==v，所以取下一个数，减少一次比较
		} else if t > v {
			lib.ExchS(a, i, gt) // 交换位置后nums[gt]>v，nums[i]的值未确定，所以再次比较交换得到的值
			gt--                // gt--，此时nums[gt]未确定
		} else if t == v {
			i++ // i++，取一个新的未确定是来比较
		}
	}

	q.sort(a, lo, lt-1, d)
	// 还有未查找的字符
	if v >= 0 {
		q.sort(a, lt, gt, d+1)
	}
	q.sort(a, gt+1, hi, d)
}
