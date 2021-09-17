package binary_search

import (
	"fmt"
	"testing"
)

func TestBranySearch(t *testing.T) {
	var b = []int{1, 2, 3, 4, 5}
	fmt.Println(BranySearch(b, 3))
	b = []int{-4, 0, 3, 2, 5}
	fmt.Println(BranySearch(b, 0))
	fmt.Println(BranySearch(b, 5))
	fmt.Println(BranySearch(b, 33))
}
