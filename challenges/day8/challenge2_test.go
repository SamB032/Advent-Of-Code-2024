package day8

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay8Challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")

	runes := make([][]rune, len(input))
	for i, val := range input {
		runes[i] = []rune(val)
	}
	got := Day8Challenge2(runes)

	if got != 126 {
		t.Errorf("Day8Challenge2 is incorrect. Got: %d, Want: %d", got, 126)
	}
}
