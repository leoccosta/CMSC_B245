/*
https://cs.brynmawr.edu/Courses/cs245/fall2022/LABS/lab5.html
"A random number generator can be perceived of as consisting of two parts:
a very long array of random numbers and a pointer to the current location in the array.
(It does not work this way, but it is a useful metaphor.) When you get a random number,
the function returns the value at the current position in the very long array and
increments the position. So, when you want a random number, you first have to create
the array and set where in the array you want to read. In Go, this requires two commands
as shown in the following little example. Once those two commands have been executed,
you use the thing returned to actually get random numbers. Note that if you do the exact 
same thing to create the array and choose a location in that array, then you will get the 
exact same sequence of random numbers. For this reason, it is common why you want reasonably 
random numbers to create the array and set the position using the current system clock time 
(as shown below).
*/

package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Printf("%d\n", r1.Intn(100))
}