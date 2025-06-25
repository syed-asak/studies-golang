//A simple calculator

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(a float64, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", operator)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin) // Sets up a reader to get input from the terminal.

	fmt.Print("Enter first number: ")                              // Prompts the user for input.
	input1, _ := reader.ReadString('\n')                           // Reads it as a string, (_) in Go is called the blank identifier, and it's used to ignore a return value that you don't care about.
	num1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64) // Trims extra spaces or newline characters and Converts it to a float64
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	operator := strings.TrimSpace(op)

	fmt.Print("Enter second number: ")
	input2, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
