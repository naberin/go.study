package main

import (
	"fmt"
)

func printCountdown(from int) {
	for from >= 0 {
		fmt.Printf("# Countdown: %v\n", from)
		from--

	}
}

func printSumFromZeroTo(to int) {
	var sum int
	for from := 0; from < to; from++ {
		total := sum
		sum += from
		fmt.Printf("# %v + %v = %v\n", total, from, sum)
	}
	fmt.Printf("# %v\n", sum)
}

func printFromTo(from, to int) {
	var i int = from
	for ; i <= to; i++ {
		fmt.Printf("# %v\n", i)
	}
}

func main() {
	fmt.Println("#############################################################")
	fmt.Println("# Sandbox")
	fmt.Println("# ========================================================= #")
	from := 10
	to := 15
	fmt.Println("# Printing from", from, "to", to)
	printFromTo(from, to)
	fmt.Println("# Printing sum from", 0, "to", to)
	printSumFromZeroTo(to)
	fmt.Println("# Printing countdown from", from)
	printCountdown(from)
	fmt.Println("#############################################################")
}
