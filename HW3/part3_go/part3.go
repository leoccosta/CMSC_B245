/**
* Instructions: https://cs.brynmawr.edu/Courses/cs245/fall2022/HWS/hw3.html
* Creates and tests a function that creates iterators to iterate through a
* slice of type int and tests the iterators and their abilities and limitations
*
* @author lcosta
* Due: October 25, 2022
**/

package main

import "fmt"

func main() {

	//Show that if you create two iterators, they do not interact:
	a := []int{1, 2, 3, 4, 5, 6, 7}
	iterator1 := iteratorGenerator(a)
	iterator2 := iteratorGenerator(a)

	for val, err := iterator1(); err == false; val, err = iterator1() {
		fmt.Printf("Iterator 1: value: %v error: %v\n", val, err)

		val2, err2 := iterator2()
		if err2 {
			break
		}

		fmt.Printf("Iterator 2: value: %v error: %v\n\n", val2, err2)
	}

	//Create an iterator then change the contents of the slice that the iterator works on
	iterator3 := iteratorGenerator(a)
	a[1] = 20

	//then use the iterator to print the slice's contents
	for val, err := iterator3(); err == false; val, err = iterator3() {
		fmt.Printf("Iterator 3: value: %v error: %v\n", val, err)
	}
	fmt.Printf("\n")

	//Create an iterator; then append new items to the slice that the iterator works on;
	a = []int{1, 2, 3, 4, 5, 6, 7}
	iterator4 := iteratorGenerator(a)
	a = append(a, 8)
	//then use the iterator to print the slice's contents
	for val, err := iterator4(); err == false; val, err = iterator4() {
		fmt.Printf("Iterator 4: value: %v error: %v\n", val, err)
	}
}

/**
* defines a function called iterator
*
* @param none
* @return an integer (for the value of the next element in an int slice
		  a boolean (for an error variable dependent on whether the value returned is valid)
**/
//
type iterator func() (int, bool)

/**
* iteratorGenerator() method generates and returns an iterator
*
* @param ourSlice the integer slice to be iterated over
* @return function of type iterator
**/
func iteratorGenerator(ourSlice []int) iterator {
	index := -1

	return func() (int, bool) {
		index++
		if index < len(ourSlice) {
			return ourSlice[index], false
		} else {
			return 0, true
		}
	}
}
