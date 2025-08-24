package main

import (
	"context"
	"errors"
	"log"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	wg := sync.WaitGroup{}

	forceShutdown := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ticker.C:
				log.Printf("App is doing its work")
			case <-ctx.Done():
				log.Printf("Shutdown signal received, starting graceful shutdown")

				select {
				case <-time.After(3 * time.Second):
					log.Printf("Cleanup operations completed")
				case <-forceShutdown:
					log.Printf("Cleanup interrupted by force shutdown")
				}

				log.Printf("Worker stopped")

				return
			}
		}
	}()

	<-ctx.Done()
	log.Println("Waiting for goroutines to finish...")

	shutdownTimeout := 5 * time.Second
	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	go func() {
		wg.Wait()
		cancel()
	}()

	<-shutdownCtx.Done()

	if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
		log.Printf("Graceful shutdown timeout exceeded, forcing shutdown")
		close(forceShutdown)
	}

	log.Println("Application shutdown successfully")
}
