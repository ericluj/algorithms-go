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
