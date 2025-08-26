package main

import "fmt"

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, v := range sl {
		set[v] = struct{}{}
	}

	fmt.Printf("Set: %v\n", set)
}
