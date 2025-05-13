package main

import "fmt"

func main() {
	// panic(interface{})
	process(10)
	process(-3)
}

func process(input int) {
	defer fmt.Println("Deffered 1")
	defer fmt.Println("Deffered 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be positive")
		//defer fmt.Println("After Panic") // this will not executed because of panic
	}
	fmt.Println("processing...", input)

}