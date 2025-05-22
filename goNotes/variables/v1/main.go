// Declaration and Interpolation examples
package main

import (
	"errors"
	"fmt"
)

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := "42"

	fmt.Println(str1, str2, str3)
	stringlength, err := 0, errors.New("simulated print error")
	// Simulate error	stringlength, err := fmt.Println("The value is", aNumber) // here generally stringlength has the 
	// of the string and any error goes to err, here we commented the line to run else part.
	if err == nil {
		fmt.Println("String length:", stringlength)
	} else {
		fmt.Println("An error occurred while printing:", err)
	}

	fmt.Printf("Value of number: %v\n", aNumber) //print the value of the aNumber
	fmt.Printf("Data type: %T\n", aNumber)       //print the datatype of the variable aNumber

}
