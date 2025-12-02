package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

type id_range struct {
	min int
	max int
}

func check_repeated(s string, max_repeats int) bool {
	for j := 1; j <= len(s)/2; j++ {
		first_part := s[:j]
		repeat_count := strings.Count(s, first_part)

		if repeat_count*len(first_part) == len(s) && repeat_count <= max_repeats {
			return true
		}
	}
	return false
}

func part1(id_ranges []id_range) int {
	// We need to identify repeating patterns of numbers, they could be as little 1 digit twice (eg: 55 = 5 twice) but no upper bound on that behaviour
	// The 'invalid' ids are ONLY made of the repeating numbers TWICE (eg: 824824 is invalid but 824824824 is not)
	var total_invalid int
	for _, id_range := range id_ranges {
		for i := id_range.min; i <= id_range.max; i++ {
			// Get the string representation of the number
			num_str := strconv.Itoa(i)
			if check_repeated(num_str, 2) {
				total_invalid += i
			}
		}
	}
	return total_invalid
}

func part2(id_ranges []id_range) int {
	// Same as part 1, identifying repeating patterns, but now the 'invalid' ids are repeated _AT LEAST_ twice (eg: 824824 is invalid, as is 824824824)
	var total_invalid int
	for _, id_range := range id_ranges {
		// Get the string representation of the number
		for i := id_range.min; i <= id_range.max; i++ {
			num_str := strconv.Itoa(i)
			// possible digits for 'repeating' can only be up to half the length of the string,
			// find all possible repeating sequences and check if the string is made entirely of that
			if check_repeated(num_str, len(num_str)) {
				total_invalid += i
			}
		}
	}
	return total_invalid
}

func main() {
	// Load the data
	lines := utils.GetFileLines("input.txt")
	var id_ranges []id_range
	for _, line := range lines {
		for range_str := range strings.SplitSeq(line, ",") {
			range_values := strings.Split(range_str, "-")
			min, _ := strconv.Atoi(range_values[0])
			max, _ := strconv.Atoi(range_values[1])
			id_ranges = append(id_ranges, id_range{min: min, max: max})
		}
	}

	fmt.Printf("Part 1: %d\n", part1(id_ranges))
	fmt.Printf("Part 2: %d\n", part2(id_ranges))
}
