package main

import (
	"fmt"
)

func Calculate(x int) int {
	fmt.Println("The val is", x, "and x*x is", (x * x))
	return x * x
}

func main() {
	fmt.Println("Hello there")
}
