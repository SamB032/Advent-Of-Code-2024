package day5

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay5challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day5Challenge2(input)

	if got != 105 {
		t.Errorf("Day5challenge2 is incorrect. got: %d, want: %d", got, 105)
	}
}
