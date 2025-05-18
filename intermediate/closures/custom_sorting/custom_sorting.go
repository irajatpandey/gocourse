package main

import "fmt"

// multiplier returns a closure that multiplies its input by a given factor.
// It maintains a state (value) that is updated with each call.
func multiplier(factor int) func(x int) int {
	value := 1 // Initial value for the closure's state
	return func(x int) int {
		value *= x          // Update the state by multiplying with x
		return value * factor // Return the product of the updated state and factor
	}
}

func main() {
	// Create a closure that multiplies by 2
	multBy2 := multiplier(2)
	fmt.Println(multBy2(3)) // Outputs: 6 (2 * 3)
	fmt.Println(multBy2(5)) // Outputs: 30 (2 * 3 * 5)

	// Create a new closure that multiplies by 4
	multBy4 := multiplier(4)
	fmt.Println(multBy4(6)) // Outputs: 24 (4 * 6)
}
