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
Output: `4.0`

```go
points := [][2]int{{0,0},{0,10},{10,10},{10,0}}
CalculatePerimeter(points)
```
Output: `40.0`

## Constraints

Every challenge has its limits. Here are yours:

- The array of points will have at least 3 elements and at most 10000 elements.
- The coordinates will be integers in the range -10000 to 10000.


Alright, explorer, it's time to venture forth and conquer this challenge! Good luck, and may your calculations be ever accurate. 🌴😎