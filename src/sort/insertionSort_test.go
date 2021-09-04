package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	sample := []int{3, 4, 5, 2, 1}
	InsertionSort(sample)
	fmt.Println(sample)

	sample = []int{3, 4, 5, 2, 1, 7, 8, -1, -3}
	InsertionSort(sample)
	fmt.Println(sample)
}
