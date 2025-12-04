package main

import (
	"fmt"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

type Grid struct {
	data [][]bool
}

func check_accessible(grid Grid, x int, y int) bool {
	// Check adjacent spaces
	adjacent := 0

	// Check up/down/left/right
	// Left
	if x > 0 && grid.data[y][x-1] {
		adjacent++
	}
	// Right
	if x < len(grid.data[y])-1 && grid.data[y][x+1] {
		adjacent++
	}
	// Up
	if y > 0 && grid.data[y-1][x] {
		adjacent++
	}
	// Down
	if y < len(grid.data)-1 && grid.data[y+1][x] {
		adjacent++
	}

	// Check diagonal spaces
	// Up Left
	if x > 0 && y > 0 && grid.data[y-1][x-1] {
		adjacent++
	}
	// Up Right
	if x < len(grid.data[y])-1 && y > 0 && grid.data[y-1][x+1] {
		adjacent++
	}
	// Down Left
	if x > 0 && y < len(grid.data)-1 && grid.data[y+1][x-1] {
		adjacent++
	}
	// Down Right
	if x < len(grid.data[y])-1 && y < len(grid.data)-1 && grid.data[y+1][x+1] {
		adjacent++
	}

	// If less than 4 spaces around are filled, it is accessible
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
			if pos {
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
				if pos {
					if check_accessible(grid, x, y) {
						grid.data[y][x] = false
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
	for _, line := range lines {
		var this_line []bool
		for _, char := range line {
			var is_roll bool
			is_roll = string(char) == "@"
			this_line = append(this_line, is_roll)
		}
		grid.data = append(grid.data, this_line)
	}

	fmt.Printf("Part 1: %d\n", part1(grid))
	fmt.Printf("Part 2: %d\n", part2(grid))
}
