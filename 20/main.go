package main

import "fmt"

func main() {
	str := "sun dog snow"
	reversed := reverseWords(str)
	fmt.Println(reversed) // "snow dog sun"
}

func reverseWords(s string) string {
	runes := []rune(s)

	reverseRunes(runes, 0, len(runes)-1)

	reverseEachWord(runes)

	return string(runes)
}

func reverseRunes(runes []rune, start, end int) {
	for left, right := start, end; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
}

func reverseEachWord(runes []rune) {
	var (
		wordStart = 0
		length    = len(runes)
	)

	for i := 0; i <= length; i++ {
		if i == length || runes[i] == ' ' {
			reverseRunes(runes, wordStart, i-1)
			wordStart = i + 1
		}
	}
}
