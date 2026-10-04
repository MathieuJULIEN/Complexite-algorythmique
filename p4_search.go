package main

func SearchV1(ligne []int, v int) int {
	for i, w := range ligne {
		if v == w {
			return  i
		}
	}
	return -1
}
