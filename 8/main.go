package main

import (
	"fmt"
)

func main() {
	var (
		num      int64 = 5 // 101 в двоичной системе
		position uint  = 0
		bit      uint  = 0
	)

	fmt.Printf("Исходное число: %d (%b)\n", num, num)

	res, prev := setBit(num, position, bit) // 4
	fmt.Printf(
		"После установки %d в позицию %d: Result: %d (%b). Previous: %d (%b).\n",
		bit, position, res, res, prev, prev,
	)

	bit++
	res, prev = setBit(num, position, bit) // 5 (без изменений)
	fmt.Printf(
		"После установки %d в позицию %d: Result: %s. Previous: %s.\n",
		bit, position, res.String(), prev.String(),
	)

	// with using previous results
	res, prev = setBit(res.Int64(), 3, 1) // 13
	fmt.Printf("Result: %s. Previous: %s.\n", res.String(), prev.String())
	res, prev = setBit(res.Int64(), 0, 0) // 12
	fmt.Printf("Result: %s. Previous: %s.\n", res.String(), prev.String())
	res, prev = setBit(res.Int64(), 2, 0) // 8
	fmt.Printf("Result: %s. Previous: %s.\n", res.String(), prev.String())
}

type (
	result   int64
	previous int64
)

func setBit(n int64, i uint, bit uint) (result, previous) {
	if bit > 1 {
		panic("invalid bit")
	}

	if bit == 1 {
		return result(n | (1 << i)), previous(n)
	} else {
		return result(n &^ (1 << i)), previous(n)
	}
}

func (r result) Int64() int64 {
	return int64(r)
}

func (p previous) Int64() int64 {
	return int64(p)
}

func (r result) String() string {
	return fmt.Sprintf("%d (%b)", r.Int64(), r.Int64())
}

func (p previous) String() string {
	return fmt.Sprintf("%d (%b)", p.Int64(), p.Int64())
}
