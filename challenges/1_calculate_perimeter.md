# 1_Calculate_Perimeter 🗺️

Ahoy, explorer! Welcome to your first Island Adventure challenge: `calculate_perimeter`. 🏝️

Your task, should you choose to accept it, is to find the total length of our island's perimeter. Don't worry, we're not sending you off with just a broken compass and a torn map. You'll have a trusty GPS device that'll give you a series of coordinates forming a closed polygonal chain (think of it as a funky line that doesn't cross itself and joins back at the start).

Now, ready to dive in? Here's what you need to do:

## Challenge Description 📝

You'll get an array of 2D points. Each point is represented as a sub-array `[x, y]`, where `x` and `y` are integer coordinates. The points connect in the order they appear in the array, and the first and last points join together to close the shape.

Your mission: write a function `CalculatePerimeter(points [][]int) float64` that takes in this array of points and returns the perimeter as a floating point number.

## Examples

You'll want to test out your function. Feel free to use these examples as a starting point:

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

Every challenge has its limits. Here are yours:

- The array of points will have at least 3 elements and at most 10000 elements.
- The coordinates will be integers in the range -10000 to 10000.

# Tests

Alright, it's time to make sure your solution holds water. Split up your tests into smaller ones, each focusing on a specific scenario. You can use these to get started:

```go
func TestCalculatePerimeter_Square_Unit(t *testing.T) {
    points := [][]int{{0,0},{0,1},{1,1},{1,0}}
    assert.Equal(t, 4.0, CalculatePerimeter(points), "The perimeter of a unit square should be 4")
}
```

```go
func TestCalculatePerimeter_Complex_Shape(t *testing.T) {
    points = [][]int{{-3,-2},{-1,4},{6,1},{3,10},{-4,9}}
    assert.InDelta(t, 27.982756057296896, CalculatePerimeter(points), 0.0000001, "The perimeter calculation should be accurate to within a small delta")
}
```

```go
func TestCalculatePerimeter_Square_Large(t *testing.T) {
    points = [][]int{{0,0},{10,0},{10,10},{0,10}}
    assert.Equal(t, 40.0, CalculatePerimeter(points), "The perimeter of a square with side length 10 should be 40")
}
```

And don't forget to test when the constraints are not met:

```go
func TestCalculatePerimeter_Fewer_Points(t *testing.T) {
    points := [][]int{{0,0},{0,1}}
    assert.Error(t, CalculatePerimeter(points), "The function should return an error if fewer than 3 points are given")
}
```

```go
func TestCalculatePerimeter_Many_Points(t *testing.T) {
    points := make([][]int, 10001)
    for i := range points {
        points[i] = []int{i, i}
    }
    assert.Error(t, CalculatePerimeter(points), "The function should return an error if more than 10000 points are given")
}
```

```go
func TestCalculatePerimeter_OutOfRange_Coordinates(t *testing.T) {
    points := [][]int{{0,0},{0,10001},{10001,10001},{10001,0}}
    assert.Error(t, CalculatePerimeter(points), "The function should return an error if any coordinates are outside the range -10000 to 10000")
}
```

Alright, explorer, it's time to venture forth and conquer this challenge! Good luck, and may your calculations be ever accurate. 🌴😎