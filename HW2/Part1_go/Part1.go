/**
* Establishes and tests Launch struct and Date struct and respective String() functions
* to organize information about a rocket launch.
*
* @author lcosta
* Due: September 26, 2022
**/

package main
import "fmt"  // imports package fmt (pronounced format) for use of Printf() functions

// The struct Launch holds information about a rocket launch. 
type Launch struct {
	date Date			// Date struct that holds the year, month, day of the launch

	// Variable descriptions (comments) taken from assignment details.
	vehicle string 		// the type of rocket
	flight string		// A descriptor of the flight. Usually unintelligible.
	site string			// The location of the rocket launch
	launchpad string	// the place within the location
	apogee int 			// max distance from earth
	category string		// I am sure that this is meaningful, just not to me.
}

/**
* String() method for a Launch struct
*
* @param stru the Launch struct being converted to a string
* @return contents of a Launch struct as a string
**/
func (stru Launch) String() string {
	return fmt.Sprintf(
		"Date: %v, Vehicle: %v, Flight: %v, Site: %v, Launchpad: %v, Apogee: %v, Category: %v",
		stru.date, stru.vehicle, stru.flight, stru.site, stru.launchpad, stru.apogee, stru.category,
	)
} // https://www.geeksforgeeks.org/methods-with-same-name-in-golang/ – func String(stru Launch) string wasn't working

// The struct Date holds information about a date (day, month, year)
type Date struct {
	year int 			// the year
	month int 			// month 1 == January, etc
	day int 			// day of month. First day of month is 1
}

/**
* String() method for a Date struct
*
* @param stru the Date struct being converted to a string
* @return contents of a Date struct as a string in the format month/day/year.
**/
func (stru Date) String() string {
	m := fmt.Sprintf("%v", stru.month)
	d := fmt.Sprintf("%v", stru.day)
	if stru.month < 10 {
		m = "0" + m
	}
	if stru.day < 10 {
		d = "0" + d
	}
	return fmt.Sprintf("%v/%v/%v", m, d, stru.year)
}

func main() {
	// Create 3 instances of Launch and store them in 3 seperate variables
	a := Launch{Date{1968, 01, 16}, "Voskhod 11A57", "Ya15002- 05", "NIIP-53", "5", 441, "Sat"}
	b := Launch{Date{1958, 05, 22}, "Nike Cajun", "-", "WI", "1", 240, "Balloon"}
	c := Launch{Date{1983, 01, 14}, "Kappa 9M", "K-9M-76", "KASC", "c", 349, "Aeron/Io"}

	// Test String() method on Launch a
	fmt.Printf("Launch a: %v\n", a.String())

	// Test whether the three Launches are programatically equivalent (they should not be)
	fmt.Printf("a == b: %v\nb == c: %v\nc == a: %v\n", a == b, b == c, a == c)

	// Create a copy of Launch a; Go creates a new instance whereas Java would create a reference (thus, an alias)
	d := a

	// Test that they have the same contents (should return true)
	fmt.Printf("a == d: %v\n", a == d)

	// Change a field in Launch d and show it is no longer equivalent to a
	d.apogee = 500
	fmt.Printf("a == d: %v\n", a == d)
}