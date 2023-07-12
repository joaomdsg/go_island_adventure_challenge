package main

import (
	"fmt"
	"math"
)

func CalculatePerimeter(points [][2]int) (float64, error) {

	if len(points) < 3 || len(points) > 1000 {
		return 0, fmt.Errorf("invalid number of points provided: got %v, want at least 3 to form a polygon, and at most 10000", len(points))
	}

	for _, point := range points {
		x := float64(point[0])
		y := float64(point[1])
		if math.Abs(x) > 10000 || math.Abs(y) > 10000 {
			return 0, fmt.Errorf("contains points outside the range -10000 to 10000")
		}
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

// Calculates the Euclidean Distance between 2 points, so that
// d = √((x2 - x1)² + (y2 - y1)²)
func calculateEuclideanDistance(point1, point2 [2]int) float64 {
	x1 := float64(point1[0])
	x2 := float64(point1[1])
	y1 := float64(point2[0])
	y2 := float64(point2[1])
	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return d
}
