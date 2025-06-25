// You 're  given two string parameters that can be parsed as float64 value.
// Parse the parameters as float64 values and return the sum of the two values.
// value1 string
// value2 string
// output result float64: the sum of the two values.
// Constraints -> The parameters are strings.Your code should handle errors that occur when they can't
// be parsed as float64 values.
// After parsing to float64 values,each parameter can be either positive or negative.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func calculate(value1 string, value2 string) float64 {
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	return float1 + float2
}

func main() {

	value1 := "10"
	value2 := "5.5"

	result := calculate(value1, value2)
	fmt.Println("sum:", result)

}
