package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dieRoll(size int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(size) + 1
}

func main() {
	fmt.Printf("Rolled a die of size %d, result: %d\n", 6, dieRoll(6))
}
