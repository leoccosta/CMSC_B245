/*
* https://cs.brynmawr.edu/Courses/cs245/fall2022/LABS/lab5.html
* Collecting Cards in Candy
*
* @author lcosta
* Date: October 05, 2022
**/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Takes input of # of pictures (i.e., 35) on the command line:
	picNum, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("Bad input: %v\n", err)
	} else {
		sum := 0
		trials := 0
		// run simulations:
		startT := time.Now()
		for trials <= 1000000 {
			sum += collect(picNum)
			trials++
		}
		// print out the expected number of candies (XXX) that kids will purchase to collect all 35 cards
		fmt.Printf("trials: %v, average: %v, runtime: %v\n", (trials - 1), sum/(trials-1), time.Since(startT))
	}
}

/**
* Simulate collecting cards
* @param c number of different cards we want to collect
* @return the total number of candies bought in order to collect all 35 cards
**/
func collect(c int) int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	purchases := 0
	numColl := 0
	collected := make([]bool, c)

	for numColl < c {
		purchases++
		rand := r1.Intn(c)
		if collected[rand] == false {
			collected[rand] = true
			numColl++
		}
		//fmt.Printf("collection: %v\n", collected) //prints status of collection
	}
	return purchases
}
