package main


func SmallestV1(pile []int) int {
	if len(pile) <= 0 {
		return -1
	}
	result := pile[0]

	for _, v := range pile {
		if v < result {
			result = v
		}
	}

	return result
}