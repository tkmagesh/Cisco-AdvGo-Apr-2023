package main

import (
	"abs-app/models"
	"fmt"
)

func DisplayArea(x interface{ Area() float32 }) {
	fmt.Println("Area : ", x.Area())
}

func main() {
	c := models.Circle{Radius: 12}
	DisplayArea(c)
	/*
		utils.PrintArea(c)
		utils.PrintPerimeter(c) */

	r := models.Rectangle{Length: 10, Breadth: 12}
	DisplayArea(r)
	/*
		utils.PrintArea(r)
		utils.PrintPerimeter(r)
	*/
}
