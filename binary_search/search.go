package binary_search

// Search takes a sorted array of int and returns the index at which the target is found.
// Returns len(nums) if target is not found
func Search(target int, nums []int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return 1
		}
	}
	middle := len(nums) / 2
	//Check if the last index of the first half is less than target
	first := nums[0:middle]
	if first[len(first)-1] >= target {
		return Search(target, first)
	}
	//Check if the first index of the second half is greater than target
	last := nums[middle:]
	if last[0] <= target {
		return len(first) + Search(target, last)
	}
	return -1
}
