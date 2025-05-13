package main

import "fmt"
func main() {

	// ... Ellipsis operator is used to pass a slice as an argument.
	fmt.Println(sum(1, 2, 3, 4, 5))

}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}