package day10

import (
	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

func DepthFirstSearchForRating(input *[][]rune, zeroPos *util.Point[int], visitedMap map[util.Point[int]]bool) int {
  uniqueTrails := 0

	//Create a stack with starting Zeros
	stack := util.Stack[util.Point[int]]{
		Arr: []*util.Point[int]{zeroPos},
		Size: 1,
	}

	for stack.Size > 0 {
		currentNode := stack.Pop()
		visitedMap[*currentNode] = true

		if (*input)[currentNode.X][currentNode.Y] == '9' {
      uniqueTrails++
			continue
		}

		neighbours := GetNeighbours(currentNode, input)

		// Add the neighbours in
		for _, neighbour := range neighbours {
      stack.Push(neighbour)
		}
	}
	return uniqueTrails
}

func Day10Challenge2(input [][]rune) int {
	zeroPos := FindZeros(&input)
  totalRating := 0

	// For each trailhead (height 0 position), calculate its rating
	for _, zero := range zeroPos {
		visitedMap := make(map[util.Point[int]]bool)
		rating := DepthFirstSearchForRating(&input, zero, visitedMap)
		totalRating += rating
	}
	return totalRating
}
