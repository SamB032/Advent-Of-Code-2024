package day5

import (
	"strings"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

type Ordering struct {
	First  int
	Second int
}

// Returns True if a string contains a "|"
func IsOdering(line string) bool {
	return strings.Contains(line, "|")
}

// Returns True if a line does not contain "|" and not empty string
func IsUpdate(line string) bool {
	return line != "" && !IsOdering(line)
}

// Parses a x|x to a int
func ParseOrdering(line string) *Ordering {
	lineSplit := strings.Split(line, "|")
	return &Ordering{util.MustAtoI(lineSplit[0]), util.MustAtoI(lineSplit[1])}
}

// Turns a string of an array to an int array
func ParseUpdate(line string) []int {
	lineSplit := strings.Split(line, ",")
	updateArr := make([]int, len(lineSplit))

	for i, val := range lineSplit {
		updateArr[i] = util.MustAtoI(val)
	}
	return updateArr
}

func SplitUpdatesOrdering(input string) ([]*Ordering, [][]int) {
	inputSplit := strings.Split(input, "\n")

	var orderings []*Ordering
	var updates [][]int

	for _, val := range inputSplit {
		if IsOdering(val) {
			orderings = append(orderings, ParseOrdering(val))
		} else if IsUpdate(val) {
			updates = append(updates, ParseUpdate(val))
		}
	}
	return orderings, updates
}

func UpdateIsInOrder(update []int, orderings []*Ordering) bool {
	// Create a map to track the index of each page in the update
	pageIndex := make(map[int]int)
	for i, page := range update {
		pageIndex[page] = i
	}

	// Check each ordering rule
	for _, rule := range orderings {
		firstIndex, firstExists := pageIndex[rule.First]
		secondIndex, secondExists := pageIndex[rule.Second]

		// If both pages exist in the update, validate their order
		if firstExists && secondExists && firstIndex >= secondIndex {
			return false // Rule violated
		}
	}

	return true
}

func GetMiddleOfArray(arr []int) int {
	return arr[int(len(arr)/2)]
}

func Day5Challenge1(input []string) int {
	inputJoined := strings.Join(input, "\n")

	orderings, updates := SplitUpdatesOrdering(inputJoined)

	var sum int

	var IsOrdered []bool = make([]bool, len(updates))
	for i, val := range updates {
		if UpdateIsInOrder(val, orderings) {
			sum += GetMiddleOfArray(val)
		}
		IsOrdered[i] = UpdateIsInOrder(val, orderings)
	}
	return sum
}
