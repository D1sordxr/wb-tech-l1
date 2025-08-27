package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	slLen := rand.Intn(1024)
	sl := make([]int, slLen)
	for i := 0; i < slLen; i++ {
		sl[i] = rand.Intn(512)
	}

	target := rand.Intn(512)
	slices.Sort(sl)

	fmt.Printf("Slice: %v\n", sl)
	fmt.Printf("Target: %d\n", target)
	result := binarySearch(sl, target)
	if result != -1 {
		fmt.Printf("Index: %d\n", result)
	} else {
		fmt.Println("Failure")
	}
}

func binarySearch(sl []int, target int) int {
	left, right := 0, len(sl)-1

	for left <= right {
		mid := left + (right-left)/2

		switch {
		case sl[mid] == target:
			return mid
		case sl[mid] < target:
			left = mid + 1
		default: // case sl[mid] > target
			right = mid - 1
		}
	}

	return -1
}
