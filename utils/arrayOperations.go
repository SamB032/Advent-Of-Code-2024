package util

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
  var mappedArray []O

  for _, value := range arr {
    mappedArray = append(mappedArray, f(value))
  }
  
  return mappedArray
}
