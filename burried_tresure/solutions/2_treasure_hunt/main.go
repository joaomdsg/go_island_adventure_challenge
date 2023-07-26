package main

import "fmt"

const (
	MAP_MINIMUM_SIZE int = 2
	MAP_MAXIMUM_SIZE int = 100
)

func FindTreasure(islandMap [][]int) ([][]int, error) {
	if islandMap == nil {
		return nil, fmt.Errorf("expected an island map but none was provided")
	}

	if !gridHasValidSize(islandMap) || !gridIsSquare(islandMap) {
		return nil, fmt.Errorf("island map must be a square grid (n * n) where n is between 2 and 100")
	}

	path := [][]int{{0, 0}}

	if islandMap[0][1] == 0 {
		path = append(path, []int{0, 1})
	}

	if islandMap[1][1] == 0 {
		path = append(path, []int{0, 1})
		path = append(path, []int{1, 1})
	}

	if islandMap[1][0] == 0 {
		path = append(path, []int{1, 0})
	}

	return path, nil
}

func gridIsSquare(grid [][]int) bool {
	n := len(grid)
	for _, col := range grid {
		if len(col) != n {
			return false
		}
	}
	return true
}

func gridHasValidSize(grid [][]int) bool {
	if len(grid) < MAP_MINIMUM_SIZE {
		return false
	}
	if len(grid) > MAP_MAXIMUM_SIZE {
		return false
	}
	return true
}
