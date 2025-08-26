package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := rand.Intn(100)
	sl := make([]int, count)
	for i := 0; i < count; i++ {
		sl[i] = rand.Intn(1000)
	}

	fmt.Printf("Before: %v.\n", sl)
	start := time.Now()
	result := quickSort(sl)
	elapsed := time.Since(start)
	fmt.Printf("After: %v. Time: %s.\n", result, elapsed)
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivotIdx := len(nums) / 2
	pivot := nums[pivotIdx]

	var (
		left   = make([]int, 0, len(nums)/2)
		middle = make([]int, 0, 1)
		right  = make([]int, 0, len(nums)/2)
	)

	for _, num := range nums {
		if num < pivot {
			left = append(left, num) // Все элементы меньше pivot
		} else if num > pivot {
			right = append(right, num) // Все элементы больше pivot
		} else {
			middle = append(middle, num) // Все элементы равные pivot
		}
	}

	return append(append(quickSort(left), middle...), quickSort(right)...)
}
