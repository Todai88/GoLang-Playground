package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x, y float64
}

type Shape interface {
	area() float64
	perimeter() float64
}

type MultiShape struct {
	shapes []Shape
}

func basicCircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func pointerCircleArea(c *Circle) float64 {
	c.r = 10
	return math.Pi * c.r * c.r
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	return r.x + r.y
}

func (c Circle) perimeter() float64 {
	return (math.Pi * 2) * c.r
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.x + r.y)
}
func totalAreas(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	r := Rectangle{2, 5}
	fmt.Println(r.area())
	fmt.Println(totalAreas(&c, &r))
	c = Circle{0, 0, 5}
	r = Rectangle{5, 5}
	shapes := MultiShape{[]Shape{c, r}}
	fmt.Println(shapes.area())
	fmt.Println(c.perimeter())
	fmt.Println(r.perimeter())
}
