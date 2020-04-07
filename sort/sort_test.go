package sort

import (
	"fmt"
	"testing"
)

var data = []int{5, 39, 6666, 9, 54, 777, 324}

func TestSelection(t *testing.T) {
	fmt.Println(data)
	Selection(data)
	fmt.Println(data)
}
