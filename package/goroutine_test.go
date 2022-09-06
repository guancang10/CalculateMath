package calculatemath

import (
	"fmt"
	"testing"
	"time"
)

func printHello() {
	fmt.Println("test")
}

func TestPrintHello(t *testing.T) {
	go printHello()
	fmt.Println("Hi")

	time.Sleep(1 * time.Second)
}
