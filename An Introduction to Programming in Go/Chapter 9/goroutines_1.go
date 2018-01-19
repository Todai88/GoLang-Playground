package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(start, index int) {
	if index == 10 || start == 10 {
		return
	}
	fmt.Println(start, ":", index)
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
	f(start, index+1)
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i, 0)
	}
	var input string
	fmt.Scanln(&input)
}
