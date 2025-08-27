package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
)

func main() {
	slLen := rand.Intn(64) + 1
	sl := make([]int, slLen)
	for i := 0; i < slLen; i++ {
		sl[i] = rand.Intn(512)
	}
	idx := rand.Intn(len(sl))

	fmt.Printf("Index %d. Slice: %v. Slice length: %d. Slice capacity: %d.\n", idx, sl, len(sl), cap(sl))

	newSl := make([]int, slLen-1)
	copy(newSl[:idx], sl[:idx])
	copy(newSl[idx:], sl[idx+1:])

	fmt.Printf("Result: %v. Result length: %d. Result capacity: %d.\n", newSl, len(newSl), cap(newSl))
}

func testFunc() { // copy(newSl, append(sl[:idx], sl[idx+1:]...)) - removed due alloc
	test := func() {
		slLen := rand.Intn(64) + 1
		sl := make([]int, slLen)
		for i := 0; i < slLen; i++ {
			sl[i] = rand.Intn(512)
		}
		idx := rand.Intn(len(sl))

		fmt.Printf("Index %d. Slice: %v. Slice length: %d. Slice capacity: %d.\n", idx, sl, len(sl), cap(sl))

		newSl := make([]int, slLen-1)
		copy(newSl, append(sl[:idx], sl[idx+1:]...))

		fmt.Printf("Result: %v. Result length: %d. Result capacity: %d.\n", newSl, len(newSl), cap(newSl))
	}

	done := make(chan os.Signal, 1)
	go func() {
		for {
			test()
		}
	}()

	signal.Notify(done, os.Interrupt)
	<-done
}
