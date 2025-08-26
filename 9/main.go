package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func main() {
	randomNumsCount := rand.IntN(1024)

	sl := make([]int, 0, randomNumsCount)
	mu := sync.RWMutex{}

	for i := 0; i < randomNumsCount; i++ {
		sl = append(sl, rand.IntN(128))
	}

	fmt.Printf("random numbers: %v\n random numbers count: %v\n", sl, len(sl))

	jobs := make(chan int, randomNumsCount/2)
	results := make(chan int, randomNumsCount/2)

	go func() {
		defer close(jobs)

		for _, n := range sl {
			func(num int) {
				mu.RLock()
				defer mu.RUnlock()

				jobs <- num
			}(n)
		}
	}()

	go func() {
		defer close(results)

		for job := range jobs {
			results <- job * job // squared job
		}
	}()

	var totalResults int
	for result := range results {
		fmt.Printf("result: %v\n", result)
		totalResults++
	}
	fmt.Printf("total results: %v\n", totalResults)
}
