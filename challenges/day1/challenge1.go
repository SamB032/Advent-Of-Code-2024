package day1

import (
	"sort"
	"math"
	"strings"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

//Split the file into a 2d array
func ParseIntoArray(content []string) ([]int, []int) {
	var leftArray []int
	var rightArray []int

	//Loop line by line, adding items to left and right arrays 
	for _, line := range content {
		parts := strings.Fields(line) // Split the line into fields by whitespace.
		if len(parts) != 2 {
			panic("Invalid input line: must contain exactly two numbers")
		}

		leftVal := util.MustAtoI(parts[0])
		leftArray = append(leftArray, leftVal)

		rightVal := util.MustAtoI(parts[1])
		rightArray = append(rightArray, rightVal)
	}
	return leftArray, rightArray
}

// Loops through and finds the difference between to minimum array and removes them
func CalculateTotalDistance(leftArray []int, rightArray []int) int {
	if len(leftArray) != len(rightArray) {
		panic("LeftArray and RigthArray are not the same size")
	}

	// Sort both arrays.
	sort.Ints(leftArray)
	sort.Ints(rightArray)

	// Calculate the total distance.
	totalDistance := 0
	for i := 0; i < len(leftArray); i++ {
		totalDistance += int(math.Abs(float64(leftArray[i] - rightArray[i])))
	}

	return totalDistance
}

func Day1Challenge1(input []string) int {
	leftArray, rightArray := ParseIntoArray(input)
	return CalculateTotalDistance(leftArray, rightArray)
}
