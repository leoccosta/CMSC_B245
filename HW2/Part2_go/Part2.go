/**
* Practice with slices. Builds on the code in Part1.go to creates a slice containing instances 
* of type Launch (a struct established in Part1) and then subslices of this slice.
*
* @author lcosta
* Due: September 26, 2022
**/

package main
import "fmt" // imports package fmt (pronounced format) for use of Printf() functions

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
		"Date: %v, Vehicle: %v, Flight: %v, Site: %v, Launchpad: %v, Apogee: %v, Category: %v\n",
		stru.date, stru.vehicle, stru.flight, stru.site, stru.launchpad, stru.apogee, stru.category,
	)
} //https://www.geeksforgeeks.org/methods-with-same-name-in-golang/ – func String(stru Launch) string wasn't working

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
	// create a slice containing structs of type Launch and create/add 8 instances of Launch to the slice
	var launches []Launch
	launches = append(launches,
		Launch{Date{1968, 1, 16}, "Voskhod 11A57", "Ya15002- 05", "NIIP-53", "5", 441, "Sat"},
		Launch{Date{1958, 5, 22}, "Nike Cajun", "-", "WI", "1", 240, "Balloon"},
		Launch{Date{1983, 1, 14}, "Kappa 9M", "K-9M-76", "KASC", "c", 349, "Aeron/Io"},
		Launch{Date{1966, 8, 25}, "R-16U", "G 22603-13T", "NIIP-5", "o", 1210, "OT"},
		Launch{Date{2003, 8, 8}, "Zenit-3SL", "SL10/DM-SL-10L", "KLA", "c", 36090, "Sat"},
		Launch{Date{2012, 9, 12}, "Terrier Lynx", "-", "WI", "c", 300, "Target"},
		Launch{Date{1978, 12, 19}, "Proton-K/DM", "295-02", "NIIP-5", "2",  51018 , "Sat"},
		Launch{Date{2018, 10, 11}, "Sineva", "VMF RF", "BLA", "f", 1000, "OT"},
	) # https://stackoverflow.com/questions/34846848/how-to-break-a-long-line-of-code-in-golang 

	// create a subslice that has only the first 5 items of launches
	subSlice1 := launches[:5]
	fmt.Printf("subSlice1 = %v\n", subSlice1)

	// create a subslice that has only the last 5 items of launches
	subSlice2 := launches[len(launches)-5:]
	fmt.Printf("subSlice2 = %v\n", subSlice2)

	// create a subslice that has the first 3 and last 2 items of launches
	subSlice3 := launches[:3]
	subSlice3 = append(subSlice3, launches[len(launches)-2:]...)
	fmt.Printf("subSlice3 = %v\n", subSlice3)

	// tests if the first item in the original slice and in subSlice1 are equal
	fmt.Printf("test: launches[0] == subSlice1[0] is %v\n", launches[0] == subSlice1[0])

	// changes the value of a field in the struct in position 1 of the original slice
	launches[0].apogee = 500

	// shows that the first item in the original slice and in subSlice1 are still equal
	fmt.Printf("test: launches[0] == subSlice1[0] is %v\n", launches[0] == subSlice1[0])

	// shows the value of the field changed in launches[0] has also changed in subSlice1[0]
	fmt.Printf("subSlice1[0].apogee = %v\n", subSlice1[0].apogee)

}

/**
* Helpful Links:
* https://cs.brynmawr.edu/Courses/cs245/fall2022/Topic03/slic_go/slic.go.txt 
* https://cs.brynmawr.edu/Courses/cs245/fall2022/Topic03/slisli_go/slisli.go.txt
**/