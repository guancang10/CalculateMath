package calculatemath

func SquareRoot(number int) int {
	return number * number
}

func Factorial(number int) int {
	if number == 1 {
		return 1
	} else {
		return number * Factorial(number-1)
	}
}
