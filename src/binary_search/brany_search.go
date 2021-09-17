package binary_search

func BranySearch(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for min <= max {
		mid := min + (max-min)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}
