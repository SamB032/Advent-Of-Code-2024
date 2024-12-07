package day7

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay7Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day7Challenge1(input)

	if got != 3749 {
		t.Errorf("Day7Challenge1 is incorrect. Got: %d, Want: %d", got, 3749)
	}
}
