package day4

func DetectXMAS(input [][]rune, x, y, i, j int) int {
	count := 0

	// Left Down
	if i < x-3 && j > 2 &&
		(input[i][j] == 'X' && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' ||
			input[i][j] == 'S' && input[i+1][j-1] == 'A' && input[i+2][j-2] == 'M' && input[i+3][j-3] == 'X') {
		count++
	}

	// Down
	if i < x-3 &&
		(input[i][j] == 'X' && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' ||
			input[i][j] == 'S' && input[i+1][j] == 'A' && input[i+2][j] == 'M' && input[i+3][j] == 'X') {
		count++
	}

	// Down Right
	if i < x-3 && j < y-3 &&
		(input[i][j] == 'X' && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' ||
			input[i][j] == 'S' && input[i+1][j+1] == 'A' && input[i+2][j+2] == 'M' && input[i+3][j+3] == 'X') {
		count++
	}

	// Right
	if j < y-3 &&
		(input[i][j] == 'X' && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' ||
			input[i][j] == 'S' && input[i][j+1] == 'A' && input[i][j+2] == 'M' && input[i][j+3] == 'X') {
		count++
	}
	return count
}

func Day4Challenge1(input [][]rune) int {
	x := len(input)
	y := len(input[0])

	count := 0

	//Loop through each x,y and investigate if we reach a X or S
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			if input[i][j] == 'X' || input[i][j] == 'S' {
				count += DetectXMAS(input, x, y, i, j)
			}
		}
	}
	return count
}
