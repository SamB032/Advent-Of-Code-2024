package day16

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay16Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	runes := make([][]rune, len(input))
	for i, val := range input {
		runes[i] = []rune(val)
	}

	got := Day16Challenge1(runes)

	if got != 7036 {
		t.Errorf("Day4Challenge1 is incorrect. Got: %d, Want: %d", got, 7036)
	}
}
