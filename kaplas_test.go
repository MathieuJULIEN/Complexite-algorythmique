package main

import (
"fmt"
"testing"
)

var sink int 

func BenchmarkSmallest(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = SmallestV1(pile)
			}
		})
	}
}

func BenchmarkDuplicate(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Shuffled(n) // préparation hors de la mesure
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = DuplicateV1(pile)
			}
		})
	}
}

func BenchmarkTowerHeight(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = TowerHeightV1(n)
			}
		})
	}
}