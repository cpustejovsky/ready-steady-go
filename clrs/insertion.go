package clrs

func InsertionSort(arr []int) []int {
	//loop through array
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		//if the previous index is greater than or equal to zero
		//and if the item in the previous index is greater than the value from the current index
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
