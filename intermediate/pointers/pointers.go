package main


func modifyValue(ptr *int) {
	*ptr++

	println("Inside function", *ptr)
}
func main() {
	var ptr *int
	var a int = 10
	ptr = &a
	println(*ptr) // prints the value of a
	println(ptr)  // prints the address of a

	modifyValue(ptr)
}