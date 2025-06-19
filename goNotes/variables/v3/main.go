package main

import ( 
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)    // bufio package in Go provides buffered I/O — 
	// it wraps around existing io.Reader or io.Writer interfaces to improve performance and
	// add convenience methods. NewReader is a function from there. os.Stdin is a parameter that means
	// input from the console.

    fmt.Print("Enter text: ")
	str, _ := reader.ReadString('\n') // in go if we want to ignore an error, use an underscore for its name, '\n' is a delimiter character or line feed.
	fmt.Println(str)                  // output the string.

	fmt.Print("Enter a number: ")
	str, _ =  reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64) // f variable for a floating point number, naming the error object
	 // with err, using a package strconv(a package for standard library and provides functions for converting between strings 
	 // and basic data types—like integers, floats, and booleans),ParseFloat is a function from this package, passing a 
	 // string object with TrimSpace to avoid any space during user passing before or after entrying
	 // 64 is the precision of the floating point number.                                                      
	if err != nil {
		fmt.Println(err) 
	} else {
		fmt.Println("Value of the number :", f)
	}	             
}

