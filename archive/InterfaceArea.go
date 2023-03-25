package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

func main() {
	sq := square{sideLength: 10}
	tr := triangle{base: 10, height: 20}
	printArea(sq)
	printArea(tr)
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func (s square) area() float64 {
	return (s.sideLength * s.sideLength)
}

func (s triangle) area() float64 {
	return (0.5 * s.base * s.height)
}
