package main

import "fmt"

func main() {
	// Define a string to iterate over
	message := "Hello World"

	// First loop: Iterate over each character in the string using range
	// 'i' is the index of the character, 'v' is the Unicode value of the character
	for i, v := range message {
		fmt.Println(i, v) // Print the index and Unicode value of each character
	}

	// Second loop: Iterate over each character in the string using range
	// This time, print the character itself using the %c format specifier
	for i, v := range message {
		fmt.Printf("%d %c\n", i, v) // Print the index and character
	}
}
