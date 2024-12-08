package util

type Queue struct {
	Arr  []int
	Size int
}

// Adds the item to the top of the array
func (queue *Queue) Push(element int) {
	queue.Arr = append(queue.Arr, element)
	queue.Size++
}

// Removes the item in the first
func (queue *Queue) Pop() int {
	if queue.Size <= 0 {
		return 0
	}

	element := queue.Arr[0] // Get the first element
	queue.Arr = RemoveIndex(queue.Arr, 0)
	return element
}
