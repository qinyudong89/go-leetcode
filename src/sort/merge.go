package sort

func Merge(left []int, m int, right []int, n int) []int {
	for n > 0 {
		if m == 0 || right[n-1] > left[m-1] {
			left[m+n-1] = right[n-1]
			n--
		} else {
			left[m+n-1] = left[m-1]
			m--
		}
	}
	return left
}
