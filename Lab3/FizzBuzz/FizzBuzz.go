/*
* Prints the numbers 1-100, except:
* if the number is divisible by 3 print "Fizz"
* if the number is divisible by 5 print "Buzz"
* [forgot this last condition:] if the number is divisible by 3 and 5 print "FizzBuzz."
* @author lcosta
* created: September 14, 2022
 */

package main // REQUIRED

func main() {
	FizzBuzz()
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 {
			println("Fizz");
		} else if i % 5 == 0 {
			println("Buzz");
		} else {
			println(i);
		}
	}
}

