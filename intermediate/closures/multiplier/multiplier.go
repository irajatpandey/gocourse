package main

import "fmt"

func multiplier(factor int) func(x int) int {
	value := 1
	return func(x int) int {
		value *= x
		return value * factor
	}
}

func main() {
	multBy2 := multiplier(2)
	fmt.Println(multBy2(3))
	fmt.Println(multBy2(5))
	multBy4 := multiplier(4)
	fmt.Println(multBy4(6))

}