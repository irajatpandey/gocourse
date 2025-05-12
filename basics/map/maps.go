// Package main demonstrates Go's map data structure and its operations
package main

import (
	"fmt"  // For formatted I/O operations
	"maps" // For map utility functions (Go 1.21+)
)

func main() {
	// ==========================================
	// MAP CREATION AND INITIALIZATION
	// ==========================================
	
	// Declaration syntax: var mapVariable map[keyType]valueType
	// Creating a map using make() - allocates and initializes a hash map
	myMap := make(map[string]int)
	fmt.Println("Empty map:", myMap) // Output: map[]
	
	// Adding key-value pairs to the map
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4
	
	fmt.Println("Map after adding elements:", myMap)
	
	// ==========================================
	// ACCESSING AND MODIFYING MAP ELEMENTS
	// ==========================================
	
	// Accessing a value by key
	// If the key exists, it returns the corresponding value
	fmt.Println("Value for key 'three':", myMap["three"]) // Output: 3
	
	// Deleting a key-value pair from the map
	delete(myMap, "two")
	fmt.Println("Map after deleting 'two':", myMap)
	
	// ==========================================
	// CHECKING FOR KEY EXISTENCE
	// ==========================================
	
	// Method 1: Using the "comma ok" idiom to check if a key exists
	// This returns both the value and a boolean indicating if the key exists
	value, ok := myMap["five"] // 'five' doesn't exist in the map
	if !ok {
		fmt.Println("Key 'five' doesn't exist in the map")
	} else {
		fmt.Println("Value for key 'five':", value)
	}
	
	// Method 2: When you only need to check existence without using the value
	// Use underscore (_) to ignore the value
	_, exists := myMap["five"]
	if !exists {
		fmt.Println("Key 'five' doesn't exist in the map")
	} else {
		fmt.Println("Key 'five' exists:", exists)
	}
	
	// ==========================================
	// ALTERNATIVE MAP INITIALIZATION
	// ==========================================
	
	// Creating a map with initial values using map literal syntax
	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("Map created with literal syntax:", myMap2)
	
	// ==========================================
	// MAP COMPARISON
	// ==========================================
	
	// Comparing two maps using maps.Equal() function (Go 1.21+)
	// This checks if both maps have the same key-value pairs
	if maps.Equal(myMap, myMap2) {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}
	
	// ==========================================
	// ITERATING OVER MAPS
	// ==========================================
	
	// Method 1: Iterating over both keys and values
	fmt.Println("\nIterating over map (keys and values):")
	for k, v := range myMap {
		fmt.Printf("  %s -> %d\n", k, v)
	}
	
	// Method 2: Iterating over values only
	fmt.Println("\nIterating over map (values only):")
	for _, v := range myMap {
		fmt.Println("  Value:", v)
	}
	
	// Method 3: You can also iterate over keys only (not shown)
	// for k := range myMap { ... }
	
	// ==========================================
	// NIL MAPS
	// ==========================================
	
	// Creating a nil map (declared but not initialized)
	var myMap3 map[string]string
	
	// Checking if a map is nil
	if myMap3 == nil {
		fmt.Println("\nWarning: myMap3 is nil - cannot add elements until initialized")
		// myMap3["key"] = "value" // This would cause a runtime panic!
	} else {
		fmt.Println("myMap3 is not nil")
	}
	
	// To make a nil map usable, initialize it with make():
	// myMap3 = make(map[string]string)
	
	// ==========================================
	// MAP LENGTH
	// ==========================================
	
	// Getting the number of key-value pairs in a map
	fmt.Println("\nNumber of elements in myMap:", len(myMap))

	//length of map
	fmt.Println(len(myMap))
}