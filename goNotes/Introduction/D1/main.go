package main

import "fmt"

func main() {
	// "\n" is an escape sequence that creates a new line
	x := "apple\norange"
	fmt.Println("Using \\n as an escape sequence:")
	fmt.Println(x)

	// `\n` is treated as raw characters, no escape
	y := `apple\norange`
	fmt.Println("\nUsing raw string (backticks):")
	fmt.Println(y)
}
