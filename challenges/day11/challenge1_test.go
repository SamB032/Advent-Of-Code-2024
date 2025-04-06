package day11

import (
	"os"
	"strconv"
	"strings"
	"testing"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func TestDay11Challenge1(t *testing.T) {
	input := util.ReadFile("input.txt")
	inputSplit := strings.Split(input[0], " ")

	inputsInt := make([]int, len(inputSplit))
	for idx, val := range inputSplit {
		i, err := strconv.Atoi(val)
		if err != nil {
			os.Exit(1)
		}
		inputsInt[idx] = i
	}

	got := Day16Challenge1(inputsInt, 25)

	if got != 55312 {
		t.Errorf("Day16Challenge1 is incorrect. Got: %d, Want: %d", got, 55312)
	}
}
