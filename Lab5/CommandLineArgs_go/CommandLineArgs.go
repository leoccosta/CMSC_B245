/* 
https://cs.brynmawr.edu/Courses/cs245/fall2022/LABS/lab5.html 
"Recall that to get the command line args in Go, you look at the Args variable 
in the os package, as is illustrated in the following very short program."
*/

package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println(os.Args)
}
