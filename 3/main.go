package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// To use this application you need to write workers count.
// Example: go run ./3/main.go --workers=50

func main() {
	var workersCount int
	flag.IntVar(&workersCount, "workers", 0, "Number of workers to use.")
	flag.Parse()

	if workersCount < 1 {
		log.Fatal("Must provide at least one worker")
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int, 100)
	wg := sync.WaitGroup{}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(jobs)

		for {
			select {
			case <-ctx.Done():
				log.Println("Data generator stopped")
				return
			case <-ticker.C:
				jobs <- rand.Int()
			}
		}
	}()

	wg.Add(workersCount)
	for i := 0; i < workersCount; i++ {
		go func(id int) {
			defer wg.Done()
			log.Printf("Worker %d started", id)
			defer log.Printf("Worker %d stopped", id)

			for job := range jobs {
				randMs := rand.Intn(1000)
				time.Sleep(time.Duration(randMs) * time.Millisecond)
				log.Printf("Worker %d received %d!\n", id, job)
			}
		}(i)
	}

	<-ctx.Done()

	done := make(chan struct{})

	go func() {
		wg.Wait()
		close(done)
	}()

	<-done
	log.Println("Application shutdown successfully")
}
