package leetcode

func TwoSum(nums []int, sum int) []int {
	var result []int
	for i, augend := range nums {
		for j, addend := range nums {
			if augend + addend == sum && i != j {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

type sumMap map[int]int

func TwoSumsTwoPassHashTable(nums []int, sum int) []int{
	sm := sumMap{}
	for i, num := range nums {
		sm[num] = i
	}
	for i, num := range nums {
		complement := sum - num
		val, ok := sm[complement]
		if ok && val != i {
			return []int{i, val}
		}
	}
	return []int{}
}

func TwoSumsOnePassHashTable(nums []int, sum int) []int{
	sm := sumMap{}
	for i, num := range nums {
		complement := sum - num
		val, ok := sm[complement];if ok {
			return []int{i, val}
		}
		sm[num] = i
	}
	return []int{}
}