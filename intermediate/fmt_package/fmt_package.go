package main

import "fmt"

func main() {

	var name string
	var age int

	fmt.Print("Enter your name and age: ")
	fmt.Scanf("%s %d", &name, &age)
	fmt.Printf("Hello %s, you are %d years old.\n", name, age)


	fmt.Println("Enter your name and age: ")
	fmt.Scanln(&name, &age)

	

} 