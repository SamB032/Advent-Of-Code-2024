package day3

import (
	"regexp"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

const RegexMulPatternExtra = `(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`

func ExtractAndSumMultiplicationsExtra(input string) int {
  re := regexp.MustCompile(RegexMulPatternExtra)
  matches := re.FindAllStringSubmatch(input, -1)

  sum := 0 
  proceedWith := true

  for _, match := range matches {
    switch match[0] {
    case "do()":
      proceedWith = true
    case "don't()":
      proceedWith = false
    default:
      if proceedWith {
        // Extract the two numbers and multiply them
        x := util.MustAtoI(string(match[2]))
        y := util.MustAtoI(string(match[3]))
        sum += x * y
      }
    }
  }
	return sum
}

func Day3Challenge2(input string) int {
  return ExtractAndSumMultiplicationsExtra(input)
}
