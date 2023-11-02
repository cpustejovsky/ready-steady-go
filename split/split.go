package split

// String takes a string and returns a slice of strings with the length of size
func String(str string, size int) []string {
	var split []string
	var temp = make([]rune, size)
	for i, char := range str {
		temp[i%size] = char
		if (i+1)%size == 0 {
			split = append(split, string(temp))
		}
	}
	return split
}
