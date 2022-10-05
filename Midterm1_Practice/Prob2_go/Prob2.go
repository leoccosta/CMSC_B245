package main
import "fmt"

// // Slice:
// func main() {
// 	bb := make([]interface{}, 2)
// 	bb[0] = "hello"
// 	bb[1] = 9
// 	fmt.Println(bb)
// }

// Array:
func main() {
	var bb [2]interface{}
	bb[0]="hello"
	bb[1] = 9
	fmt.Println(bb)
}