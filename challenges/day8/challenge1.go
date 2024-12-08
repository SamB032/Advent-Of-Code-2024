package day8

import (
	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

// Parse the input and extract the locations of the antennas
func ParseInput(input [][]rune) map[rune][]util.Point[int] {
	freqs := map[rune][]util.Point[int]{}

	for i, line := range input {
		for j, ch := range line {
			if ch == '.' {
				continue
			}
			freqs[ch] = append(freqs[ch], util.Point[int]{i + 1, j})
		}
	}
	return freqs
}

func Day8Challenge1(input [][]rune) int {
	freqs := ParseInput(input)

	antiNodes := map[util.Point[int]]struct{}{}

	for _, locs := range freqs {
		// Loop through each pair of antennas and calculate the antiNodes
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				delta := util.Point[int]{locs[b].X - locs[a].X, locs[b].Y - locs[a].Y}

				// Create corrdinates the possible anti nodes
				anti1 := util.Point[int]{locs[a].X - delta.X, locs[a].Y - delta.Y}
				anti2 := util.Point[int]{locs[b].X + delta.X, locs[b].Y + delta.Y}

				if anti1.X >= 0 && anti1.X < 50 && anti1.Y >= 0 && anti1.Y < 50 {
					antiNodes[anti1] = struct{}{}
				}

				if anti2.X >= 0 && anti2.X < 50 && anti2.Y >= 0 && anti2.Y < 50 {
					antiNodes[anti2] = struct{}{}
				}
			}
		}
	}
	return len(antiNodes)
}
