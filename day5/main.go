package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part1(fresh_ranges [][]int, available_ingredients []int) int {
	var fresh_count int
	for _, ingredient := range available_ingredients {
		for _, fresh_range := range fresh_ranges {
			if ingredient >= fresh_range[0] && ingredient <= fresh_range[1] {
				fresh_count++
				break
			}
		}
	}
	return fresh_count
}

func part2(fresh_ranges [][]int) int {
	// Attempt to merge the ranges
	var merged_ranges [][]int
	// Start by sorting the current ranges by their lower bound
	slices.SortFunc(fresh_ranges, func(a, b []int) int {
		return a[0] - b[0]
	})
	for _, fresh_range := range fresh_ranges {
		if len(merged_ranges) == 0 {
			merged_ranges = append(merged_ranges, fresh_range)
			continue
		}
		last_range := merged_ranges[len(merged_ranges)-1]
		// Check if the current range overlaps the last merged range
		if fresh_range[0] <= last_range[1] {
			// Merge the ranges by extending the upper bound if needed
			if fresh_range[1] > last_range[1] {
				last_range[1] = fresh_range[1]
			}
		} else {
			merged_ranges = append(merged_ranges, fresh_range)
		}
	}

	// Now find the total number of IDs in the merged ranges
	var total_count int
	for _, merged_range := range merged_ranges {
		total_count += (merged_range[1] - merged_range[0] + 1)
	}
	return total_count
}

func main() {
	// Load the data
	dat, _ := os.Open("input.txt")
	// Scan over the input to make an array of lines
	var fresh_ranges [][]int
	var available_ingredients []int
	scanner := bufio.NewScanner(dat)
	var reading_ingredients bool = false
	for scanner.Scan() {
		// After a blank line switch to reading the available ingredients
		this_line := scanner.Text()
		if this_line == "" {
			reading_ingredients = true
			continue
		}
		if !reading_ingredients {
			var this_range []int
			line_split := strings.Split(this_line, "-")
			lower, _ := strconv.Atoi(line_split[0])
			upper, _ := strconv.Atoi(line_split[1])
			this_range = []int{lower, upper}
			fresh_ranges = append(fresh_ranges, this_range)
		} else {
			ingredient, _ := strconv.Atoi(this_line)
			available_ingredients = append(available_ingredients, ingredient)
		}
	}
	// fmt.Println(fresh_ranges)
	// fmt.Println(available_ingredients)

	fmt.Printf("Part 1: %d\n", part1(fresh_ranges, available_ingredients))
	fmt.Printf("Part 2: %d\n", part2(fresh_ranges))
}
