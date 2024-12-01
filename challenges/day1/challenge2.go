package day1

func CalculateFrequency(leftArray []int, rightArray []int) map[int]int {
	frequencyMap := make(map[int]int)

	for _, leftVal := range leftArray {
		count := 0

		//Count the number of times leftVal appears in rightArray
		for _, rightVal := range rightArray {
			if leftVal == rightVal {
				count++
			}
		}
		if val, ok := frequencyMap[leftVal]; ok {
			//If key already exists, add to it
			frequencyMap[leftVal] = val + count

		} else {
			// Create new key and add the value
			frequencyMap[leftVal] = count
		}
	}
	return frequencyMap
}

func CalculateSimularity(freqMap map[int]int) int {
	sum := 0

	//Loop over the map and sum up key * val
	for key, val := range freqMap {
		sum += key * val
	}

	return sum
}

func Day1Challenge2(input []string) int {
	leftArray, rightArray := ParseIntoArray(input)

	freqMap := CalculateFrequency(leftArray, rightArray)
	simularity := CalculateSimularity(freqMap)

	return simularity
}
