package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Treasure_At_Starting_Cell(t *testing.T) {
	islandMap := [][]int{
		{0, 1},
		{1, 2},
	}
	path, _ := FindTreasure(islandMap)
	assert.Equal(t, [][]int{{0, 0}}, path)
}

func Test_No_Island_Map(t *testing.T) {
	_, err := FindTreasure(nil)
	assert.Error(t, err)
}
