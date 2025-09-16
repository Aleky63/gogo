package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	width, height float64
}
type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Shaper interface {
	Area() float64
	Perimeter() float64
}

func PrintInfo(title string, s Shaper) {
	fmt.Printf("%s\nПлощадь: %.2f\nПериметр: %.2f\n", title, s.Area(), s.Perimeter())
}

func main54() {
	rect := Rectangle{
		width:  15,
		height: 10,
	}
	cir := Circle{
		radius: 8,
	}
	PrintInfo("Прямоугольник:", rect)
	PrintInfo("Круг:", cir)
}

// https://stepik.org/lesson/1500880/step/1?unit=1520997
