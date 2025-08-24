package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	const (
		tickerInterval = 50 * time.Millisecond
		appDuration    = 5 * time.Second
	)

	start := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	time.AfterFunc(appDuration, func() {
		cancel()
	})

	jobs := make(chan int, 100)
	wg := sync.WaitGroup{}

	ticker := time.NewTicker(tickerInterval)
	defer ticker.Stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(jobs)

		for {
			select {
			case <-ticker.C:
				jobs <- rand.Intn(2025)
			case <-ctx.Done():
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for job := range jobs {
			log.Printf("Received number %d", job)
		}
	}()

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
}
