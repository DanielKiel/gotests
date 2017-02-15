package main

import (
	"fmt"
)

var firstName = "Daniel" 
var lastName = "Koch"

const Pi = 3.14

/**
* only testing some go basics
*/
func main() {
	fmt.Println(sum(4, 6))
	fmt.Println(devide(10, 2))

	a, b := combineString(combineString("Hello", "go"))
	fmt.Println(a, b)

	fmt.Println(firstName, lastName)

	fmt.Println(simpleFor)
}

func simpleFor() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func forever() {
	//only write for and we have an endless for
	for {

	}
}

func sum(x int, y int) int {
	return x + y
}

func devide(x int, y int) int {
	return x / y
}

func combineString(x, y string) (string, string) {
	return x, y 
}