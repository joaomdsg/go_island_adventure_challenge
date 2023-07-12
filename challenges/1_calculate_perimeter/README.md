# Challenge #1 - Calculate Perimeter 🗺️

Ahoy, explorer! Welcome to your first Island Adventure challenge: `Calculate Perimeter`. 🏝️

Your task, should you choose to accept it, is to find the total length of our island's perimeter. Don't worry, we're not sending you off with just a broken compass and a torn map. You'll have a trusty GPS device that'll give you a series of coordinates forming a closed polygonal chain (think of it as a funky line that doesn't cross itself and joins back at the start).

Now, ready to dive in? Here's what you need to do:

## Challenge Description 📝

You'll get an array of 2D points. Each point is represented as a sub-array `[x, y]`, where `x` and `y` are integer coordinates. The points connect in the order they appear in the array, and the first and last points join together to close the shape.

Your mission: write a function `CalculatePerimeter(points [][]int) float64` that takes in this array of points and returns the perimeter as a floating point number.

## Examples

You'll want to test out your function. Feel free to use these examples as a starting point:

```go
points := [][2]int{{0,0},{0,1},{1,1},{1,0}}
CalculatePerimeter(points)
```
Output: `4`

```go
points := [][2]int{{-3,-2},{-1,4},{6,1},{3,10},{-4,9}}
CalculatePerimeter(points)
```
Output: `48.57564046313958`

## Constraints

Every challenge has its limits. Here are yours:

- The array of points will have at least 3 elements and at most 10000 elements.
- The coordinates will be integers in the range -10000 to 10000.

# Acceptance Tests

Alright, it's time to make sure your solution holds water. When you fell confident of your implementation, copy these tests, one by one, to your test file. Save in between a check if the tests pass. And if all tests have passed, concratulations!! You have now been deemed worthy of the next challenge! 🏆

```go
func TestCalculatePerimeter_Square_Unit(t *testing.T) {
	points := [][2]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	p, _ := CalculatePerimeter(points)
	assert.Equal(t, 4.0, p, "The perimeter of a unit square should be 4")
}
```

```go
func TestCalculatePerimeter_Complex_Shape(t *testing.T) {
	points := [][2]int{{-3, -2}, {-1, 4}, {6, 1}, {3, 10}, {-4, 9}}
	p, _ := CalculatePerimeter(points)
	assert.InDelta(t, 48.57564046313958, p, 0.0000001, "The perimeter calculation should be accurate to within a small delta")
}
```

```go
func TestCalculatePerimeter_Too_Many_Points(t *testing.T) {
	points := make([][2]int, 10001)
	for i := range points {
		points[i] = [2]int{0, 0}
	}
	_, err := CalculatePerimeter(points)
	assert.Error(t, err, "The function should return an error if more than 10000 points are given")
}
```

And don't forget the constraints tests:

```go
func TestCalculatePerimeter_OutOfRange_Coordinates_UpperBound(t *testing.T) {
	points := [][2]int{{0, 0}, {0, 10001}, {10001, 10001}, {10001, 0}}
	_, err := CalculatePerimeter(points)
	assert.Error(t, err, "The function should return an error if any coordinates are outside the range -10000 to 10000")
}
```

```go
func TestCalculatePerimeter_OutOfRange_Coordinates_LowerBound(t *testing.T) {
	points := [][2]int{{0, 0}, {0, -10001}, {-10001, -10001}, {-10001, 0}}
	_, err := CalculatePerimeter(points)
	assert.Error(t, err, "The function should return an error if any coordinates are outside the range -10000 to 10000")
}
```

Alright, explorer, it's time to venture forth and conquer this challenge! Good luck, and may your calculations be ever accurate. 🌴😎