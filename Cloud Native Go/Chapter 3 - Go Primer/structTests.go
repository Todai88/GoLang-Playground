package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) toString() {
	fmt.Println(p.name, p.age)
}
func main() {
	var p = Person{}
	p.toString()
	var p2 = Person{"Beth", 25}
	p2.toString()
	var p3 = Person{"", 0}
	p3.toString()
}
