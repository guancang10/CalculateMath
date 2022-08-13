package calculatemath

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M){
	fmt.Println("Running")

	m.Run()

	fmt.Println("Stopped")
}

func TestSquareRoot(t *testing.T){
	result := SquareRoot(2,15)
	assert.Equal(t,int64(1073741824),result,"Result not 1073741824")
	fmt.Println("test")
}

func TestSquareRoot2(t *testing.T){
	result := SquareRoot2(2,30)
	if result != 1073741824{
		errorResult := result
		panic(errorResult)
	}
}

func TestSquareRoot3(t *testing.T){
	result := SquareRoot3(2,30)
	if result != 1073741824{
		errorResult := result
		panic(errorResult)
	}
}

func TestFactorial(t *testing.T){
	result := Factorial(5)
	if result != 120{
		panic("Result is not 120")
	}
}