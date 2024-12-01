package day1

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay1Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day1Challenge1(input)

	if got != 11 {
		t.Errorf("Day1Challenge1 is incorret. Got: %d, Want: %d", got, 11)
	}
}
