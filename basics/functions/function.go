package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	// func <name> (<params>) (<return type>) {}
	fmt.Println(add(10, 20))

	// anonymous function

	greet := func() {
		fmt.Println("Hello Anonymous Function!")
	}

	greet()

	result := applyOperation(5, 6, add)
	fmt.Println("Result of operation:", result)

	result2 := createMultiplicationFunction(3)
	fmt.Println(result2(4)) // 3*4 = 12
}


// Function that take as function as an argument

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function returning a function
func createMultiplicationFunction(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}