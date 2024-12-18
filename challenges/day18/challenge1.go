package day18

import (
	"container/heap"
	"strings"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

const WALL = '#'
const EMPTY = '.'

var directions = []util.Point[int]{
	{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0},
}

type NeighbourWithCost struct {
	Point *util.Point[int]
	Cost  int
}

// Returns the max x and y coorinates
func CalculateGridSize(points []*util.Point[int]) (int, int) {
	Xmax := 0
	Ymax := 0

	for _, point := range points {
		if point.X > Xmax {
			Xmax = point.X
		}
		if point.Y > Ymax {
			Ymax = point.Y
		}
	}
	return Xmax + 1, Ymax + 1 // +1 to account for zero-based indexing
}

func CreateGrid(points []*util.Point[int], Xmax int, Ymax int) *[][]rune {
	// Create a Xmax by Ymax gird
	grid := make([][]rune, Ymax)
	for i := range grid {
		grid[i] = make([]rune, Xmax)
	}

	// Fill the empty space with '.'
	for i := 0; i < Ymax; i++ {
		for j := 0; j < Xmax; j++ {
			grid[i][j] = EMPTY
		}
	}

	// Place the points on the grid marked by 'x'
	for _, point := range points {
		grid[point.Y][point.X] = WALL
	}

	return &grid
}

func ParsePoints(input []string) []*util.Point[int] {
	var points []*util.Point[int]

	for _, row := range input {
		rowSplit := strings.Split(row, ",")
		points = append(points, &util.Point[int]{X: util.MustAtoI(rowSplit[0]), Y: util.MustAtoI(rowSplit[1])})
	}
	return points
}

func GetNeighbours(currentPoint *util.Point[int], currentCost int, grid *[][]rune) []*NeighbourWithCost {
	var neighbours []*NeighbourWithCost

	rows := len(*grid)
	cols := len((*grid)[0])

	// Iterate over all possible directions
	for _, dir := range directions {
		newX := currentPoint.X + dir.X
		newY := currentPoint.Y + dir.Y

		if newX >= 0 && newX < rows && newY >= 0 && newY < cols && (*grid)[newX][newY] == EMPTY {
			// Add the neighbour
			neighbours = append(neighbours, &NeighbourWithCost{
				Point: &util.Point[int]{X: newX, Y: newY},
				Cost:  currentCost + 1,
			})
		}
	}
	return neighbours
}

func DijkstrasAlgorithm(grid *[][]rune, Xmax int, Ymax int) int {
	startingPoint := &util.Point[int]{X: 0, Y: 0}
	endingPoint := &util.Point[int]{X: Xmax - 1, Y: Ymax - 1}

	visitedMap := make(map[util.Point[int]]bool)

	queue := util.NewPriorityQueue([]*util.Item[util.Point[int]]{{
		Value:    startingPoint,
		Priority: 0,
	},
	})

	// Initialise the Dijkstras Algorithm
	for queue.Len() > 0 {
		currentNode := heap.Pop(queue).(*util.Item[util.Point[int]])
		visitedMap[*currentNode.Value] = true

		if *currentNode.Value == *endingPoint {
			return currentNode.Priority
		}

		neighbours := GetNeighbours(currentNode.Value, currentNode.Priority, grid)

		// Add the neighbours in
		for _, neighbour := range neighbours {
			//If not in visited, add to queue
			if _, alreadyVisited := visitedMap[*neighbour.Point]; !alreadyVisited {
				heap.Push(queue, &util.Item[util.Point[int]]{
					Value:    neighbour.Point,
					Priority: neighbour.Cost,
				})
			}
		}
	}
	return -1
}

func Day18Challenge1(input []string) int {
	points := ParsePoints(input)
	Xmax, Ymax := CalculateGridSize(points)
	grid := CreateGrid(points, Xmax, Ymax)

	return DijkstrasAlgorithm(grid, Xmax, Ymax)
}
