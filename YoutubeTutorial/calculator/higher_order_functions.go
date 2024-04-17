package calculator

func SumFromAToB(a, b int) int {
	return ProcessFromAToB(a, b, 0, func(x, y int) int {
		return x + y
	})
}

func MultiplyFromAToB(a, b int) int {
	return ProcessFromAToB(a, b, 1, func(x, y int) int {
		return x * y
	})
}

func AddX(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func ProcessFromAToB(a, b, initValue int, fn func(int, int) int) int {
increment := AddX(1)

	if a > b {
		return initValue
	}
	return fn(a, ProcessFromAToB(increment(a), b, initValue, fn))
}
