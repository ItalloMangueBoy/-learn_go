package main

import (
	"fmt"
	"math"
)

type GeometricFigure interface {
	Area() float64
}

func TellArea(f GeometricFigure) {
	fmt.Printf("A area desta figura é %f \n", f.Area())
}

type Rectangle struct {
	heigth float64
	width  float64
}

func (f Rectangle) Area() float64 {
	fmt.Println("Calculando area do retangulo")
	return f.heigth * f.width
}

type Circle struct {
	radius float64
}

func (f Circle) Area() float64 {
	fmt.Println("Calculando area do circulo")
	return math.Pow(f.radius, 2) * math.Pi
}


func main() {
	fmt.Println([]interface{}{1, "sdfasd", true})

	TellArea(Rectangle{5, 2})
	TellArea(Circle{5})
}
