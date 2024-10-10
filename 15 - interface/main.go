package main

import (
	"fmt"
	"math"
)

// GeometricFigure is an interface represents any geometric figure methods
type GeometricFigure interface {
	Area() float64
}

// TellArea prints an message indicating the given GeometricFigure area  
func TellArea(f GeometricFigure) {
	fmt.Printf("A area desta figura é %f \n", f.Area())
}

// Rectangle represents an rectangle figure 
type Rectangle struct {
	heigth float64
	width  float64
}

//  Area compute and  returns the referenced rectangle area
func (f Rectangle) Area() float64 {
	return f.heigth * f.width
}

// Circle represents an circle figure 
type Circle struct {
	radius float64
}

//  Area compute and  returns the referenced cicle area
func (f Circle) Area() float64 {
	return math.Pow(f.radius, 2) * math.Pi
}


func main() {
	fmt.Println([]interface{}{1, "sdfasd", true})

	TellArea(Rectangle{5, 2})
	TellArea(Circle{5})
}
