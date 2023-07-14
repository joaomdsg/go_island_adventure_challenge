package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// With this map and legend in hand, it's time to put your treasure hunting skills to the ultimate test.
// Tease out your solution incrementally, write a little test first, and make it pass, then repeat. Remember,
// one small step for man, one giant leap for mankind... if you're too ambitious, the island's magic will ~
// erase your progress up to the last successful state.
//
// Don't forget, each rule and condition is a vital part of the map to your treasure! Make sure to test them all thoroughly.
//
// Here is a test to start you out:

func TestFindTreasure_Simple_Path(t *testing.T) {
	islandMap := [][]int{
		{0, 1},
		{1, 2}}
	path, _ := FindTreasure(islandMap)
	assert.Equal(t, [][]int{{0, 0}}, path, "The function should return a correct path to the treasure")
}
