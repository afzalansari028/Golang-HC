package main

import "fmt"

// Define an interface called "Shape" with a single method called "Area"
type Shape interface {
	Area() float64
}

// Define a concrete type called "Rectangle" that implements the "Shape" interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the "Area" method for the "Rectangle" type
func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

// Define another concrete type called "Circle" that also implements the "Shape" interface
type Circle struct {
	Radius float64
}

// Implement the "Area" method for the "Circle" type
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// A function that takes an interface as an argument
func PrintArea(s Shape) {
	fmt.Printf("Area of shape is: %f\n", s.Area())
}

func main() {
	// Create a Rectangle and a Circle
	r := Rectangle{Width: 5, Height: 10}
	c := Circle{Radius: 7}

	// fmt.Println("***", r.Area())

	// Call the PrintArea function with each shape
	PrintArea(r)
	PrintArea(c)
}
