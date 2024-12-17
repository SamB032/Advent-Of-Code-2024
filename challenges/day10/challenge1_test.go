package day10

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay10Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	runes := make([][]rune, len(input))
	for i, val := range input {
		runes[i] = []rune(val)
	}

	got := Day10Challenge1(runes)

	if got != 36 {
		t.Errorf("Day10Challenge1 is incorrect. Got: %d, Want: %d", got, 36)
	}
}
