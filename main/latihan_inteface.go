package main

import (
	"fmt"
)

type Movement interface {
	GetName() (string, int)
	GetAge() int
}

func walking(move Movement) {
	fmt.Println(move.GetName())
}

type Person struct {
	name string
	age  int
	foot int
}

func (person Person) GetName() (string, int) {
	return person.name, person.foot
}

func (person Person) GetAge() int {
	return person.age
}

func main() {
	fmt.Println("test")
	person := Person{name: "fikri", age: 12, foot: 2}
	walking(person)
}
