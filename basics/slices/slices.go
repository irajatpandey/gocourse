package main

import "fmt"

func main() {

	// Declaring and initializing slices

	// Syntax: var sliceName []ElementType = []ElementType{Element1, Element2,...}
	// var numbers1 []int                       // Declares a nil slice of integers
	var numbers2 = []int{1, 2, 3, 4, 5, 6, 7, 8} // Initializes a slice with 8 integers
	// numbers3 := []int{1, 2, 3}              // Shorthand declaration and initialization

	// numbers4 := make([]int, 5)             // Creates a slice of length 5 with zero values

	// Slicing a slice
	slice := numbers2[0:3] // Creates a new slice from index 0 to 2: [1, 2, 3]
	fmt.Println("Sliced (0:3):", slice)

	// Appending elements to a slice
	slice = append(slice, 9) // Appends 9 to the slice: [1, 2, 3, 9]
	fmt.Println("After append:", slice)

	// Copying a slice
	sliceCopy := make([]int, len(numbers2)) // Creates a destination slice with same length
	copy(sliceCopy, numbers2)               // Copies all elements from numbers2 to sliceCopy
	fmt.Println("Copied slice:", sliceCopy)

	// Working with Nil Slices
	// var nilSlice []int // Declares a nil slice (zero length and capacity)

	// Creating and populating a 2D slice
	fmt.Println("\nTwo Dimensional Slice")
	twoD := make([][]int, 3) // Slice with 3 inner slices
	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength) // Each inner slice has length i+1
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j // Populate with sum of indices
		}
		fmt.Println("Row", i, ":", twoD[i])
	}
}
