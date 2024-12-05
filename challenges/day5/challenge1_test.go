package day5

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay5challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day5Challenge1(input)

	if got != 143 {
		t.Errorf("Day5challenge1 is incorrect. got: %d, want: %d", got, 143)
	}
}
