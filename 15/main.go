package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

func someFunc() string {
	v := createHugeString(1 << 10)

	result := make([]byte, 100) // создание новой строки
	copy(result, v[:100])

	return string(result) // возврат новой строки без ссылок на v
}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ticker := time.NewTicker(time.Millisecond * 50)
	defer ticker.Stop()

	for {
		var justString string

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			justString = someFunc()
			fmt.Printf("data: %s\nlength: %d\n", justString, len(justString))
		}
	}
}

func createHugeString(size int) string {
	const chars = "abcdefghijklmnopqrstuvwxyz"

	sl := make([]byte, size)
	for i := range size {
		sl[i] = chars[rand.Intn(len(chars))]
	}

	return string(sl)
}
