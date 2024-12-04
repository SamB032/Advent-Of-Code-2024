package day4

func DetectXMASShape(input [][]rune, i, j int) int {
	//Detect Diagonal
	if ((input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') || (input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M')) &&
		((input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S') || (input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M')) {
		return 1
	}
	return 0
}

func Day4Challenge2(input [][]rune) int {
	x := len(input)
	y := len(input[0])

	count := 0

	//Loop through each x,y and investigate if we reach a X or S
	for i := 1; i < x-1; i++ {
		for j := 1; j < y-1; j++ {
			if input[i][j] == 'A' {
				count += DetectXMASShape(input, i, j)
			}
		}
	}
	return count
}
