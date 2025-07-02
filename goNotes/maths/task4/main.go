package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //buffered reader to read user input from the keyboard

	fmt.Println("Enter numbers separated by spaces (e.g., 10 5 3 8 2):")
	line, err := reader.ReadString('\n') //Read characters until you see a newline (\n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Split the input string by spaces
	numStrings := strings.Fields(line) // Automatically splits by any whitespace

	// Convert to integers
	var numbers []int              //declares a slice of integers named numbers
	for _, s := range numStrings { // numStrings is a slice of strings (like ["12", "45", "7", "3"]).
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", s)
			return
		}
		numbers = append(numbers, num)
	}

	// Sort the numbers
	sort.Ints(numbers)

	// Print the result
	fmt.Println("Sorted numbers:", numbers)
}
