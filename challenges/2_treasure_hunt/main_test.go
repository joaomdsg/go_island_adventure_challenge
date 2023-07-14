package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTreasure_Simple_Path(t *testing.T) {
	islandMap := [][]int{
		{0, 1},
		{1, 2}}
	path, _ := FindTreasure(islandMap)
	assert.Equal(t, [][]int{{0, 0}}, path, "The function should return a correct path to the treasure")
}
