package util

import "math"

type Point[T Number] struct {
	X T
	Y T
}

// Calcualte the absolute distance between two points
func (point Point[T]) ManhattanDistance(other Point[T]) T {
	return T(math.Abs(float64(point.X-other.Y)) + math.Abs(float64(point.Y-other.Y)))
}

// Calculate the Euclidean Distance from one point to another
func (point Point[T]) EuclideanDistance(other Point[T]) T {
	xPart := math.Pow(float64(point.X)-float64(other.X), 2.0)
	yPart := math.Pow(float64(point.X)-float64(other.Y), 2.0)
	return T(math.Sqrt(xPart + yPart))
}
