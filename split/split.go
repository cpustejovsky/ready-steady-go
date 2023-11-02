package split

func String(str string, size int) []string {
	var parts []string
	for i := 0; i < len(str); i++ {
		start := i * size
		end := start + size
		if start > len(str) || end > len(str) {
			break
		}
		parts = append(parts, str[start:end])
	}
	return parts
}

// StringUsingRunes takes a string and returns a slice of strings with the length of size
// This function is significantly slower than using the indexes of the str parameter.
func StringUsingRunes(str string, size int) []string {
	var split []string
	var temp string
	for _, char := range str {
		if len(temp) == size {
			split = append(split, temp)
			temp = string(char)
		} else {
			temp += string(char)
		}
	}
	return split
}
