// Package main demonstrates variadic functions in Go
package main

import "fmt"

// sum is a variadic function that accepts any number of integers
// and returns their sum. The '...' syntax makes this a variadic function.
func sum(nums ...int) int {
	// Inside the function, 'nums' is treated as a slice of integers
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// join is a variadic function with a regular parameter
// The variadic parameter must be the last parameter
func join(separator string, words ...string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += separator
		}
		result += word
	}
	return result
}

func main() {
	// Basic usage: passing multiple arguments directly
	fmt.Println("Sum of 1, 2, 3, 4, 5:", sum(1, 2, 3, 4, 5))
	
	// Variadic functions can be called with any number of arguments
	fmt.Println("Sum of 10, 20:", sum(10, 20))
	fmt.Println("Sum with no arguments:", sum())
	
	// Passing a slice to a variadic function using the '...' operator
	numbers := []int{5, 10, 15, 20, 25}
	fmt.Println("Sum of slice:", sum(numbers...))
	
	// Example with a regular parameter and variadic parameter
	fmt.Println("Joined words:", join(", ", "apple", "banana", "cherry"))
}