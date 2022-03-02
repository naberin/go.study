package main

import (
	"fmt"
	"strconv"
)

/**
Go Quirks:
- Go does not have ternary operators, if-else IS the idiomatic way
*/
func fizzBuzzButOnlyEvalNumber(N int) (val string) {

	if N%3 == 0 {
		val += "Fizz"
	}
	if N%5 == 0 {
		val += "Buzz"
	}
	if val == "" {
		val = strconv.Itoa(N)
	}
	return
}

func getLimit() int {
	return 20
}

/**
Go quirks:
- if clauses can have short statements. This is useful for when the evaluated variable is only needed within the if-scope
*/
func add(requestCount, possessedCount int) int {

	// maxValue is only useful for the following evaluation
	if maxValue := getLimit(); requestCount >= maxValue {
		return maxValue
	}
	return requestCount + possessedCount
}

func tip(total float64) float64 {
	var multiplier float64
	if total < 10 {
		multiplier = 5.0
	} else if total < 15 {
		multiplier = 10.0
	} else if total < 30 {
		multiplier = 12.5
	} else if total < 60 {
		multiplier = 15.0
	} else {
		multiplier = 18.0
	}
	return total * (multiplier / 100)
}

func main() {
	var number int = 15
	var request int = 2
	var total_cost float64 = 67.75
	fmt.Println("#############################################################")
	fmt.Println("# Sandbox")
	fmt.Println("# ========================================================= #")
	fmt.Printf("# Is number %v FizzBuzz? %v\n", number, fizzBuzzButOnlyEvalNumber(number))
	fmt.Printf("# Requesting to add %v into list of %v items: %v items\n", request, number, add(request, number))
	fmt.Printf("# Calculate fake tip on %f is %f\n", total_cost, tip(total_cost))
	fmt.Println("#############################################################")

}
