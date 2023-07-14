package main

import "fmt"

func FindTreasure(islandMap [][]int) ([][]int, error) {
	if islandMap == nil {
		return nil, fmt.Errorf("expected an island map, but none was provided")
	}
	path := [][]int{{0, 0}}
	return path, nil
}
