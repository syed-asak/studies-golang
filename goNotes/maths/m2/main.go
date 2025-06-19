package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100          // using = as we have define sum before.
	fmt.Printf("The Sum is now %v\n\n", sum) // %v is a placeholder used with two linefeed \n

	fmt.Println("The value of PI is", math.Pi) // value of PI

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference) // .2f means 2 character after decimal point in the output.

}
