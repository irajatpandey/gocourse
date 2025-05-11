package main

import "fmt"

func main() {

	// Declare an array of integers with a size of 5
	var numbers [5]int

	// Print the initial state of the numbers array (all elements are zero by default)
	fmt.Println(numbers)

	// Assign values to specific elements in the numbers array
	numbers[0] = 100
	numbers[4] = 200

	// Print the updated numbers array
	fmt.Println(numbers)

	// Declare and initialize an array of strings with fruit names
	fruits := [4]string{"Apple", "Orange", "Banana", "Grapes"}
	fmt.Println("Fruits Array:", fruits)

	// Demonstrate array copying
	origialArray := [3]int{1, 2, 3}
	copiedArray := origialArray // copiedArray is a copy of origialArray
	copiedArray[0] = 99         // Modify the first element of copiedArray
	fmt.Println(origialArray)   // Print original array to show it remains unchanged
	fmt.Println(copiedArray)    // Print copied array to show the change

	// Loop through the numbers array using a traditional for loop
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d ", numbers[i]) // Print each element in the numbers array
	}

	// Loop through the fruits array using range
	for index, value := range fruits {
		fmt.Printf("\nIndex: %d Value:%s\n", index, value) // Print index and value of each element
	}

	// Use a blank identifier to ignore the index in the range loop
	for _, value := range fruits {
		fmt.Printf("\nValue:%s\n", value) // Print only the value of each element
	}

	// Compare two arrays for equality
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{4, 5, 6}

	if array1 == array2 {
		fmt.Println("Arrays are equal")
	} else {
		fmt.Println("Arrays are not equal")
	}

	// Declare and initialize a 2D array (matrix)
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix) // Print the entire matrix

	// Loop through the 2D array and print each element
	for _, value := range matrix {
		for _, val := range value {
			fmt.Print(val, " ") // Print each element in the row
		}
		fmt.Println() // New line after each row
	}
}
