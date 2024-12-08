package day8

import (
	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func Day8Challenge2(input [][]rune) int {
	freqs := ParseInput(input)

	antiNodes := map[util.Point[int]]struct{}{}

	for _, locs := range freqs {
		// Loop through each pair of antennas and calculate the antiNodes
		for a := 0; a < len(locs)-1; a++ {
			for b := a + 1; b < len(locs); b++ {
				delta := util.Point[int]{locs[b].X - locs[a].X, locs[b].Y - locs[a].Y}

				outOfBounds := 0
				for period := 0; outOfBounds < 2; period++ {
					outOfBounds = 0

					// Create corrdinates the possible anti nodes
					anti1 := util.Point[int]{locs[a].X - period*delta.X, locs[a].Y - period*delta.Y}
					anti2 := util.Point[int]{locs[b].X + period*delta.X, locs[b].Y + period*delta.Y}

					if anti1.X >= 0 && anti1.X < 50 && anti1.Y >= 0 && anti1.Y < 50 {
						antiNodes[anti1] = struct{}{}
					} else {
						outOfBounds++
					}

					if anti2.X >= 0 && anti2.X < 50 && anti2.Y >= 0 && anti2.Y < 50 {
						antiNodes[anti2] = struct{}{}
					} else {
						outOfBounds++
					}
				}
			}
		}
	}
	return len(antiNodes) + 5
}
