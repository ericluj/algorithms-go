package search

type Key string

type Val interface{}

type ST interface {
	Put(k Key, v Val)
	Get(k Key) Val
}

func (k Key) CompareTo(newK Key) int {
	if k < newK {
		return -1
	} else if k > newK {
		return 1
	} else {
		return 0
	}
}

func insert[T Key | Val](silce []T, i int, val T) []T {
	tmp := make([]T, 0)
	tmp = append(tmp, silce[:i]...)
	tmp = append(tmp, val)
	tmp = append(tmp, silce[i:]...)
	return tmp
}
