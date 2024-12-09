package day9

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay9Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")[0]

	got := Day9Challenge1(input)

	if got != 1928 {
		t.Errorf("Day9Challenge1 is incorrect. Got: %d, Want: %d", got, 1928)
	}
}
