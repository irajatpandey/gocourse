package main

import "fmt"

func main() {
	// Integer formatting Verbs
	// %b - base 2
	// %d - base 10
	// %+d - base 10 with sign
	// %o - base 8
	// %O - base 8 (with leading zero)
	// %x - base 16 (lowercase)
	// %X - base 16 (uppercase)
	// %#x - base 16 with leading 0x
	// %4d - Pad with spaces (width 4, right-aligned)
	// %-4d - Pad with spaces (width 4, left-aligned)
	// %04d - Pad with zeros (width 4, right-aligned)

	int := 18
	fmt.Printf("Integer formatting Verbs")
	fmt.Printf("Binary: %b\n", int)
	fmt.Printf("Decimal: %d\n", int)
	fmt.Printf("Decimal with sign: %+d\n", int)
	fmt.Printf("Octal: %o\n", int)
	fmt.Printf("Octal with leading zero: %O\n", int)
	fmt.Printf("Hexadecimal (lowercase): %x\n", int)
	fmt.Printf("Hexadecimal (uppercase): %X\n", int)
	fmt.Printf("Hexadecimal with leading 0x: %#x\n", int)
	fmt.Printf("Padded with spaces (width 4, right-aligned): %4d\n", int)
	fmt.Printf("Padded with spaces (width 4, left-aligned): %-4d\n", int)
	fmt.Printf("Padded with zeros (width 4, right-aligned): %04d\n", int)

	// Strings formatting Verbs

	// %s - string
	// %q - quoted string
	// %x - hexadecimal representation of bytes
	// %X - hexadecimal representation of bytes (uppercase)
	// %c - character (rune)
	// %p - pointer
	// %T - type of the value
	// %v - default format
	// %+v - default format with field names
	// %#v - Go-syntax representation of the value
	// %t - boolean
	str := "Hello, World!"
	fmt.Printf("String formatting Verbs")
	fmt.Printf("String: %s\n", str)
	fmt.Printf("Quoted string: %q\n", str)
	fmt.Printf("Hexadecimal representation of bytes: %x\n", str)
	fmt.Printf("Hexadecimal representation of bytes (uppercase): %X\n", str)
	fmt.Printf("Character (rune): %c\n", str[0])
	fmt.Printf("Pointer: %p\n", &str)
	fmt.Printf("Type of the value: %T\n", str)
	fmt.Printf("Default format: %v\n", str)
	fmt.Printf("Default format with field names: %+v\n", str)
	fmt.Printf("Go-syntax representation of the value: %#v\n", str)
	fmt.Printf("Boolean: %t\n", true)
	fmt.Printf("Boolean: %t\n", false)

	
	// %e - scientific notation (lowercase)

}