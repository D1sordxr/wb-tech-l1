package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	const workersCount = 2

	arr := []int{2, 4, 6, 8, 10}

	type jobData struct {
		value int
		index int
	}

	jobs := make(chan jobData, workersCount)
	wg := sync.WaitGroup{}

	results := make([]int, len(arr))
	mu := sync.Mutex{}

	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(id int) {
			defer wg.Done()

			log.Printf("Worker %d starting...\n", id)
			defer log.Printf("Worker %d stopped.\n", id)

			for j := range jobs {
				log.Printf("Worker %d started work at %d...\n", id, j.value)

				time.Sleep(1 * time.Second)

				log.Printf("Worker %d finished work at %d! Result is %d.\n", id, j.value, j.value*j.value)

				mu.Lock()
				results[j.index] = j.value * j.value
				mu.Unlock()
			}
		}(i)
	}

	log.Printf("All workers started. Ready to work!\n")
	time.Sleep(1 * time.Second)

	for i, j := range arr {
		jobs <- jobData{value: j, index: i}
	}

	close(jobs)
	wg.Wait()

	log.Printf("All workers finished. Job is done!\n")

	for k, v := range results {
		fmt.Printf("%d squared is %d\n", arr[k], v)
	}
}
