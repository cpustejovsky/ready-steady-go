package bitwise

import (
	"strings"
)

func Or(x, y int) int {
	return x | y
}

func And(x, y int) int {
	return x & y
}

func IsEvenModulo(x int) bool {
	return x%2 == 0
}

func IsEvenBitwise(x int) bool {
	return x&1 == 0
}

func IsOddModulo(x int) bool {
	return x%2 != 0
}

func IsOddBitwise(x int) bool {
	return x&1 == 1
}

const (
	UPPER = 1 // upper case
	LOWER = 2 // lower case
	CAP   = 4 // capitalizes
	REV   = 8 // reverses
)

func Procstr(str string, conf byte) string {
	// reverse string
	rev := func(s string) string {
		runes := []rune(s)
		n := len(runes)
		for i := 0; i < n/2; i++ {
			runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
		}
		return string(runes)
	}

	// query configs
	if (conf & UPPER) != 0 {
		str = strings.ToUpper(str)
	}
	if (conf & LOWER) != 0 {
		str = strings.ToLower(str)
	}
	if (conf & CAP) != 0 {
		str = strings.Title(str)
	}
	if (conf & REV) != 0 {
		str = rev(str)
	}
	return str
}
