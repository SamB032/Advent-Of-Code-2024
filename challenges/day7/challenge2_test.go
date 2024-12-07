package day7

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay7Challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day7Challenge2(input)

	if got != 11387 {
		t.Errorf("Day7Challenge2 is incorrect. Got: %d, Want: %d", got, 11387)
	}
}
