package day2

import (
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay2Challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")

	got := Day2Challenge2(input)

	if got != 4 {
		t.Errorf("Day2Challenge2 is incorrect. Got: %d, Want: %d", got, 4)
	}
}
