package main

func TowerHeightV1(n int) int {
	result := 0
	if n < 1 {
		return -1
	}
	for i := n; i > 0; i-- {
		result += i
	}
	return result
}
