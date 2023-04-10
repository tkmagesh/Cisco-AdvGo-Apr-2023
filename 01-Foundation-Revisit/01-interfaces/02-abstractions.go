package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

// utility functions
// ver-1.0
/*
func PrintArea(c Circle) {
	fmt.Println("Area :", c.Area())
}
*/

// ver-2.0
/*
func PrintArea(x interface{}) {
	switch o := x.(type) {
	case Circle:
		fmt.Println("Area :", o.Area())
	case Rectangle:
		fmt.Println("Area :", o.Area())
	default:
		fmt.Println("Argument is not an object with Area() method")
	}
}
*/

// ver-3.0
// The assumption is that "Circle" & "Rectangle" has "Area() methods"
/*
func PrintArea(x interface{}) {
	switch o := x.(type) {
	case Circle:
		fmt.Println("Area :", o.Area())
	case Rectangle:
		fmt.Println("Area :", o.Area())
	default:
		fmt.Println("Argument is not an object with Area() method")
	}
}
*/

// ver-4.0
/*
func PrintArea(x interface{}) {
	switch o := x.(type) {
	case interface{ Area() float32 }: //Any object that has an Area() method:
		fmt.Println("Area :", o.Area())
	default:
		fmt.Println("Argument is not an object with Area() method")
	}
}
*/

// ver-5.0
/*
func PrintArea(x interface{ Area() float32 }) {
	fmt.Println("Area :", x.Area())
}
*/

// ver-6.0
type AreaFinder interface{ Area() float32 }

func PrintArea(x AreaFinder) {
	fmt.Println("Area :", x.Area())
}

type Rectangle struct {
	Length  float32
	Breadth float32
}

func (r Rectangle) Area() float32 {
	return r.Length * r.Breadth
}

func main() {
	c := Circle{Radius: 12}
	PrintArea(c)

	r := Rectangle{Length: 10, Breadth: 12}
	// fmt.Println("Area :", r.Area())
	PrintArea(r)

	// PrintArea(100)

	// Implement the above concepts for "Perimeter()"
	// Perimeter of circle = 2 * Pi * r
	// Perimeter of Rectangle = 2 * (length + breadth)
}
