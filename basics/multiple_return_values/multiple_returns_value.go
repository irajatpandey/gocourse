// Package main demonstrates different ways to return multiple values from functions in Go
package main

import "fmt"

// divide calculates the quotient and remainder when dividing a by b
// It returns two integers: the quotient and the remainder
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// getDivisionResult performs the same division operation as divide()
// but uses named return values which allows for a "naked" return
// This can improve readability in simple functions
func getDivisionResult(a int, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return // naked return - automatically returns the named return values
}

// compare compares two integers and returns a descriptive string
// If the values are equal, it returns an error instead
// This demonstrates returning a value with an error
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	}
	if a < b {
		return "b is greater than a", nil
	}
	return "", fmt.Errorf("a and b are equal")
}

func main() {
	// Example 1: Basic multiple return values
	q, r := divide(10, 3)
	fmt.Println("divide(10, 3) =", q, r)

	// Example 2: Handling errors with multiple return values
	result, err := compare(6, 6)
	if err != nil {
		fmt.Println("compare(6, 6) error:", err)
	} else {
		fmt.Println("compare(6, 6) result:", result)
	}

	// Example 3: Using a function with named return values
	q2, r2 := getDivisionResult(9, 4)
	fmt.Println("getDivisionResult(9, 4) =", q2, r2)
}