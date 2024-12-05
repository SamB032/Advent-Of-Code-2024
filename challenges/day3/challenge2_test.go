package day3

import (
	"strings"
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay3challenge2(t *testing.T) {
	input := util.ReadFile("input.txt")
	inputJoined := strings.Join(input, "")

	got := Day3Challenge2(inputJoined)

	if got != 48 {
		t.Errorf("Day3challenge2 is incorrect. got: %d, want: %d", got, 48)
	}
}
