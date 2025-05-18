package main

import "fmt"

func init(){
	fmt.Println("init function executed1...!")
}
func init(){
	fmt.Println("init function executed2...!")
}
func init(){
	fmt.Println("init function executed3...!")
}


func main() {
	fmt.Println("Inside the main function")
}