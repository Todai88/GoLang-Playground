package main

import "fmt"

type attackMove struct {
	defense, attack int
}

type sword struct {
	attackMove
	name      string
	twoHanded bool
}

type gun struct {
	attackMove
	name             string
	bulletsRemaining int
}

type weapon interface {
	Wield() bool
}

func (g gun) Wield() bool {
	fmt.Println("Wielded ", g.name, ". Dealt: ", g.attack, "dmg")
	return true
}

func (s sword) Wield() bool {
	fmt.Println("Wielded ", s.name, ". Dealt: ", s.attack, "dmg")
	return true
}

func wielder(w weapon) bool {
	fmt.Println("Wielding...")
	return w.Wield()
}
func main() {
	s := sword{attackMove: attackMove{attack: 10, defense: 5}, name: "bastard sword", twoHanded: true}
	wielder(s)
}
