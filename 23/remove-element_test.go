package main

import (
	"math/rand"
	"testing"
)

func BenchmarkRemoveElementImproved(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slLen := rand.Intn(64) + 1
		sl := make([]int, slLen)
		for i := 0; i < slLen; i++ {
			sl[i] = rand.Intn(512)
		}
		idx := rand.Intn(len(sl))

		newSl := make([]int, slLen-1)
		copy(newSl[:idx], sl[:idx])
		copy(newSl[idx:], sl[idx+1:])
	}
}

func BenchmarkRemoveElementOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slLen := rand.Intn(64) + 1
		sl := make([]int, slLen)
		for i := 0; i < slLen; i++ {
			sl[i] = rand.Intn(512)
		}
		idx := rand.Intn(len(sl))

		newSl := make([]int, slLen-1)
		copy(newSl, append(sl[:idx], sl[idx+1:]...))
	}
}
