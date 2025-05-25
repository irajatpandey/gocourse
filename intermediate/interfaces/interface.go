package main

import (
    "fmt"
    "math"
)

// geometry interface defines methods that any geometric shape must implement
type geometry interface {
    area() float64
    perimeter() float64
}

// rectangle struct represents a rectangle with width and height
type rectangle struct {
    width, height float64
}

// circle struct represents a circle with a radius
type circle struct {
    radius float64
}

// area calculates the area of a rectangle (implements geometry)
func (r rectangle) area() float64 {
    return r.height * r.width
}

// area calculates the area of a circle (implements geometry)
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

// perimeter calculates the perimeter of a rectangle (implements geometry)
func (r rectangle) perimeter() float64 {
    return 2 * (r.width + r.height)
}

// perimeter calculates the perimeter of a circle (implements geometry)
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}

// measure takes any geometry type and prints its details, area, and perimeter
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println("Area:", g.area())
    fmt.Println("Perimeter:", g.perimeter())
}

func main() {
    // Create a rectangle and a circle
    r := rectangle{width: 3, height: 4}
    c := circle{radius: 3}

    // Pass both to measure, demonstrating interface usage
    measure(r)
    measure(c)
}