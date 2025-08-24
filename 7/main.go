package main

import (
	"context"
	"log"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type CustomStore struct {
	store map[int]int
	mu    sync.RWMutex
}

func newCustomStore(limit int) *CustomStore {
	return &CustomStore{
		store: make(map[int]int, limit),
		mu:    sync.RWMutex{},
	}
}

func (c *CustomStore) Get(key int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	v, ok := c.store[key]
	return v, ok
}

func (c *CustomStore) Set(key int, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *CustomStore) GetAll() map[int]int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	copiedStore := make(map[int]int)
	for k, v := range c.store {
		copiedStore[k] = v
	}
	return copiedStore
}

func main() {
	const randLimit = 1000

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	customStore := newCustomStore(randLimit)
	syncStore := sync.Map{}

	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				randKey := rand.Intn(randLimit)
				randValue := rand.Intn(randLimit)

				customStore.Set(randKey, randValue)
				syncStore.Store(randKey, randValue)
			}
		}
	}()

	separationLog := func() {
		log.Println("======================================")
	}

	go func() {
		readTicker := time.NewTicker(50 * time.Millisecond)
		defer readTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-readTicker.C:
				randKey := rand.Intn(randLimit)

				if v, exists := customStore.Get(randKey); exists {
					log.Printf("custom store value: %v, key: %v", v, randKey)
				} else {
					log.Printf("custom store value on key %v not found", randKey)
				}

				if v, exists := syncStore.Load(randKey); exists {
					if value, ok := v.(int); ok {
						log.Printf("sync store value: %v, key: %v", value, randKey)
					}
				} else {
					log.Printf("sync store value on key %v not found", randKey)
				}
				separationLog()
			}
		}
	}()

	go func() {
		statsTicker := time.NewTicker(time.Second)
		defer statsTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-statsTicker.C:
				allValues := customStore.GetAll()
				log.Printf("CustomStore items: %d", len(allValues))

				syncCount := 0
				syncStore.Range(func(key, value interface{}) bool {
					syncCount++
					return true
				})
				log.Printf("SyncMap items: %d", syncCount)
				separationLog()
			}
		}
	}()

	<-ctx.Done()
}
