package main

import "fmt"

func main() {
	a, b := 10, 20
	fmt.Printf("Before: a: %v, b: %v\n", a, b)

	a, b = b, a
	fmt.Printf("After: a: %v, b: %v\n", a, b)

	swapNums(&a, &b)
	fmt.Printf("Rollback: a: %v, b: %v\n", a, b)

	swapArithmetic(&a, &b)
	fmt.Printf("Arithmetic: a: %v, b: %v\n", a, b)

	swapXOR(&a, &b)
	fmt.Printf("XOR: a: %v, b: %v\n", a, b)
}

func swapNums(a, b *int) {
	*a, *b = *b, *a
}

func swapArithmetic(a, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

func swapXOR(a, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}
