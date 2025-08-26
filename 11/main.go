package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	targets := make(map[int]struct{})

	for _, v := range a {
		targets[v] = struct{}{}
	}

	result := make([]int, 0, len(a)/2)

	for _, v := range b {
		if _, ok := targets[v]; ok {
			result = append(result, v)
		}
	}

	fmt.Printf("Intersection result: %v\n", result)

	funcResult := intersection[int](a, b)

	fmt.Printf("Intersection func result: %v\n", funcResult)
}

func intersection[T comparable](a, b []T) []T {
	targets := make(map[T]struct{}, len(a))
	for _, v := range a {
		targets[v] = struct{}{}
	}

	result := make([]T, 0, len(a)/2)
	for _, v := range b {
		if _, ok := targets[v]; ok {
			result = append(result, v)
		}
	}

	return result
}
