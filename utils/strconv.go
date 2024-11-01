package util

import "strconv"

func MustAtoI(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return result
}
