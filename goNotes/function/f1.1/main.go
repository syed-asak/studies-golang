// program to explain parameters and returns

package main
import "fmt"

func main()  {
	/* declare local variable*/
	var x int = 50
	var y int = 40
	var sub_value int

	/* calling the function to get the sub of a values */
	sub_value = sub(x,y)

	fmt.Printf("Subtraction value is : %d\n", sub_value)

}

 /* function returns the sum of two numbers */

   func sub(num1, num2 int) int
   var result int
   result = num1 - num2
   return result
   }

