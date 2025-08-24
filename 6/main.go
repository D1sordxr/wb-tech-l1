package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

const workTime = time.Millisecond * 500

func main() {
	examples := []func(){
		firstExample,
		secondExample,
		thirdExample,
		fourthExample,
		fifthExample,
		sixthExample,
		seventhExample,
	}

	wg := sync.WaitGroup{}

	wg.Add(len(examples))
	for _, example := range examples {
		go func(e func()) {
			defer wg.Done()
			e()
		}(example)
	}

	wg.Wait()
}

func firstExample() {
	stop := make(chan struct{}, 1)
	done := make(chan struct{}, 1)

	go func() {
		defer close(done)
		defer log.Println("First example: stopped via channel")

		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	time.AfterFunc(workTime, func() {
		close(stop)
	})

	<-done
}

func secondExample() {
	ctx, cancel := context.WithTimeout(context.Background(), workTime)
	defer cancel()

	go func() {
		defer log.Println("Second example: stopped via context")

		for {
			select {
			case <-ctx.Done():
				return
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()
}

func thirdExample() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("Third example: stopped via Goexit")

		timer := time.NewTimer(workTime)
		defer timer.Stop()

		<-timer.C
		runtime.Goexit()
	}()

	wg.Wait()
}

func fourthExample() {
	var (
		stop atomic.Bool
		wg   sync.WaitGroup
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("Fourth example: stopped via condition")

		for {
			if !stop.Load() {
				time.Sleep(10 * time.Millisecond)
				continue
			}
			return
		}
	}()

	time.AfterFunc(workTime, func() {
		stop.Store(true)
	})

	wg.Wait()
}

func fifthExample() {
	dataChan := make(chan int, 1)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("Fifth example: stopped via channel close")

		for {
			_, ok := <-dataChan
			if !ok {
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	time.AfterFunc(workTime, func() {
		close(dataChan)
	})

	wg.Wait()
}

func sixthExample() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if r := recover(); r != nil {
				log.Println("Sixth example: stopped via panic/recover:", r)
			}
		}()

		timer := time.NewTimer(workTime)
		defer timer.Stop()

		<-timer.C
		panic("planned shutdown")
	}()

	wg.Wait()
}

func seventhExample() {
	stop := make(chan struct{})
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer log.Println("Seventh example: stopped via timer")

		timer := time.NewTimer(workTime)
		defer timer.Stop()

		for {
			select {
			case <-timer.C:
				return
			case <-stop:
				return
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	time.AfterFunc(workTime, func() {
		close(stop)
	})

	wg.Wait()
}
