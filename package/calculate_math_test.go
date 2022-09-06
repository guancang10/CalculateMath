package calculatemath

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/assert"
)

type Request struct {
	name     string
	number   int64
	mutiply  int64
	expected int64
}

func (request Request) CallMe() {
	fmt.Println(request.name)
}

func BenchmarkSquareroot2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRoot2(2, 1000)
	}
}

func BenchmarkSquareroot3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRoot3(2, 1000)
	}
}

func BenchmarkSquareroot4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRoot4(2, 1000)
	}
}

func BenchmarkSquareroot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareRoot(2, 10, FalseNumber)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("Running")

	m.Run()

	fmt.Println("Stopped")
}

func EndApp() {
	message := recover()
	fmt.Println("Error : ", message)
	fmt.Println("Done")
}

func TestTableDynamic(t *testing.T) {
	data := []Request{
		{
			name:     "Test1",
			number:   5,
			mutiply:  5,
			expected: 3125,
		},
		{
			name:     "Test2",
			number:   2,
			mutiply:  10,
			expected: 1024,
		},
		{
			name:     "Test3",
			number:   3,
			mutiply:  15,
			expected: 14348907,
		},
	}

	// test := []struct {
	// 	name     string
	// 	number   int64
	// 	mutiply  int64
	// 	expected int64
	// }{
	// 	{
	// 		name:     "Test1",
	// 		number:   5,
	// 		mutiply:  5,
	// 		expected: 3125,
	// 	},
	// 	{
	// 		name:     "Test2",
	// 		number:   2,
	// 		mutiply:  10,
	// 		expected: 1024,
	// 	},
	// 	{
	// 		name:     "Test3",
	// 		number:   3,
	// 		mutiply:  15,
	// 		expected: 14348907,
	// 	},
	// }
	defer EndApp()
	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			v.CallMe()
			result := SquareRoot(v.number, v.mutiply, FalseNumber)
			assert.Equal(t, v.expected, result, "result is not")
		})
	}
}

func TestSquareRoot(t *testing.T) {
	result := SquareRoot(2, 30, FalseNumber)
	assert.Equal(t, int64(1073741824), result, "Result not 1073741824")
	fmt.Println("test")

	t.Run("SubTest", func(t *testing.T) {
		result := SquareRoot(2, 10, FalseNumber)
		assert.Equal(t, int64(1024), result, "Result not 32")
	})
}

func TestSquareRoot2(t *testing.T) {
	result := SquareRoot2(2, 30)
	if result != 1073741824 {
		errorResult := result
		panic(errorResult)
	}
}

func TestSquareRoot3(t *testing.T) {
	result := SquareRoot3(2, 30)
	if result != 1073741824 {
		errorResult := result
		panic(errorResult)
	}
}

func TestFactorial(t *testing.T) {
	result := Factorial(5)
	if result != 120 {
		panic("Result is not 120")
	}
}
