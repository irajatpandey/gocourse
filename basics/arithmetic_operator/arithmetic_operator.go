package main	

import "fmt"
func main() {
	// Variable Declaration
	var a, b int = 10, 20
	var result int
	result = a + b
	fmt.Println("Addition of two numbers is",result)

	result = a - b
	fmt.Println("Subtraction of two numbers is",result)

	result = a * b
	fmt.Println("Multiplication of two numbers is",result)

	result = a / b
	fmt.Println("Division of two numbers is",result) // It will give zero because it's integer division

	result = a % b
	fmt.Println("Modulus of two numbers is",result)

	const p float64 = 22 / 7.0
	fmt.Println(p)	 
	
}