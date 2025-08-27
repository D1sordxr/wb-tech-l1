package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	const duration = time.Second * 2

	sleepFunctions := []func(time.Duration){
		withEndMessage(Sleep),
		withEndMessage(ChanSleep),
		withEndMessage(TimerSleep),
	}

	for i, fn := range sleepFunctions {
		fmt.Printf("=== Sleep function #%d ===\n", i+1)
		fn(duration)
	}
}

func Sleep(duration time.Duration) {
	target := time.Now().Add(duration)
	for time.Now().Before(target) {
		runtime.Gosched()
	}
}

func ChanSleep(duration time.Duration) {
	park := make(chan struct{})
	time.AfterFunc(duration, func() { close(park) })
	<-park
}

func TimerSleep(duration time.Duration) {
	ticker := time.NewTimer(duration)
	defer ticker.Stop()
	<-ticker.C
}

func withEndMessage(sleepFunc func(time.Duration)) func(time.Duration) {
	return func(duration time.Duration) {
		start := time.Now()
		defer func() {
			elapsed := time.Since(start)
			fmt.Printf("Sleep end. Duration: %v, Elapsed: %v\n", duration, elapsed)
		}()

		sleepFunc(duration)
	}
}
