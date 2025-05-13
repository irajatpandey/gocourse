/*
Package main demonstrates multiple return values in Go.

This example showcases different techniques for returning multiple values from functions,
which is a powerful feature in Go that allows functions to return more than one result.
Common use cases include:
- Returning a result and an error
- Returning multiple parts of a calculation
- Returning optional values along with success indicators
*/
package main

import (
	"fmt"
	"time"
)

// divide calculates the quotient and remainder when dividing a by b.
// It demonstrates the basic syntax for returning multiple values.
//
// Parameters:
//   - a: the dividend (number being divided)
//   - b: the divisor (number to divide by)
//
// Returns:
//   - int: the quotient (a / b)
//   - int: the remainder (a % b)
//
// Note: This function will panic if b is 0 (division by zero).
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// getDivisionResult performs the same division operation as divide()
// but uses named return values which allows for a "naked" return.
//
// Named return values:
//   - quotient: stores the result of integer division
//   - remainder: stores the modulus result
//
// This approach can improve readability in simple functions and
// is particularly useful when the return values have clear meanings.
func getDivisionResult(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // naked return - automatically returns the named return values
}

// compare compares two integers and returns a descriptive string.
// If the values are equal, it returns an error instead.
//
// This demonstrates the idiomatic Go pattern of returning (value, error),
// where the error is nil when the operation succeeds.
//
// Parameters:
//   - a, b: the integers to compare
//
// Returns:
//   - string: a description of the comparison result
//   - error: nil if a comparison was made, or an error if values are equal
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	}
	if a < b {
		return "b is greater than a", nil
	}
	return "", fmt.Errorf("a and b are equal")
}

// getTimeInfo returns current time information with multiple values.
// This demonstrates returning multiple values of different types.
//
// Returns:
//   - time.Time: the current time
//   - string: the weekday name
//   - int: the hour of day (0-23)
func getTimeInfo() (time.Time, string, int) {
	now := time.Now()
	return now, now.Weekday().String(), now.Hour()
}

// Main function demonstrates the usage of multiple return values
func main() {
	// Example 1: Basic multiple return values
	// ----------------------------------------
	q, r := divide(10, 3)
	fmt.Printf("Example 1: divide(10, 3) = quotient: %d, remainder: %d\n", q, r)
	
	// We can also ignore values we don't need using the blank identifier (_)
	onlyQuotient, _ := divide(20, 6)
	fmt.Printf("          Only using quotient: 20 ÷ 6 = %d\n", onlyQuotient)

	// Example 2: Handling errors with multiple return values
	// -----------------------------------------------------
	fmt.Println("\nExample 2: Error handling with multiple returns")
	
	// Case: Equal values - should return an error
	result, err := compare(6, 6)
	if err != nil {
		fmt.Printf("          compare(6, 6) error: %v\n", err)
	} else {
		fmt.Printf("          compare(6, 6) result: %s\n", result)
	}
	
	// Case: Different values - should return a result with nil error
	result, err = compare(10, 5)
	if err != nil {
		fmt.Printf("          compare(10, 5) error: %v\n", err)
	} else {
		fmt.Printf("          compare(10, 5) result: %s\n", result)
	}

	// Example 3: Using a function with named return values
	// ---------------------------------------------------
	q2, r2 := getDivisionResult(9, 4)
	fmt.Printf("\nExample 3: getDivisionResult(9, 4) = quotient: %d, remainder: %d\n", q2, r2)
	
	// Example 4: Multiple returns with different types
	// -----------------------------------------------
	currentTime, weekday, hour := getTimeInfo()
	fmt.Printf("\nExample 4: Current time info:\n")
	fmt.Printf("          Time: %v\n", currentTime.Format("15:04:05"))
	fmt.Printf("          Day: %s\n", weekday)
	fmt.Printf("          Hour: %d\n", hour)
}