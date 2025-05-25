// explaining reassignment concept
package main

import "fmt"

func main() {
	x := 10

	{
		fmt.Println("value of x:", x)
		//new local variable x
		x, y := 20, 30
		fmt.Println("value of x:", x)
		fmt.Println("value of y:", y)
	}
	fmt.Println("value of x:", x)
}
