package main

import (
	"fmt"
)

func main() {
	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence()) 
}

func adder() func() int {
	i := 0 
	fmt.Println("Previous Value of i:", i)
	return func () int{
		i++
		fmt.Println("Current value of i:", i)
		return i
	}
}
