package clrs

func SelectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		minimumIndex := i
		for currentIndex := i + 1; currentIndex < len(nums); currentIndex++ {
			if nums[currentIndex] < nums[minimumIndex] {
				minimumIndex = currentIndex
			}
		}
		nums[i], nums[minimumIndex] = nums[minimumIndex], nums[i]
	}
	return nums
}
