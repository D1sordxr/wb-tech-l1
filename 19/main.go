package main

import (
	"fmt"
	"strings"
)

func main() {
	data := strings.Trim(strings.Repeat("大丈夫、", 10), "、") + "!"
	fmt.Printf("Data: %s\n", data)
	fmt.Printf("Reversed data: %s\n", reverseString(data))

}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
