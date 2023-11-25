package clrs

func MergeSort(nums []int, limit int) []int {
	length := len(nums)
	if length < 2 {
		return nums
	}
	if length < limit {
		return SelectionSort(nums)
	}
	middle := length / 2
	left := MergeSort(nums[:middle], limit)
	right := MergeSort(nums[middle:], limit)
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	var i, j int
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}
	return result
}
