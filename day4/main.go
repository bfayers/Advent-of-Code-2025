package main

import (
	"fmt"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

type Pos struct {
	x    int
	y    int
	roll bool
}

type Grid struct {
	data [][]Pos
}

func check_accessible(grid Grid, x int, y int) bool {
	// Check adjacent spaces
	adjacent := 0

	// Check up/down/left/right
	if x > 0 && grid.data[y][x-1].roll {
		adjacent++
	}
	if x < len(grid.data[y])-1 && grid.data[y][x+1].roll {
		adjacent++
	}
	if y > 0 && grid.data[y-1][x].roll {
		adjacent++
	}
	if y < len(grid.data)-1 && grid.data[y+1][x].roll {
		adjacent++
	}

	// Check diagonal spaces
	if x > 0 && y > 0 && grid.data[y-1][x-1].roll {
		adjacent++
	}
	if x < len(grid.data[y])-1 && y > 0 && grid.data[y-1][x+1].roll {
		adjacent++
	}
	if x > 0 && y < len(grid.data)-1 && grid.data[y+1][x-1].roll {
		adjacent++
	}
	if x < len(grid.data[y])-1 && y < len(grid.data)-1 && grid.data[y+1][x+1].roll {
		adjacent++
	}
	if adjacent < 4 {
		return true
	}
	return false
}

func part1(grid Grid) int {
	// Find all rolls in grid that have less than 4 of the 8 adjacent spaces filled
	// Treat grid borders as unfilled
	var total_accessible int
	for y, line := range grid.data {
		for x, pos := range line {
			if pos.roll {
				if check_accessible(grid, x, y) {
					total_accessible++
				}
			}
		}
	}
	return total_accessible
}

func part2(grid Grid) int {
	// In part 2 we can remove rolls that are accessible in the first instance to make more available until we can no longer remove any rolls
	var total_accessible int
	for part1(grid) > 0 {
		for y, line := range grid.data {
			for x, pos := range line {
				if pos.roll {
					if check_accessible(grid, x, y) {
						grid.data[y][x].roll = false
						total_accessible++
					}
				}
			}
		}
	}
	return total_accessible
}

func main() {
	// Load the data
	lines := utils.GetFileLines("input.txt")
	var grid Grid
	for y, line := range lines {
		var this_line []Pos
		for x, char := range line {
			var is_roll bool
			is_roll = string(char) == "@"
			this_line = append(this_line, Pos{x: x, y: y, roll: is_roll})
		}
		grid.data = append(grid.data, this_line)
	}

	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}
