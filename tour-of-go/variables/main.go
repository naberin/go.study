package main

import (
	"fmt"
	"math"
)

func defaultValues() {
	var num int
	var floats float64
	var boolean bool
	var strings string
	fmt.Printf("# %v %f %v %q\n", num, floats, boolean, strings)
}

/**
Go quirks:
- there is a difference between := and =
- = is only used for assignment, hence the need for var
- := is used for initialization and assignment
*/
func initializedWithoutAutomaticTypes() {
	var number, boolean, strings = 1, true, "#"
	fmt.Println(strings, number, boolean)

}

/**
Go quirks:
- var and := cannot be used together
*/
func intializedWithoutVar() {
	number := 1
	strings, boolean := "#", true
	fmt.Println(strings, number, boolean)
}

func numericTypes() {
	var maxInt64 int64 = math.MaxInt
	var maxUInt64 uint64 = 1<<64 - 1
	fmt.Println("# max-sized-int", maxInt64)
	fmt.Println("# max-unsigned-int", maxUInt64)
}

/**
Go quirks:
- variable types and value can be printed with printF's %T and %v substitution variables
*/
func blockVariablesWithFormatting() {
	var (
		opA int  = 1
		opB uint = math.MaxInt
	)
	fmt.Printf("# type: %T, value: %v\n", opA, opA)
	fmt.Printf("# type: %T, value: %v\n", opB, opB)
}

/**
Go quirks:
- variables must be explicitly casted
*/
func convertingValues() {
	number := 5
	var half float64 = float64(number) / 2
	var long uint64 = uint64(half)

	fmt.Printf("# %v %f %v\n", number, half, long)
}

func constantValues() {
	const INTERVAL = 2
	fmt.Printf("# Scheduled interval is set to %v\n", INTERVAL)
}

func main() {
	fmt.Println("#############################################################")
	fmt.Println("# Variables")
	fmt.Println("# ========================================================= #")
	defaultValues()
	initializedWithoutAutomaticTypes()
	intializedWithoutVar()
	numericTypes()
	blockVariablesWithFormatting()
	constantValues()
	fmt.Println("#############################################################")
}
