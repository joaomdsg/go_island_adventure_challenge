package main

import "math"

func CalculatePerimeter(points [][2]int) float64 {
	p := 0.0
	for idx, point := range points {
		if idx == 0 {
			p += calculateDistance(point, points[len(points)-1])
			continue
		}
		p += calculateDistance(point, points[idx-1])
	}
	return p
}

// Calculates distance between 2 points using:
// d = √((x2 - x1)² + (y2 - y1)²)
func calculateDistance(point1, point2 [2]int) float64 {
	x1 := float64(point1[0])
	x2 := float64(point1[1])
	y1 := float64(point2[0])
	y2 := float64(point2[1])
	d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return d
}
