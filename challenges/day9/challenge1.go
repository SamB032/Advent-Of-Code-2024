package day9

const emptyBlock = '.'

// Generate a rune array of a specific length
func GenerateRuneArray(ch rune, length int) []rune {
	var arr []rune = make([]rune, length)
	for i := range arr {
		arr[i] = ch
	}
	return arr
}

// Generates the input in the long format with "12....23..232"
func ParseInput(input string) []rune {
	//isBlock should alternate between block and free space
	var fileId []rune
	currentNumber := '0'

	for i, val := range input {
		isBlock := i%2 == 0     // Modulo operation to iterator between block and empty
		intVal := int(val) - 48 // -48 as we go from ascii to int

		if isBlock {
			fileId = append(fileId, GenerateRuneArray(currentNumber, intVal)...)
			currentNumber++
		} else {
			fileId = append(fileId, GenerateRuneArray(emptyBlock, intVal)...)
		}
	}
	return fileId
}

// Gets first empty character
func GetFirstEmptyIndex(input []rune) int {
	for i, val := range input {
		if val == emptyBlock {
			return i
		}
	}
	return -1
}

// Get the last non empty character in a string
func GetLastCharacterIndex(input []rune) int {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != emptyBlock {
			return i // Return the index of the last non-empty character
		}
	}
	return -1
}

// Compacts the fileId by removing the empty space
func CompactId(input []rune) []rune {
	counter := 0
	tempInput := make([]rune, len(input))
	copy(tempInput, input)

	for {
		// Get first empty index, if we can't get one, exit
		firstEmptyIndex := GetFirstEmptyIndex(tempInput)
		if firstEmptyIndex == -1 {
			break
		}

		lastCharacterIndex := GetLastCharacterIndex(tempInput)

		if lastCharacterIndex < firstEmptyIndex {
			break
		}

		// Make the swap, loss is ok
		tempInput[firstEmptyIndex] = tempInput[lastCharacterIndex]
		tempInput[lastCharacterIndex] = emptyBlock
		counter++

		// If counter is a Modulo of 5, save the tempInput into input
		if counter%4 == 0 {
			input = make([]rune, len(tempInput))
			copy(input, tempInput)
		}
	}
	return input
}

func SumOverCompactId(input []rune) int {
	sum := 0

	for i, val := range input {
		if val == emptyBlock {
			continue
		}
		intVal := int(val) - 48 // -48 as we go from ascii to int (Ascii a = 48)
		sum += i * intVal
	}
	return sum
}

func Day9Challenge1(input string) int {
	parsedInput := ParseInput(input)
	compactId := CompactId(parsedInput)
	return SumOverCompactId(compactId)
}
