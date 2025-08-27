package main

import "fmt"

func main() {
	firstStr, secondStr, thirdStr := "abcd", "abCdefAaf", "aabcd"
	results := func(args ...bool) []bool { return append([]bool{}, args...) }
	fmt.Printf("%v", results(isUnique(firstStr), isUnique(secondStr), isUnique(thirdStr)))
}

func isUnique(str string) bool {
	seen := make(map[rune]struct{})

	for _, char := range str {
		lowerChar := char
		if char >= 'A' && char <= 'Z' {
			lowerChar = char + 32 // ASCII to lower
		}

		if _, ok := seen[lowerChar]; ok {
			return false
		}
		seen[lowerChar] = struct{}{}
	}

	return true
}
