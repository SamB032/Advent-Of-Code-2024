package util

type Queue[T any] struct {
	Arr  []*T
	Size int
}

// Adds the item to the top of the array
func (queue *Queue[T]) Push(element *T) {
	queue.Arr = append(queue.Arr, element)
	queue.Size++
}

// Removes the item in the first
func (queue *Queue[T]) Pop() *T {
	if queue.Size <= 0 {
		panic("Queue Empty")
	}

	element := queue.Arr[0] // Get the first element
	queue.Arr = RemoveIndex(queue.Arr, 0)
	queue.Size--
	return element
}
