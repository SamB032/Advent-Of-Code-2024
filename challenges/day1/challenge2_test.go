package day1

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay1Challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day1Challenge2(input)

	if got != 31 {
		t.Errorf("Day1Challenge2 is incorret. Got: %d, Want: %d", got, 31)
	}
}
