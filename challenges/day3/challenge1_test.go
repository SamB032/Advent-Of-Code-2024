package day3

import (
	"strings"
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay3Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")
	inputJoined := strings.Join(input, "")

	got := Day3Challenge1(inputJoined)

	if got != 161 {
		t.Errorf("Day3Challenge1 is incorrect. Got: %d, Want: %d", got, 161)
	}
}
