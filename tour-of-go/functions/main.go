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

func main() {
	fmt.Println("#############################################################")
	fmt.Println("# Functions")
	fmt.Println("# ========================================================= #")

	div, rem := divideWithRemainder(14, 6)

	fmt.Println("# Adding 2 and 3 equals", add(2, 3))
	fmt.Println("# Subtracting 3 from 5 equals", subtract(5, 3))
	fmt.Println("# Dividing 14 by 6 equals", div, "with remainder", rem)
	fmt.Println("# Pi is", math.Pi)

	rand.Seed(86)
	fmt.Println("# Random integer", rand.Intn(1000))
	fmt.Println("# Random integer", rand.Intn(1000))
	fmt.Println("# Random integer", rand.Intn(1000))
	fmt.Println("#############################################################")

}
