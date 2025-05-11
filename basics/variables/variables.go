package main

import (
	"fmt"
)
var middleName = "Cane"
func main() {
	// Variable declarations (commented out examples):
	// var age int                // Declares an integer variable 'age' with a default value of 0
	// var name string = "John"   // Declares a string variable 'name' initialized with "John"
	// var name1 = "John"         // Declares a string variable 'name1' using type inference
	// var isMarried bool         // Declares a boolean variable 'isMarried' with a default value of false

	// count := 10 // Short variable declaration, only allowed within functions
	lastName := "Doe" // Declares and initializes a string variable 'lastName' with "Doe"

	// Default Values of Variables:
	// Numeric Types: 0
	// Boolean Type: false
	// String Type: "" (empty string)

	// ***** Scope of Variables *****
	firstName := "John" // 'firstName' is a local variable, accessible only within 'main'
	printName()         // Calls the 'printName' function
	fmt.Println(firstName) // Prints the local 'firstName' variable from 'main'
	fmt.Println(middleName) // Error: 'middleName' is not declared in this scope
	fmt.Println(lastName)   // Prints the 'lastName' variable, accessible within 'main'
}

func printName() {
	firstName := "Michael" // 'firstName' is a local variable, accessible only within 'printName'
	fmt.Println(firstName) // Prints the local 'firstName' variable from 'printName'
}
