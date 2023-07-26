# Challenge #2 - Island Treasure Hunt 🏖️🔍

Greetings, fearless adventurer! 🏴‍☠️

Your next task is to embark on a treasure hunt. This island is famous for its magical puzzles and encrypted clues. Only the most intelligent, agile, and diligent adventurers (like you!) can hope to find the treasure. 💎

You'll be given a two-dimensional map grid of the island, represented as a 2D array. Each array cell contains an integer between 0 and 9, representing the straight-line distance from that cell to the treasure.

Here's how it works:

- Each integer clue corresponds to the number of steps in any direction to the treasure. For instance, if a cell contains the number 5, the treasure is 5 steps away from that cell, in either north, south, east, or west direction.
- You can move in any of the four cardinal directions: north, east, south, or west (but not diagonally).
- The hunt starts at cell [0, 0] and ends when you find a cell with the number 0, which indicates the treasure's location.

Your mission: write a function FindTreasure(map [][]int) ([][]int, error) that takes in this 2D map array and returns the path to the treasure as a 2D array of the cell indices.

Before you begin, there are several constraints and conditions you need to be aware of:

1. The map must be a square grid (n * n), where n is between 1 and 100. If it's not, your function should return an error.
2. The map cannot be empty. If it is, your function should return an error.
3. The numbers on the cells must accurately represent the straight-line distance to the treasure. If they don't, your function should return an error.
4. There must be exactly one treasure on the map. If there is more than one, your function should return an error.
5. The algorithm should always move to the adjecent cell with the lowest number. If there are multiple adjecent cells with that same lowest number, the movement should be in this order of preference: East, South, West, North.

## Example 📝

Consider this example map:

```go
[[3, 2, 1],
 [2, 1, 0],
 [1, 2, 3]]
```

The path to the treasure from [0, 0] would be:

```go
[[0, 0], [0, 1], [0, 2], [1, 2]]
```

Let's break it down:

You start at [0, 0] which contains the number 3. This indicates that the treasure is 3 cells away.

You check the number of each adjacent cell to find out that there are 2 adjecent cells with the lowest number of 2.

Seeing thrat there is more than one possible move, you should follow the order of preference: 1. East, 2. South, 3. West, 4. North.

You move one cell east to [0, 1] which contains the number 2. This tells you the treasure is now 2 cells away.

Again, you check the number of each adjecent cell and repeat the same logic to then move east to [0, 2] where the number is 1, indicating that the treasure is now just 1 cell away.

Finally, you move south to the cell with the lower adjecent number, [1, 2] which contains the number 0 - the treasure!

Alright, explorer, it's time to embark on your treasure hunt! Good luck, and may your path lead you straight to the treasure! 🌴🏴‍☠️🔍