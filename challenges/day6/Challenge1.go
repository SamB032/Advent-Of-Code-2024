package day6

import "fmt"

var Directions = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

// Gets the x,y position of the the starting point
func GetStartingNode(input [][]rune) (int, int) {
	x := len(input)
	y := len(input[0])

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if input[i][j] == '^' {
				return i, j
			}
		}
	}
	return 0, 0
}

func Day6Challenge1(input [][]rune) int {
	guard_x, guard_y := GetStartingNode(input)

	x := len(input)
	y := len(input[0])

	// Keep track of what the guard has visited
	directionIndex := 0
  visited := make(map[string]bool)
  visited[fmt.Sprintf("%d,%d", guard_x, guard_y)] = true

	//Walk the path of the gaurd
	for {
		// Get the next position of the guard
		delta_x := Directions[directionIndex][0]
		delta_y := Directions[directionIndex][1]

		next_x := guard_x + delta_x
		next_y := guard_y + delta_y

		// Check for out of bounds
		if next_x < 0 || next_x >= x || next_y < 0 || next_y >= y {
			break
		}

		// change direction when obstacle encountered
		if input[next_x][next_y] == '#' {
			directionIndex = (directionIndex + 1) % 4
			continue
		}

		// update position and visited
		guard_x = next_x
		guard_y = next_y
    position := fmt.Sprintf("%d,%d", guard_x, guard_y)
    if !visited[position] {
			visited[position] = true
		}
	}
  return len(visited)
}
