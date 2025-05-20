package main

import "fmt"

// ✅ Struct Definition
type Rectangle struct {
	Length  float64
	Breadth float64
}

// ✅ Method with Value Receiver
// This method calculates area but does NOT modify the struct
// 'r' is just a copy of Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

// ✅ Method with Pointer Receiver
// This method modifies the original struct values
func (r *Rectangle) DoubleDimensions() {
	r.Length *= 2
	r.Breadth *= 2
}

// ✅ Method with Value Receiver: Perimeter Calculation
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Breadth)
}

// ✅ Method with Pointer Receiver: Reset dimensions to zero
func (r *Rectangle) Reset() {
	r.Length = 0
	r.Breadth = 0
}

func main() {
	// 🔸 Create a Rectangle object
	rect := Rectangle{Length: 5, Breadth: 3}

	// 🔹 Call value receiver method - does not change struct
	area := rect.Area()
	fmt.Println("Area:", area) // Output: Area: 15

	// 🔹 Call perimeter method
	perimeter := rect.Perimeter()
	fmt.Println("Perimeter:", perimeter) // Output: Perimeter: 16

	// 🔹 Call pointer receiver method - modifies struct
	rect.DoubleDimensions()
	fmt.Println("After Doubling - Length:", rect.Length, "Breadth:", rect.Breadth)
	// Output: After Doubling - Length: 10 Breadth: 6

	// 🔹 Call Reset method to zero the fields
	rect.Reset()
	fmt.Println("After Reset - Length:", rect.Length, "Breadth:", rect.Breadth)
	// Output: After Reset - Length: 0 Breadth: 0

	// ✅ Extra: You can also call pointer methods using value type, Go auto-converts!
	anotherRect := Rectangle{Length: 2, Breadth: 2}
	anotherRect.DoubleDimensions() // even though it's not a pointer
	fmt.Println("Another Rect Doubled - Length:", anotherRect.Length, "Breadth:", anotherRect.Breadth)
}
