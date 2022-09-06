package calculatemath

func SquareRoot(number int, mutiply int) int {
	result := 0
	for i := 0; i < mutiply; i++ {
		result += number * number
	}
	return result
}

func Factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * Factorial(number-1)
	}
}
