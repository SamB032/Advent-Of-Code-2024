package day2

import (
	util "github.com/SamB032/Advent-Of-Code-2024/utils"
)

//For some reason, removing an index does not work, so result to a function to manually copy it
func CopyArrayWithoutIndex(input []int, index int) []int {
  var output []int

  for i, v := range input {
    if i != index {
      output = append(output, v)
    }
  }
  return output
}

//Returns true if the array is increasing or decreasing (diff by 3)
func isIncreasingDecreasingTollerance(arr []int) bool {
  //Return true if it is already increasing or decreasing
  if isIncreasingDecreasing(arr) {
    return true
  }

  //Check safely by removing each level at a time 
  for i := 0; i < len(arr); i++ {
    modifiedArr := CopyArrayWithoutIndex(arr, i)

    if isIncreasingDecreasing(modifiedArr) {
      return true
    }
  }
  return false
}

func Day2Challenge2(input []string) int {
  parsedInputs := ParseInputs(input)

  //Array of true and falses, whether they are increasing or decreasing
  trendOfLevels := util.ArrayMap(parsedInputs, isIncreasingDecreasingTollerance)
  
  //Array only contain true occurences
  postitiveNegativeTrends := util.ArrayFilter(trendOfLevels, func (a bool) bool { return a == true } )

  return len(postitiveNegativeTrends)
}
