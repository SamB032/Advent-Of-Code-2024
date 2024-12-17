package day10

import (
	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

var directions = []util.Point[int]{
	{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0},
}

func FindZeros(input *[][]rune) []*util.Point[int] {
	rows := len(*input)
	cols := len((*input)[0])

	var zeroPos []*util.Point[int]

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if (*input)[i][j] == '0' {
				zeroPos = append(zeroPos, &util.Point[int]{X: i, Y: j})
			}
		}
	}
	return zeroPos
}

func GetNeighbours(currentPoint *util.Point[int], input *[][]rune) []*util.Point[int] {
	var neighbours []*util.Point[int]

	rows := len(*input)
	cols := len((*input)[0])

	// Iterate over all possible directions
	for _, dir := range directions {
		newX := currentPoint.X + dir.X
		newY := currentPoint.Y + dir.Y

		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && (*input)[currentPoint.X][currentPoint.Y]+1 == (*input)[newX][newY] {
			// Add the neighbour
			neighbours = append(neighbours, &util.Point[int]{X: newX, Y: newY})
		}
	}
	return neighbours
}

func DepthFirstSearch(input *[][]rune, zeroPos *util.Point[int]) int {
	visitedMap := make(map[util.Point[int]]bool)
	foundNines := make(map[util.Point[int]]bool)

	//Create a stack with starting Zeros
	stack := util.Stack[util.Point[int]]{
		Arr: []*util.Point[int]{zeroPos},
		Size: 1,
	}

	for stack.Size > 0 {
		currentNode := stack.Pop()
		visitedMap[*currentNode] = true

		if (*input)[currentNode.X][currentNode.Y] == '9' {
			foundNines[*currentNode] = true
			continue
		}

		neighbours := GetNeighbours(currentNode, input)

		// Add the neighbours in
		for _, neighbour := range neighbours {
			//If not in visited, add to queue
			if _, alreadyVisited := visitedMap[*neighbour]; !alreadyVisited {
				stack.Push(neighbour)
			}
		}
	}
	return len(foundNines)
}

func Day10Challenge1(input [][]rune) int {
	zeroPos := FindZeros(&input)

	totalScore := 0
	for _, zero := range zeroPos {
		totalScore += DepthFirstSearch(&input, zero)
	}
	return totalScore 
}
