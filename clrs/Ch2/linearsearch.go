package Ch2

func LinearSearch(nums []int, val int) *int {
	/*
		Loop Invariant:
			At the start of each iteration of the for-range loop,
			there is no index k < j such that nums[k] == val

		At Initialization:
			Prior to first iteration, there is no pointer that could be assigned so index remains nil

		At Maintenance:
			It remains true before iteration n because no assigment has been reached because the loop would've been
			terminated. The same applies for before iteration n+1.

		At Termination:
			If the loop terminates before ranging through all nums,
			it is because we have found an index that contains the search value

			If the loop terminates after ranging through all nums,
			it is because we could not find a satisfactory index, leaving nil as the correct answer.
	*/
	for i, num := range nums {
		if num == val {
			return &i
		}
	}
	return nil
}
