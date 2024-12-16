package day16

import (
	"container/heap"
	"fmt"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

const wall = '#'
const rotationScore = 1000

var directions = []util.Point[int]{
	{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0},
}

type NeighbourWithCost struct {
	Point *util.Point[int]
	Cost  int
}

func FindChar(input *[][]rune, search rune) (*util.Point[int], error) {
	for x, rows := range *input {
		for y, _ := range rows {
			if (*input)[x][y] == search {
				return &util.Point[int]{X: x, Y: y}, nil
			}
		}
	}
	return nil, fmt.Errorf("Cannot find search term: %c", search)
}

func GetNeighbours(currentNode *util.Point[int], direction int, currentCost int, input *[][]rune) []*NeighbourWithCost {
	var neighbours []*NeighbourWithCost

	rows := len(*input)
	cols := len((*input)[0])

	// Iterate over all possible directions
	for i, dir := range directions {
		newX := currentNode.X + dir.X
		newY := currentNode.Y + dir.Y

		// Check if the new position is within bounds and not a wall
		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && (*input)[newX][newY] != wall {
			if i != direction { // If the direction has changed (rotation), add a rotation cost
				currentCost += rotationScore
			}

			// Add the neighbour and its cost to the list
			neighbours = append(neighbours, &NeighbourWithCost{
				Point: &util.Point[int]{X: newX, Y: newY},
				Cost:  currentCost + 1,
			})
		}
	}
	return neighbours
}

func ReconstructPath(startingPoint *util.Point[int], endingPoint *util.Point[int],
	visitedMap map[util.Point[int]]*util.Point[int]) []util.Point[int] {
	path := []util.Point[int]{}
	current := endingPoint

	for current != nil {
		path = append(path, *current)
		next, exists := visitedMap[*current] // Get the previous node

		if !exists || next == nil {
			break // If the previous node does not exist, break out of the loop
		}
		current = next
	}
	Reverse(path)
	return path
}

// Helper function to reverse a slice
func Reverse(path []util.Point[int]) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}

func FindShortestPath(input *[][]rune) []util.Point[int] {
	startingPoint, err1 := FindChar(input, rune('S'))
	if err1 != nil {
		panic(err1)
	}

	endingPoint, err2 := FindChar(input, rune('E'))
	if err2 != nil {
		panic(err2)
	}

	visitedMap := make(map[util.Point[int]]*util.Point[int])

	queue := util.NewPriorityQueue([]*util.Item[util.Point[int]]{{
		Value:    startingPoint,
		Priority: 0,
	},
	})

	visitedMap[*startingPoint] = nil

	direction := 1 // Facing East

	// Initialise the BFS Loop
	for queue.Len() > 0 {
		currentNode := heap.Pop(queue).(*util.Item[util.Point[int]])

		if *currentNode.Value == *endingPoint {
			// Reconstruct the path and return it
			fmt.Println(currentNode.Priority)
			return ReconstructPath(startingPoint, endingPoint, visitedMap)
		}

		neighbours := GetNeighbours(currentNode.Value, direction, currentNode.Priority, input)

		// Add the neighbours in
		for _, neighbour := range neighbours {
			//If not in visited, add to queue
			if _, alreadyVisited := visitedMap[*neighbour.Point]; !alreadyVisited {
				heap.Push(queue, &util.Item[util.Point[int]]{
					Value:    neighbour.Point,
					Priority: neighbour.Cost,
				})
				visitedMap[*neighbour.Point] = currentNode.Value
			}
		}
	}
	return nil
}

func DeterminePoints(path []util.Point[int]) int {
	baseScore := len(path) - 1
	turnCount := 0

	//Determine the number of 90 degree turns
	for i := 1; i < len(path)-1; i++ {
		prev := path[i-1] // Point i-1
		curr := path[i]   // Point i
		next := path[i+1] // Point i+1

		// Calculate differences to determine direction
		dx1 := curr.X - prev.X
		dy1 := curr.Y - prev.Y
		dx2 := next.X - curr.X
		dy2 := next.Y - curr.Y

		// 90-degree turn occurs if the direction changes between axes
		if dx1*dy2 != dx2*dy1 {
			turnCount += rotationScore
		}
	}
	return baseScore + turnCount + 1000 // For direction at start
}

func Day16Challenge1(input [][]rune) int {
	path := FindShortestPath(&input)
	return DeterminePoints(path)
}
