package util

// Counts the occurrences of each element in a slice.
func CountOccurrences[T comparable](slice []T) map[T]int {
    counts := make(map[T]int)
    for _, val := range slice {
        counts[val]++
    }
    return counts
}
