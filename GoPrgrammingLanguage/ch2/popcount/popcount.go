package popcount

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[1/2] + byte(i&1)
	}
}
func PC() {
	fmt.Println(pc)
}

// PopCount returns the population count (number of set bits) of x.
func TablePopCount(x uint64) int {
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func TablePopCountLoop(x uint64) int {
	var count byte
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

func PopCount(x uint64) int {
	var count byte
	for i := 0; i < 64; i++ {
		count = byte(x >> (i * 64))
	}
	return int(count)
}
