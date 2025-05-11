package main
import (
	"fmt"
)
const MAXTRIES = 5
const PI = 3.14
func main() {
	
	fmt.Println(MAXTRIES)
	fmt.Println(PI)

	// Declare multiple constants in one line.
	const (
		MONDAY = 1
		TUESDAY = 2
		WEDNESDAY = 3
		THURSDAY = 4
		FRIDAY = 5
		SATURDAY = 6
		SUNDAY = 7
	)
	fmt.Println(MONDAY, TUESDAY)
	fmt.Println(WEDNESDAY, THURSDAY)
	fmt.Println(FRIDAY, SATURDAY)
	fmt.Println(SUNDAY)
}