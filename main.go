package main

import (
	"fmt"
)

// Testing action

func Calculate(x int) int {
	fmt.Println("The val is", x, "and x*x is", (x * x))
	return x * x
}

func main() {
	fmt.Println("Hello there")
}
