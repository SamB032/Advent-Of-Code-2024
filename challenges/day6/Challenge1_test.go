package day6

import (
  "testing"

  util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay6Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	runes := make([][]rune, len(input))
	for i, val := range input {
		runes[i] = []rune(val)
	}

	got := Day6Challenge1(runes)

	if got != 41 {
		t.Errorf("Day6Challenge1 is incorrect. Got: %d, Want: %d", got, 41)
	}
}
