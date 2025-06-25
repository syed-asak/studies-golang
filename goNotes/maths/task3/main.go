package main

import (
	"fmt"
)

func fibonacci(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Enter number of Fibonacci terms: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Please enter a positive integer")
		return
	}

	fmt.Printf("Fibonacci sequence (%d terms):\n", n)
	fibonacci(n)
}
