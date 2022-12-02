/**
* Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw3.html
* Creates and tests a function that makes and returns an incrementer
*
* @author lcosta
* Due: October 25, 2022
**/

package main

import "fmt"

/**
* main() creates two incrementers and prints their outputs to show they work independently
**/
func main() {
	inc3 := makeIncrementer(3)
	inc7 := makeIncrementer(-7)

	fmt.Printf(inc3())
	fmt.Printf(inc7())
	fmt.Printf(inc3())
	fmt.Printf(inc7())
	fmt.Printf(inc3())
	fmt.Printf(inc7())
	fmt.Printf(inc3())
	fmt.Printf(inc7())
	fmt.Printf(inc3())
	fmt.Printf(inc7())
}

/**
* makeIncrementor() method to create incrementors
*
* @param increment an integer by which the incrementor increments
* @return an incrementor function that returns a string showing its status as an incrementor
**/
func makeIncrementer(increment int) func() string {
	count := 0 - increment

	return func() string {
		count += increment
		return fmt.Sprintf("Incrementor %v: %v\n", increment, count)
	}
}
