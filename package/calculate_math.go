package calculatemath

func SquareRoot(number int64, mutiply int64) int64 {
	if mutiply == 0 {
		return 1
	} else {
		return number * number * SquareRoot(number, mutiply-1)
	}
}

func SquareRoot2(number int64, mutiply int64) int64 {
	if mutiply == 0 {
		return 1
	} else {
		return number * SquareRoot2(number, mutiply-1)
	}
}

func SquareRoot3(number int64, mutiply int64) int64 {
	var result int64 = 1
	if mutiply == 0 {
		return 1
	}
	for i := 1; i <= int(mutiply); i++ {
		result *= number
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
