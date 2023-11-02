package split

// String takes a string and returns a slice of strings with the length of size
func String(str string, size int) []string {
	var split []string
	var temp []rune
	for _, char := range str {
		if len(temp) == size {
			split = append(split, string(temp))
			temp = []rune{char}
		} else {
			temp = append(temp, char)
		}
	}
	return append(split, string(temp))
}
