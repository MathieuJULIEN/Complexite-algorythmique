package main

func CountV1(pile []int, plafond int) []int {
	liste := make([]int, plafond+1)
	for _, a := range pile {
		if a <= plafond && a >= 0 {
			liste[a]++
		}
	}
	return liste
}
