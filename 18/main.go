package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var (
		counter       atomic.Uint64
		wg            sync.WaitGroup
		numGoroutines = rand.Intn(512) + 1
	)

	fmt.Printf("Starting %d goroutines\n", numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			incrementTicker := time.NewTicker(time.Microsecond)
			defer incrementTicker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-incrementTicker.C:
					counter.Add(1)
				}
			}
		}(i)
	}

	go func() {
		statusTicker := time.NewTicker(time.Second * 5)
		defer statusTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-statusTicker.C:
				fmt.Printf("Counter = %d\n", counter.Load())
			}
		}
	}()

	<-ctx.Done()

	wg.Wait()

	fmt.Printf("Final counter = %d\n", counter.Load())
}
