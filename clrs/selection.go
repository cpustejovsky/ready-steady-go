package clrs

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minIdx := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minIdx] && j != i {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums

}
