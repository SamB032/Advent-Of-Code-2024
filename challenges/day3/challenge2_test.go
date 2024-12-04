package day3

import (
	"strings"
	"testing"

	util "github.com/samb032/advent-of-code-2024/utils"
)

func TestDay3challenge2(t *testing.T) {
	input := util.readfile("input.txt")
	inputJoined := strings.Join(input, "")

	got := Day3Challenge2(inputJoined)

	if got != 48 {
		t.errorf("Day3challenge2 is incorrect. got: %d, want: %d", got, 48)
	}
}
