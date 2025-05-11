package main

import "fmt"

func forLoop() {
	// Simple for loop to print numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i) // Print the current value of i
	}

	// Declare a slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6}

	// Iterate over the slice using range
	for index, value := range numbers {
		// Print the index and value of each element in the slice
		fmt.Println(index, value)
	}

	// Loop to print numbers from 10 down to 1
	for i := range 10 {
		// Print the countdown from 10
		fmt.Println(10 - i)
	}
}


func whileLoop() {
	i := 1
	for i <= 5 {
		fmt.Println("Iteration:", i)
		i++
	}
}
func main() {
	forLoop() 
	whileLoop()
}
