package recursion

func SumUpRecursive(x int) int {
	//Handle Base Case
	if x < 1 {
		return x
	}
	return x + SumUpRecursive(x-1)
}

func SumUp(x int) int {
	var sum int
	for i := x; i > 0; i-- {
		sum = sum + i
	}
	return sum
}

func Fibonacci(x int) int {
	if x < 2 {
		return x
	}
	return Fibonacci(x-1) + Fibonacci(x-2)
}
