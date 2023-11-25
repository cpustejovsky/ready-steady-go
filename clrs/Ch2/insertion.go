package Ch2

func InsertionSort(arr []int) []int {
	//loop through array
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		//this along with the for while loop are used to test the loop invariant at maintenance
		j := i - 1
		//while the previous index is greater than or equal to zero
		//and the item in the previous index is greater than the value from the current index
		for j >= 0 && arr[j] > key {
			//set the index 1 after j to be the value of the value at index j
			arr[j+1] = arr[j]
			//decrement j
			j--
		}
		//assign the value to the index 1 after j to the value from the current index
		//if the for loop never triggers, the assignment is arr[i] = arr[i]
		arr[j+1] = key
	}
	return arr
}

func ReverseInsertionSort(arr []int) []int {
	//loop through array
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		//this along with the for while loop are used to test the loop invariant at maintenance
		j := i - 1
		//while the previous index is greater than or equal to zero
		//and the item in the previous index is less than the value from the current index
		for j >= 0 && arr[j] < key {
			//set the index 1 after j to be the value of the value at index j
			arr[j+1] = arr[j]
			//decrement j
			j--
		}
		//assign the value to the index 1 after j to the value from the current index
		//if the for loop never triggers, the assignment is arr[i] = arr[i]
		arr[j+1] = key
	}
	return arr
}
