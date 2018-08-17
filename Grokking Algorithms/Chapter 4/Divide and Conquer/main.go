package main

import (
	"fmt"
	"math/rand"
)

func sum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return numbers[0] + sum(numbers[1:])
}

func count(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return 1 + count(list[1:])
}

func max(list []int, highestNumber int) int {
	if len(list) == 0 {
		return highestNumber
	}
	if list[0] > highestNumber {
		highestNumber = list[0]
	}
	return max(list[1:], highestNumber)
}

func quickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}
	// Pick a random pivot
	pivotIndex := rand.Int() % len(list)

	// Move pivot to the right
	left, right := 0, len(list)-1

	list[pivotIndex], list[right] = list[right], list[pivotIndex]

	for i := range list {
		if list[i] < list[right] {
			list[i], list[left] = list[left], list[i]
			left++
		}
	}

	list[left], list[right] = list[right], list[left]
	quickSort(list[:left])
	quickSort(list[left+1:])
	return list
}

func main() {
	bs := []int{16, 15, 19, 20, 13, 5, 3, 2, 9, 8, 6, 11, 4, 18, 7, 10, 1, 12, 14, 17}
	fmt.Printf("sum: \t%d\n", sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("length: %d\n", count([]int{1, 2, 3}))
	fmt.Printf("max: \t%d\n", max([]int{1, 5, 3}, 0))
	fmt.Printf("bs: \t%d\n", quickSort(bs))
}
