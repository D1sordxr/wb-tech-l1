package main

import "fmt"

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 19.9}

	store := map[int][]float32{}

	for _, t := range temperatures {
		groupKey := int(t/10) * 10
		store[groupKey] = append(store[groupKey], t)
	}

	fmt.Printf("Grouped temperatures: %v\n", store)
}
