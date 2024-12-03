package day2

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay2Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day2Challenge1(input)

	if got != 2 {
		t.Errorf("Day2Challenge1 is incorrect. Got: %d, Want: %d", got, 2)
	}
}
