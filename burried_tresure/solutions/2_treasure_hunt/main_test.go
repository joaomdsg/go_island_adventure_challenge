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

func Test_Treasure_At_1East(t *testing.T) {
	islandMap := [][]int{
		{1, 0},
		{2, 1},
	}
	path, _ := FindTreasure(islandMap)
	assert.Equal(t, [][]int{{0, 0}, {0, 1}}, path)

}

func Test_Treasure_At_1East_1South(t *testing.T) {
	islandMap := [][]int{
		{2, 1},
		{1, 0},
	}
	path, _ := FindTreasure(islandMap)
	assert.Equal(t, [][]int{{0, 0}, {0, 1}, {1, 1}}, path)

}

func Test_Treasure_At_1South(t *testing.T) {
	islandMap := [][]int{
		{1, 2},
		{0, 1},
	}
	path, _ := FindTreasure(islandMap)
	assert.Equal(t, [][]int{{0, 0}, {1, 0}}, path)
}

func Test_No_Island_Map(t *testing.T) {
	_, err := FindTreasure(nil)
	assert.Error(t, err)
}

func Test_Non_Square_Map(t *testing.T) {
	islandMap := [][]int{
		{1, 0},
	}
	_, err := FindTreasure(islandMap)
	assert.Error(t, err)
}

func Test_Map_Is_Too_Large(t *testing.T) {
	islandMapRow := make([]int, 101)
	islandMap := make([][]int, 0)
	for i := 0; i < 101; i++ {
		islandMap = append(islandMap, islandMapRow)
	}
	_, err := FindTreasure(islandMap)
	assert.Error(t, err)
}

func Test_Map_Is_Too_Small(t *testing.T) {
	islandMap := [][]int{ {0} }
	_, err := FindTreasure(islandMap)
	assert.Error(t, err)
}
