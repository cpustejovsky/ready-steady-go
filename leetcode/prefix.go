package leetcode

func longestCommonPrefix(strings []string) string {
	var count int
	var temp int
	if len(strings) == 0 {
		return ""
	}
	if len(strings) == 1 {
		return strings[0]
	}
	word := strings[0]
	for i := 1; i < len(strings); i++ {
		for j := 0; j < len(strings[i]); j++ {
			if j < len(word) {
				if word[j] == strings[i][j] {
					temp++
				} else {
					break
				}
			}
		}
		if i > 1 && temp > count {
			break
		}
		count = temp
		temp = 0
	}
	return strings[0][0:count]
}
