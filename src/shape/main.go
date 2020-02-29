package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sidelength float64
}

type triangle struct {
	heigth float64
	base   float64
}

func main() {
	squareArea := square{10.64}

	triangleArea := triangle{5, 6}

	printArea(squareArea)
	printArea(triangleArea)

}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.heigth
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
