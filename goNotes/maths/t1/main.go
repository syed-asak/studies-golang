package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC) // we are creating a time object t
	// for the specific moment in time when the go programming was lunched.

	fmt.Printf("Go launched at %s\n", t) // Printf is used for formating here

	n := time.Now()
	fmt.Printf("The time currently is %s\n", n) // shows current date and time.
	fmt.Printf("This object's type is %T\n", n) // finding the variable/object type with Uppercase T

	fmt.Printf(n.Format(time.ANSIC) + "\n") // creating a time object by parsing a string in a format. "\n" is use for append.

	tomorrow := n.AddDate(0, 0, 1)                 // 0 for the year 0 for the month and 1 for the day.
	fmt.Printf(tomorrow.Format(time.ANSIC) + "\n") //tomorrow is one day latter than today

	format := "Mon 2006-02-01"                 // this day is mnemonic and easy to remember.
	fmt.Printf(tomorrow.Format(format) + "\n") //formated value, printf outputs to a target while S printf returns a string to a variable.

}
