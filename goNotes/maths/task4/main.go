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
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter numbers separated by spaces (e.g., 10 5 3 8 2):")
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Split the input string by spaces
	numStrings := strings.Fields(line) // Automatically splits by any whitespace

	// Convert to integers
	var numbers []int
	for _, s := range numStrings {
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
