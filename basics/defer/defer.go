package main

import "fmt"

func main() {
	process(10) 
}

func process(i int) {
	defer fmt.Println("Deffer i value:", i)
	defer fmt.Println("First Deffered statement executed.")
	defer fmt.Println("Second Deffered statement executed.")
	i++
	fmt.Println("Process function executed.")
	fmt.Println("i value:", i)
}