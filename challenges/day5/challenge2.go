package day5

import (
	"strings"
)

func ReorderUpdate(update []int, orderings []*Ordering) []int {
	swapped := true

	// Perform the bubble sort algorithm until the ordering rules are met
	for swapped {
		swapped = false
		for i := 0; i < len(update)-1; i++ {
			x, y := update[i], update[i+1]

			// Check if x|y exists and x must come before y
			for _, rule := range orderings {
				if rule.First == y && rule.Second == x {
					// Swap x and y
					update[i], update[i+1] = update[i+1], update[i]
					swapped = true
				}
			}
		}
	}
	return update
}

func Day5Challenge2(input []string) int {
	inputJoined := strings.Join(input, "\n")

	orderings, updates := SplitUpdatesOrdering(inputJoined)
	sum := 0

	for _, update := range updates {
		if !UpdateIsInOrder(update, orderings) {
			// Reorder the update
			corrected := ReorderUpdate(update, orderings)

			// Find the middle page of the corrected update
			middlePage := GetMiddleOfArray(corrected)
			sum += middlePage
		}
	}
	return sum
}
