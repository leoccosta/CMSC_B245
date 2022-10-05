/*
https://cs.brynmawr.edu/Courses/cs245/fall2022/LABS/lab5.html
"strconv.Atoi() converts a string containing a number to an int. It returns 
two values, the number itself, and an error value (in err) in case there it 
was unable to convert the string into an int. The error value returned is nil 
if successful. The program below tests to see if err is nil before printing 
the value of the integer command line argument. For instance:"
*/

package main
import ( "fmt"
         "strconv"
)
func main() {
	v1,err1 := strconv.Atoi("3")
	if err1==nil {
		fmt.Printf("%d\n", v1)
	}
	v2,e2 := strconv.Atoi("this fails")
	if e2==nil {
		fmt.Printf("%d\n", v2)
	} else {
		fmt.Printf("ERR: %v\n", e2)
	}
}