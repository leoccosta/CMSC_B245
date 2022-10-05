package main

/*
* Small illustration of function returning a function in Go
* and closures
*
* @author gtowell
* Created: Aug 22, 2021
* Modified: Sep 22, 2021
 */
import "fmt"

func main() {
	b6 := a(6) 
	b42 := a(42)
	fmt.Printf("b6 :%v\n", b6(5)) //12
	fmt.Printf("b42:%v\n", b42(7)) //50
	fmt.Printf("b6 :%v\n", b6(5))
	fmt.Printf("b42:%v\n", b42(7))
	fmt.Printf("b6 :%v\n", b6(25))
	fmt.Printf("b42:%v\n", b42(-11))
	var bb funcType
	{
		aa := 7
		bb = func(mult int) int {
			return aa*mult
		}
		fmt.Printf("bb7:%v\n", bb(7))
		cc(bb)
	}
	//bb is still in scope
	fmt.Printf("bb7:%v\n", bb(7))

	//locally defined function is used rather than global
	//always in Go, local shadows global
	cc := func(height int) int {
		return 7*4*height
	}
	fmt.Printf("cc:%v\n", cc(3))
} // functions b6 and b42; b6 and b42 created by a function (a)

// Not strictly necessary, but defining types for funtions being passed, or
// returned make the code more readable, IMHO
type funcType func(int) int

/*
* Function that creates and returns a function.
* The created function simply returns the value passed into this
* function plus, the value passed into the created function.
* See the main method for examples
*
* @param int the value to increment
* @return the created function
 */
func a(outerParam int) funcType {
	count := 0
	b := func(incr int) int {
		count++
		return count + outerParam + incr
	}
	return b
}

// this is exactly the same as the function a
// just slightly differently phrased
func a2 (outerParam int) func(int) int {
	count := 0
	return func(incr int) int {
		count++
		return count + outerParam + incr
	}
}

// Here we have the function receiving another function
// as an argument. Since functions are first class in Go
// this works too.
func cc(fun funcType) {
	fmt.Printf("bb7 in cc:%v\n", fun(6))
}