package day18

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay18Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day18Challenge1(input)

	if got != 22 {
		t.Errorf("Day18Challenge1 is incorrect. Got: %d, Want: %d", got, 22)
	}
}
