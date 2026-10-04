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

func BenchmarkSearch(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Sorted(n)
		m := Random(1, n)
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				sink = SearchV1(pile, m[0])
			}
		})
	}
}

var sinkSlice []int
func BenchmarkCount(b *testing.B) {
	for _, n := range []int{1_000, 10_000, 100_000} {
		pile := Sorted(n)
		m := Random(n, n)
		plafond := n
		if len(m)>0 {
			plafond = m[0]
		}
		b.Run(fmt.Sprintf("V1/n=%d", n), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sinkSlice = CountV1(pile, plafond)
			}
		})
	}
}