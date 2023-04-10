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

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
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

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Length + r.Breadth)
}

type PerimeterFinder interface{ Perimeter() float32 }

func PrintPerimeter(x PerimeterFinder) {
	fmt.Println("Perimeter :", x.Perimeter())
}

// ver-1.0
/*
func PrintStats(x interface {
	interface{ Area() float32 }
	interface{ Perimeter() float32 }
}) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}
*/

//ver-2.0
/*
func PrintStats(x interface {
	Area() float32
	Perimeter() float32
}) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}
*/

//ver-3.0
/*
func PrintStats(x interface {
	AreaFinder
	PerimeterFinder
}) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}
*/

//ver-4.0

type ShapeStatsFinder interface {
	AreaFinder
	PerimeterFinder
}

func PrintStats(x ShapeStatsFinder) {
	PrintArea(x)      // interface { Area() float32 }
	PrintPerimeter(x) // interface { Perimeter() float32 }
}

func main() {
	c := Circle{Radius: 12}
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintStats(c)

	r := Rectangle{Length: 10, Breadth: 12}
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintStats(r)

}
