package util

type Stack[T any] struct {
	Arr  []*T
	Size int
}

// Adds the item to the top of the array
func (stack *Stack[T]) Push(element *T) {
	stack.Arr = append(stack.Arr, element)
	stack.Size++
}

// Removes the item in the first
func (stack *Stack[T]) Pop() *T {
	if stack.Size <= 0 {
		panic("stack Empty")
	}
	element := stack.Arr[stack.Size-1] // Get the last element
	stack.Arr = RemoveIndex(stack.Arr, stack.Size-1)
	stack.Size--
	return element
}
