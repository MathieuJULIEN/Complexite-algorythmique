package main



func DuplicateV1(pile []int) int {
	seen := make(map[int]struct{})

	for _, v := range pile {
		if _, b := seen[v]; b {
			return v
		} else {
			seen[v] = struct{}{}
		}
	}
	return -1
}
