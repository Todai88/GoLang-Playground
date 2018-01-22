package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(indexName, indexNumber int) {
	if indexName == 10 || indexNumber == 10 {
		return
	}
	fmt.Println(indexName, ":", indexNumber)
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
	f(indexName, indexNumber+1)
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i, 0)
	}
	// Test
	var input string
	fmt.Scanln(&input)
}
