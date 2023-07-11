# Island Adventure Challenge - 1_Calculate_Perimeter

## Challenge Description

Welcome, explorer! As part of your Island Adventure, the first challenge 'Calculate_Perimeter' is to determine the perimeter of the island. You have a GPS device that provides you a series of coordinates (pairs of integers), which form a closed polygonal chain representing the shoreline of the island. Your task is to calculate the total length of this perimeter.

The coordinates are given in an array of 2D points. Each point is represented by an array `[x, y]` where `x` and `y` are the integer coordinates of the point. The points are connected in the order they appear in the array, and the last point is connected to the first point.

Write a function `CalculatePerimeter(points [][]int) float64` to solve this task. The function should take an array of points (2D integer arrays), and return the perimeter as a floating point number.

## Examples

```go
points := [][]int{{0,0},{0,1},{1,1},{1,0}}
CalculatePerimeter(points)
```
Output: `4`

```go
points := [][]int{{-3,-2},{-1,4},{6,1},{3,10},{-4,9}}
CalculatePerimeter(points)
```
Output: `27.982756057296896`

## Constraints

- The array of points will have at least 3 elements and at most 10000 elements.
- The coordinates will be integers in the range -10000 to 10000.

# Tests

You should write tests to ensure your solution is correct. Here are some assertions you might include in your tests:

```go
func TestCalculatePerimeter(t *testing.T) {
    points := [][]int{{0,0},{0,1},{1,1},{1,0}}
    assert.Equal(t, 4.0, CalculatePerimeter(points), "The perimeter of a unit square should be 4")

    points = [][]int{{-3,-2},{-1,4},{6,1},{3,10},{-4,9}}
    assert.InDelta(t, 27.982756057296896, CalculatePerimeter(points), 0.0000001, "The perimeter calculation should be accurate to within a small delta")

    points = [][]int{{0,0},{10,0},{10,10},{0,10}}
    assert.Equal(t, 40.0, CalculatePerimeter(points), "The perimeter of a square with side length 10 should be 40")
}
```