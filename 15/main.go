package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// var justString string - использование глобальной переменной плохая практика

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
		var justString string // локальная переменная

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
