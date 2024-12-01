package util

import (
	"fmt"
	"math"
)

//Filters an array keeping values where the function applied to the values is true
func ArrayFilter[I any](arr []I, f func(I) bool) []I {
	var filteredArr []I

	for _, values := range arr {
		if f(values) {
			filteredArr = append(filteredArr, values)
		}
	}
	return filteredArr
}

//Reduce the array to a single value
func ArrayReduce[I any](arr []I, f func(I, I) I, initialValue I) I {
	var result = initialValue

	for _, value := range arr {
		result = f(result, value)
	}

	return result
}

//Map values of an array to a new value and type
func ArrayMap[I any, O any](arr []I, f func(I) O) []O {
	var mappedArray []O = make([]O, len(arr))

	for i, value := range arr {
		mappedArray[i] = f(value)
	}

	return mappedArray
}

// Returns true if a array contains that item
func ArrayContains[I comparable](arr []I, searchTerm I) bool {
	for _, value := range arr {
		if value == searchTerm {
			return true
		}
	}
	return false
}

// Return the min and max value of a number array
func MinMax[I Number](arr []I) (I, I) {
	if len(arr) == 0 {
		panic(fmt.Errorf("empty array passed into MinMax Function"))
	}

	maxValue := arr[0]
	minValue := arr[0]

	for _, value := range arr {
		if maxValue < value {
			maxValue = value
		}
		if minValue > value {
			minValue = value
		}
	}

	return minValue, maxValue
}

// Return the index and value of the minimum in the array
func MinIndex(arr []int) (int, int) {
	var minIndex int = 0
	var minVal int = math.MaxInt

	for i, val := range arr {
		//Update the min value and index with new min
		if val < minVal {
			minIndex = i
			minVal = val
		}
	}
	return minIndex, minVal
}

//Removes an index from an array
func RemoveIndex[I any](slice []I, index int) []I {
	return append(slice[:index], slice[index+1:]...)
}
