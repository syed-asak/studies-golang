// difference betn "\n" and '\n'.
package main

import "fmt"

func main() {
// newline sequence is treated as a special value	
	x := "apple\norange"
	fmt.Println(x)

// newline sequence is treated as two raw characters

	y:= 'apple\norange'
	fmt.Print\n(y)

}
