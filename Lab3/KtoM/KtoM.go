/*
* Converts kilometers to miles and miles to kilometers. Part of Lab 3 for CMSC B245.
* @author lcosta
* created: September 14, 2022
 */

package main
import (
	"fmt"
	"strings"
	"os"
	"strconv"
)

func main() {
	f,_ := strconv.ParseFloat(os.Args[1], 32) // still returns a float64 https://yourbasic.org/golang/convert-string-to-float/
	fmt.Printf(KtoM(float32(f), os.Args[2]))
}

func KtoM(i float32, unitFull string) string  {
	unit := strings.ToLower(unitFull[0:1])

	    // convert miles to kilometers
        if unit == "m" {
			// fmt.Sprintf does printf but instead of printing, returns a string
            return fmt.Sprintf("%f kilometers\n", i*1.609)
        }

        // convert kilometers to miles
        if unit == "k" {
			return fmt.Sprintf("%v miles\n", i/1.609)
        }

        return "invalid input\n"
}