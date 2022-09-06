package calculatemath

type UpdatedNumber func(int64, int64) int64

func FalseNumber(BaseNumber int64, ChangeNumber int64) int64 {
	return BaseNumber * ChangeNumber
}

func SquareRoot(number int64, mutiply int64, FalseNumber UpdatedNumber) int64 {
	if mutiply == 0 {
		return 1
	} else if mutiply == 1 {
		return number * SquareRoot(number, mutiply-1, FalseNumber)
	} else {
		return number * number * SquareRoot(number, mutiply-2, FalseNumber)
	}
}
func SquareRoot4(number int64, mutiply int64) int64 {
	if mutiply == 0 {
		return 1
	}
	temp := SquareRoot4(number, mutiply/2)
	if mutiply%2 == 0 {
		return temp * temp
	} else {
		return number * temp * temp
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
