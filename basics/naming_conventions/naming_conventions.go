package main

import "fmt"

func main() {
	// PascalCase: Used for naming structs, interfaces, and exported functions
	// Example: MyStruct, MyInterface, MyFunction

	// snake_case: Typically used for filenames and sometimes for variable names in other languages
	// Example: user_id, first_name

	// UPPERCASE: Used for constants, especially when they are exported
	// Example: PI = 3.14

	// mixedCase: Commonly used for local variables and function names
	// Example: javaScript, isValid

	// Constant declaration using UPPERCASE
	const MAXRETRIES = 5 // Represents the maximum number of retries allowed

	// Variable declaration using mixedCase
	employeeID := 100 // Represents the ID of an employee

	// Print the constant and variable values
	fmt.Println(MAXRETRIES)  // Outputs the value of MAXRETRIES
	fmt.Println(employeeID)  // Outputs the value of employeeID
}
