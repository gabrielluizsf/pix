package pix

import "sort"

func sortKeys(data intMap) []int {
	keys := make([]int, len(data))
	i := 0

	for k := range data {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	return keys
}
