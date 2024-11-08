package util

import (
	"strconv"
	"unicode"
)

// Converts a string to int
func MustAtoI(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return result
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func IsDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
