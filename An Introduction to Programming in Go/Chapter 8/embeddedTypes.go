package main

import "fmt"

type Person struct {
	Name string
}
type Android struct {
	Person
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is ", p.Name)
}

func main() {
	p := Person{"Kalle"}
	p.Talk()
	a := new(Android)
	a.Person = Person{"Kalle"}
	a.Talk()
}
