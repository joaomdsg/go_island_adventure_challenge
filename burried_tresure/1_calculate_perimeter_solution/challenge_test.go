package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCR(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestXxx(t *testing.T) {
	points := [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	CalculatePerimeter(points)
}
