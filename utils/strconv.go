package util

import "strconv"

// Converts a string to int
func MustAtoI(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return result
}

// Converts a slice of strings to slice to integers
func StingsToInt(strs []string) []int {
	ints := make([]int, len(strs))

	for idx, val := range strs {
		ints[idx] = MustAtoI(val)
	}
	return ints
}
