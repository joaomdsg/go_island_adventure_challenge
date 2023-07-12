package main

import (
	"fmt"
	"math"
)

func CalculatePerimeter(points [][2]int) (float64, error) {
	if len(points) < 3 || len(points) > 1000 {
		return 0, fmt.Errorf("invalid number of points provided: got %v, want at least 3 to form a polygon, and at most 10000", len(points))
	}
	p := 0.0
	for idx, point := range points {
		if idx == 0 {
			p += calculateEuclideanDistance(point, points[len(points)-1])
			continue
		}
		p += calculateEuclideanDistance(point, points[idx-1])
	}
	return p, nil
}

// Calculates the Euclidean Distance between 2 points.
// d = √((x2 - x1)² + (y2 - y1)²)
func calculateEuclideanDistance(point1, point2 [2]int) float64 {
	x1 := float64(point1[0])
	x2 := float64(point1[1])
	y1 := float64(point2[0])
	y2 := float64(point2[1])
	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return d
}
