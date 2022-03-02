package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(opA int, opB int) int {
	return opA + opB
}

func subtract(opA, opB int) int {
	return opA - opB
}
func divideWithRemainder(opA, opB int) (int, int) {
	return opA / opB, opA % opB
}

/**
Go Quirks:
- functions can have a "naked" return if a blank return statement with named return variables are returned
- named-return variables allow for a blank return statement?
- we cannot use the same params as named-return variables
- named return variables are auto-declared and initialized to zero
*/
func multiplyNumbers(opA, opB, multiplier int) (processedA, processedB int) {
	processedA = opA * multiplier
	processedB = opB * multiplier
	return
}

func main() {
	fmt.Println("#############################################################")
	fmt.Println("# Functions")
	fmt.Println("# ========================================================= #")

	div, rem := divideWithRemainder(14, 6)
	multipliedA, multipliedB := multiplyNumbers(13, 17, 3)

	fmt.Println("# Adding 2 and 3 equals", add(2, 3))
	fmt.Println("# Subtracting 3 from 5 equals", subtract(5, 3))
	fmt.Println("# Dividing 14 by 6 equals", div, "with remainder", rem)
	fmt.Println("# Multiplying 13, 17 by 3 equals", multipliedA, "and", multipliedB)
	fmt.Println("# Pi is", math.Pi)

	rand.Seed(86)
	fmt.Println("# Random integer", rand.Intn(1000))
	fmt.Println("# Random integer", rand.Intn(1000))
	fmt.Println("# Random integer", rand.Intn(1000))
	fmt.Println("#############################################################")

}
