package day4

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay4challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")

	runes := make([][]rune, len(input))
	for i, val := range input {
		runes[i] = []rune(val)
	}

	got := Day4Challenge2(runes)

	if got != 9 {
		t.Errorf("Day4challenge2 is incorrect. got: %d, want: %d", got, 9)
	}
}
