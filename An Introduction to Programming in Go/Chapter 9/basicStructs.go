package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func basicCircleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func pointerCircleArea(c *Circle) float64 {
	c.r = 10
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}
func main() {
	c := Circle{0, 0, 5}
	fmt.Println(c.area())
	fmt.Println(basicCircleArea(c))
	fmt.Println(pointerCircleArea(&c))
	fmt.Println(basicCircleArea(c))
	fmt.Println(c.area())
}
