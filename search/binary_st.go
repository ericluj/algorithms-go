package search

import (
	"fmt"
	"strings"
)

// 二分查找（基于有序数组）
type BinanryST struct {
	keys []Key
	vals []Val
}

func NewBinanryST() *BinanryST {
	return &BinanryST{
		keys: make([]Key, 0),
		vals: make([]Val, 0),
	}
}

func (b *BinanryST) Get(k Key) Val {
	i := b.Rank(k, 0, len(b.keys)-1)
	if i < len(b.keys) && b.keys[i] == k { // 找到了
		return b.vals[i]
	}
	return nil
}

func (b *BinanryST) Put(k Key, v Val) {
	// 递归二分查找
	i := b.Rank(k, 0, len(b.keys)-1)

	// 迭代二分查找
	// i := b.Rank2(k)

	if i < len(b.keys) && b.keys[i] == k { // 找到了更新值
		b.vals[i] = v
		return
	}

	//找不到插入新元素
	b.keys = insert(b.keys, i, k)
	b.vals = insert(b.vals, i, v)
}

// 递归二分查找 返回的是比查询的k小的键数量
func (b *BinanryST) Rank(k Key, low, high int) int {
	// 向左没查到（所有数都比k大），low = 0		high = -1
	// 向右没查到（所有数都比k小），low = len	high = len - 1
	if low > high {
		return low
	}

	keys := b.keys
	mid := low + (high-low)/2
	cmp := keys[mid].CompareTo(k)

	if cmp == 0 {
		return mid
	} else if cmp < 0 { // vals[mid] < vals[k]
		low = mid + 1
	} else if cmp > 0 { // vals[mid] > vals[k]
		high = mid - 1
	}

	return b.Rank(k, low, high)
}

// 迭代二分查找 返回的是比查询的k小的键数量
func (b *BinanryST) Rank2(k Key) int {
	low := 0
	high := len(b.keys) - 1

	for low <= high {
		mid := low + (high-low)/2
		cmp := b.keys[mid].CompareTo(k)

		if cmp == 0 {
			return mid
		} else if cmp < 0 {
			low = mid + 1
		} else if cmp > 0 {
			high = mid - 1
		}
	}

	return low
}

func (b *BinanryST) String() string {
	keys := b.keys
	vals := b.vals

	arr := make([]string, 0)
	for i := range keys {
		arr = append(arr, fmt.Sprintf("%s:%v", keys[i], vals[i]))
	}
	return strings.Join(arr, ",")
}
