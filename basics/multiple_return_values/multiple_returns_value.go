/*
 * Multiple Return Values in Go
 * ===========================
 *
 * Go allows functions to return multiple values, which is a powerful feature
 * that distinguishes it from many other programming languages. This example
 * demonstrates various techniques and patterns for using multiple return values.
 *
 * Key Concepts:
 * ------------
 * 1. Basic multiple returns - Returning two or more values of the same type
 * 2. Named return values - Declaring return variable names in the function signature
 * 3. Error handling pattern - The idiomatic (value, error) return pattern
 * 4. Mixed type returns - Returning values of different types
 * 5. Ignoring returns - Using the blank identifier (_) to ignore unwanted return values
 *
 * Common Use Cases:
 * ---------------
 * - Returning a result and an error
 * - Returning multiple parts of a calculation
 * - Returning optional values along with success indicators
 * - Returning complex data structures with status information
 */
package main

import (
	"fmt"     // For formatted I/O operations
	"strconv" // For string conversion functions
	"strings" // For string manipulation functions
	"time"    // For time-related functions
)

// =====================================================================
// FUNCTION DEFINITIONS
// =====================================================================

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

// parseWithError demonstrates the common Go pattern of returning
// a value along with an error that indicates success or failure.
//
// Parameters:
//   - str: the string to parse as an integer
//
// Returns:
//   - int: the parsed integer value (0 if parsing failed)
//   - error: nil if parsing succeeded, otherwise an error describing the failure
func parseWithError(str string) (int, error) {
	return strconv.Atoi(str)
}

// =====================================================================
// MAIN FUNCTION
// =====================================================================

// Main function demonstrates the usage of multiple return values
func main() {
	printHeader("MULTIPLE RETURN VALUES IN GO")

	// Example 1: Basic multiple return values
	// ----------------------------------------
	printSectionHeader("1. Basic Multiple Returns")
	
	q, r := divide(10, 3)
	fmt.Printf("  divide(10, 3) = quotient: %d, remainder: %d\n", q, r)
	
	// We can also ignore values we don't need using the blank identifier (_)
	onlyQuotient, _ := divide(20, 6)
	fmt.Printf("  Only using quotient: 20 ÷ 6 = %d\n", onlyQuotient)

	// Example 2: Handling errors with multiple return values
	// -----------------------------------------------------
	printSectionHeader("2. Error Handling Pattern")
	
	// Case: Equal values - should return an error
	result, err := compare(6, 6)
	if err != nil {
		fmt.Printf("  compare(6, 6) error: %v\n", err)
	} else {
		fmt.Printf("  compare(6, 6) result: %s\n", result)
	}
	
	// Case: Different values - should return a result with nil error
	result, err = compare(10, 5)
	if err != nil {
		fmt.Printf("  compare(10, 5) error: %v\n", err)
	} else {
		fmt.Printf("  compare(10, 5) result: %s\n", result)
	}

	// Example 3: Using a function with named return values
	// ---------------------------------------------------
	printSectionHeader("3. Named Return Values")
	
	q2, r2 := getDivisionResult(9, 4)
	fmt.Printf("  getDivisionResult(9, 4) = quotient: %d, remainder: %d\n", q2, r2)
	
	// Example 4: Multiple returns with different types
	// -----------------------------------------------
	printSectionHeader("4. Multiple Returns with Different Types")
	
	currentTime, weekday, hour := getTimeInfo()
	fmt.Printf("  Current time: %v\n", currentTime.Format("15:04:05"))
	fmt.Printf("  Day of week: %s\n", weekday)
	fmt.Printf("  Hour of day: %d\n", hour)
	
	// Example 5: The common (value, error) pattern
	// -------------------------------------------
	printSectionHeader("5. Parsing with Error Handling")
	
	// Successful case
	num, err := parseWithError("42")
	if err != nil {
		fmt.Printf("  Failed to parse '42': %v\n", err)
	} else {
		fmt.Printf("  Successfully parsed '42' as: %d\n", num)
	}
	
	// Error case
	num, err = parseWithError("not-a-number")
	if err != nil {
		fmt.Printf("  Failed to parse 'not-a-number': %v\n", err)
	} else {
		fmt.Printf("  Successfully parsed 'not-a-number' as: %d\n", num)
	}
	
	printFooter()
}

// =====================================================================
// HELPER FUNCTIONS FOR FORMATTING OUTPUT
// =====================================================================

// printHeader prints a formatted header for the program output
func printHeader(title string) {
	fmt.Println("=======================================================")
	fmt.Printf("  %s\n", title)
	fmt.Println("=======================================================")
	fmt.Println()
}

// printSectionHeader prints a formatted section header
func printSectionHeader(title string) {
	fmt.Println()
	fmt.Printf("▶ %s\n", title)
	fmt.Println("  " + strings.Repeat("-", len(title)+1))
}

// printFooter prints a formatted footer for the program output
func printFooter() {
	fmt.Println()
	fmt.Println("=======================================================")
	fmt.Println("  End of demonstration")
	fmt.Println("=======================================================")
}