package day11

import (
	"fmt"
	"os"
	"strconv"
)

var cache map[int][]int

func convertStringToIntArray(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("Error converting %s to int\n", s)
		os.Exit(1)
	}
	return val
}

func removeLeadingZeros(s string) int {
	start := 0
	for start < len(s) && s[start] == '0' {
		start++
	}
	if start == len(s) {
		return 0
	}
	trimmed := s[start:]
	return convertStringToIntArray(trimmed)
}

/*
- Engraved with the number 0 -> replaced by a stone engraved with the number 1
- Engraved with a number that has an even number of digits -> it is replaced by two stones.
	- left half of the digits are engraved on the new left stone
	- right half of the digits are engraved on the new right stone. (The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
- Else -> replaced with new stone multiplied with 2024
*/
func NewStones(stoneValue int) []int {
	stonesDigits := strconv.Itoa(stoneValue)

	if stoneValue == 0 {
		return []int{1}
	} else if len(stonesDigits) % 2 == 0 {
		midpoint := len(stonesDigits) / 2
		lhs := convertStringToIntArray(stonesDigits[:midpoint])
		rhs := removeLeadingZeros(stonesDigits[midpoint:])

		return []int{lhs, rhs}
	} else {
		return []int{stoneValue * 2024}
	}
}

func Day16Challenge1(input []int, numbeOfBlinks int) int {
	cache = make(map[int][]int)

	// Loop over every blink
	for blink := 1; blink <= numbeOfBlinks; blink++ {
		// Loop over every element in stones
		i := 0
		for i < len(input) {
			newStones, ok := cache[input[i]]
			if !ok {
				// No hit in cache, calculate value and store in cache
				newStones = NewStones(input[i])
				cache[input[i]] = newStones
			}

			var nextInput []int
			nextInput = append(nextInput, input[:i]...) // Add elements up to the current position
			nextInput = append(nextInput, newStones...) // Add the new stones at position i
			nextInput = append(nextInput, input[i+1:]...) // Add remaining elements after the current position
			input = nextInput

			// Move to after insertion
			i += len(newStones)
		}
	}
	return len(input)
}
