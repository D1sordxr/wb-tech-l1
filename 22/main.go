package main

import (
	"flag"
	"fmt"
	"math/big"
)

// To use this application you need to write a and b ints.
// Example: go run ./22/main.go --a=1048577 --b=1048577

func main() {
	var a, b int64
	flag.Int64Var(&a, "a", 0, "first number (must be > 2^20)")
	flag.Int64Var(&b, "b", 0, "second number (must be > 2^20)")
	flag.Parse()

	if a <= 1048576 || b <= 1048576 {
		fmt.Println("Error: both numbers must be greater than 2^20 (1,048,576)")
		fmt.Printf("Your input: a=%d, b=%d\n", a, b)
		return
	}

	bigA, bigB := big.NewInt(a), big.NewInt(b)

	fmt.Printf("%d * %d = %s\n", a, b, multiply(bigA, bigB).String())
}

func multiply(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}
