package main

import (
	"fmt"
	"math"
	"math/rand"
)

func add(opA int, opB int) int {
	return opA + opB
}

func main() {
	fmt.Println("#############################################################")
	fmt.Println("# Functions")
	fmt.Println("# ========================================================= #")

	fmt.Println("Adding 2 and 3 equals", add(2, 3))
	fmt.Println("Pi is", math.Pi)

	rand.Seed(86)
	fmt.Println("Random integer", rand.Intn(1000))
	fmt.Println("Random integer", rand.Intn(1000))
	fmt.Println("Random integer", rand.Intn(1000))
	fmt.Println("#############################################################")

}
