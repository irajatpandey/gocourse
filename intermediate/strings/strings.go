package main

import "fmt"

func main() {
	// Regular double-quoted string
	message := "Hello GO!"

	// Raw string literal (backticks), doesn't interpret escape sequences
	message1 := `Hello \n GO!`

	// Print the messages
	println(message)  // Output: Hello GO!
	println(message1) // Output: Hello \n GO!

	// String properties
	println(len(message))    // Output: Length of the string
	println(message[0])      // Output: Unicode code point (byte value) of first character 'H'

	// String concatenation
	greetings := "Hello"
	name := "Alice"
	println(greetings + " " + name) // Output: Hello Alice

	// String comparison
	str1 := "Apple"
	str2 := "Banana"
	fmt.Println(str1 == str2) // false
	fmt.Println(str1 != str2) // true
	fmt.Println(str1 < str2)  // true  (ASCII of A is less than B)
	fmt.Println(str1 > str2)  // false

	// Iterating over a string (rune-wise)
	for i, char := range message {
		fmt.Printf("Character at %d is %c\n", i, char)
	}

	// Rune count (accurate count of characters, not bytes)
	fmt.Println("Rune count:", len([]rune(message)))

	// String construction
	greetingWithName := greetings + " " + name
	fmt.Println(greetingWithName) // Output: Hello Alice

	// Rune example (character to Unicode)
	var ch rune = 'a'
	fmt.Printf("Unicode point of %c is %d\n", ch, ch) // Output: Unicode point of a is 97
}
