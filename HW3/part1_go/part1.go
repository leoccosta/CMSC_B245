/**
* Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw3.html
* This conclusively proves whether Go uses static or dynamic scoping.
*
* @author lcosta
* Due: October 25, 2022
**/

package main

import "fmt"

var x string
var y string

/**
* a() method takes an x input and prints out x and y (which under static scoping must be globally defined)
*
* @param x string a string that gets printed in the method's print statement
**/
func a(x string) {
	fmt.Printf("x:%v  y:%v\n", x, y)
	// y ends up undefined when this function is called unless the global y is changed
	// meanwhile, in dynamic scope, y would have the value of the last value assigned to the name y
}

func main() {
	// create a value of x to be defined in the main() scope
	x = "defined in main()"
	a(x)

	// "scope 1":
	{
		y := "defined in scope 1"
		fmt.Printf("y: %v\n", y) // proof y was indeed defined in this scope

		a(x) // y value does not carry into function, despite being called within scope 1
	}

	// "scope 2":
	{
		x := "defined in scope 2"
		y = "global y assigned"
		a(x) // x value does carry into function, but only because it is a parameter
	}
}
