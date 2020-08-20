package iteration

import (
	"bytes"
)

func Repeat(chars string, num int) string {
	var repeatedChars bytes.Buffer
	for i := 0; i < num; i++ {
		repeatedChars.WriteString(chars)
	}
	return repeatedChars.String()
}
