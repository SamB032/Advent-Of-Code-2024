package day3

import (
	"regexp"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

const RegexMulPattern = `mul\((\d+),(\d+)\)`

func ExtractAndSumMultiplications(input string) int {
  re := regexp.MustCompile(RegexMulPattern)
  matches := re.FindAllStringSubmatch(input, -1)

  sum := 0 

  for _, match := range matches {
    // Extract the two numbers and multiply them
    x := util.MustAtoI(string(match[1]))
    y := util.MustAtoI(string(match[2]))
    sum += x * y
  }
	return sum
}

func Day3Challenge1(input string) int {
  return ExtractAndSumMultiplications(input)
}
