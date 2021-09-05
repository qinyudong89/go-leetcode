package sort

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	left := []int{1, 2, 3, 5, 0, 0}
	right := []int{2, 5, 6}
	m := 4
	n := 3
	fmt.Println(Merge(left, m, right, n))
}
