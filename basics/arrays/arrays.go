package main

import "fmt"

func main() {

	// var arrayName [size]elementType
	var numbers [5]int

	fmt.Println(numbers)

	numbers[0] = 100
	numbers[4] = 200

	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Orange", "Banana", "Grapes"}
	fmt.Println("Fruits Array:", fruits)

	origialArray := [3]int{1, 2, 3}
	copiedArray := origialArray
	copiedArray[0] = 99
	fmt.Println(origialArray)
	fmt.Println(copiedArray)

	// Loops on Array
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("%d ", numbers[i])
	}

	// using range function

	for index, value := range fruits {
		fmt.Printf("\nIndex: %d Value:%s\n", index, value)
	}
	// underscore is a blank identifier which means we don't need the index variable here. 
	for _, value := range fruits {
		fmt.Printf("\nValue:%s\n", value)
	}
	
	// compare arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{4, 5, 6}

	if array1 == array2 {
		fmt.Println("Arrays are equal")
	} else {
		fmt.Println("Arrays are not equal")
	}

	// 2D Array
	 var matrix[3][3]int = [3][3]int{
		 {1, 2, 3},
		 {4, 5, 6},
		 {7, 8, 9},
	 }
	 fmt.Println(matrix) 

	 for _,value := range matrix {
	 	for _, val := range value {
	 		fmt.Print(val," ")
	 	}
		fmt.Println()
	 }


	 //

}