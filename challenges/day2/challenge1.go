package day2

import (
	"math"
	"strconv"
	"strings"

	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

//Returns true if the array is increasing or decreasing (diff by 3)
func isIncreasingDecreasing(arr []int) bool {
  if len(arr) <= 2 {
    return false
  }

  isIncreasing := true
	isDecreasing := true

  //Loop through the array determine the trend
	for i := 0; i < len(arr)-1; i++ {
    diff := math.Abs(float64(arr[i+1] - arr[i]))
    if diff < 1 || diff > 3 {
      return false
    } else if arr[i] > arr[i+1] {
			isIncreasing = false
		} else if arr[i] <= arr[i+1] {
			isDecreasing = false
		}
	}
	return isIncreasing || isDecreasing
}

// Turns the levels to a 2d int array
func ParseInputs(input []string) [][]int {
	var arr [][]int

  //Loop through each line
	for _, line := range input {
		var nestedArr []int
    splitLine := strings.Split(line, " ")

    //Loop through each number, convert it to int and add to arr
		for _, level := range splitLine {
      //Skip string if not a number
      number, err := strconv.Atoi(level)
      if err != nil {
        continue
      }
      nestedArr = append(nestedArr, number) 
		}
    arr = append(arr, nestedArr)
	}
  return arr
}

func Day2Challenge1(input []string) int {
  parsedInputs := ParseInputs(input)

  //Array of true and falses, whether they are increasing or decreasing
  trendOfLevels := util.ArrayMap(parsedInputs, isIncreasingDecreasing)
  
  //Array only contain true occurences
  postitiveNegativeTrends := util.ArrayFilter(trendOfLevels, func (a bool) bool { return a == true } )

  return len(postitiveNegativeTrends)
}
