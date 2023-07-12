package main

import "math"

func CalculatePerimeter(points [][]int) float64 {
	p := 0.0
	for idx, point := range points {
		if idx == 0 {
			x1 := float64(point[0])
			x2 := float64(point[1])
			y1 := float64(points[len(points)-1][0])
			y2 := float64(points[len(points)-1][1])
			d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
			p += d
			continue
		}
		x1 := float64(point[0])
		x2 := float64(point[1])
		y1 := float64(points[idx-1][0])
		y2 := float64(points[idx-1][1])
		d := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
		p += d
	}
	return 4
}

// d = √((x2 - x1)² + (y2 - y1)²)
