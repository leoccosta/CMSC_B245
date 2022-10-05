/*
* Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/LABS/lab4.html
* @author lcosta
* created: September 21, 2022
*/

/*
* Into Terminal:
* go mod init aaa
* go run Sorter.go
*/

package main
import "fmt"

type myStruct struct {
	num int
	flt float32
	str string
} // only need braces for defining new types when the type is a struct

func (s myStruct) String() string {
	return fmt.Sprintf("Int: %v Float: %f String: %s\n", s.num, s.flt, s.str)
} // https://cs.brynmawr.edu/Courses/cs245/fall2021/LABS/student_go/student.go.txt

type structures []myStruct // any instance of structures contains a slice of type myStruct
type dollar float32 // why? it's just a float32? but it documents itself – by the name we know it's currency

func main() {
	var mySlice structures 

	// way 1 of initializing a struct:
	var strct1 myStruct
	strct1.num = 14
	strct1.flt = 0.1
	strct1.str = "what's up"

	// way 2 of initializing a struct:
	strct2 := myStruct{5, 5.2, "hi"}
	strct3 := myStruct{-29, 8.3, "ola"}
	// TODO create 7 more instances

	// println(strct1.String())
	// println(strct2.String())
	// println(strct3.String())

	mySlice = append(mySlice, strct1, strct2, strct3)
	fmt.Println(mySlice)

}