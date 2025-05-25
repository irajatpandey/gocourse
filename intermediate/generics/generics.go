package main	

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

func main() {

	a, b := 1, 2
	a, b = swap(a, b)

	fmt.Println(a, b)

	a1, b1 := "Raju", "Shyam"
	a1, b1 = swap(a1, b1)
	fmt.Println(a1, b1)
}