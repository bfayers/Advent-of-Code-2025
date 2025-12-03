package main

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

func part1(lines []string) int {
	var total_output_joltage int
	for _, bank := range lines {
		// Find all 2 battery pairings, in the right order
		var pairings []int
		for i := 0; i <= len(bank)-1; i++ {
			for j := i + 1; j <= len(bank)-1; j++ {
				this_pairing, _ := strconv.Atoi(string(bank[i]) + string(bank[j]))
				pairings = append(pairings, this_pairing)
			}
		}
		total_output_joltage += slices.Max(pairings)
	}
	return total_output_joltage
}

func findMaxCombination(bank string) int {
	if len(bank) < 12 {
		return 0
	}
	var result string
	remaining := 12
	for i := 0; i < len(bank) && remaining > 0; i++ {
		if len(bank)-i == remaining {
			result += bank[i:]
			break
		}
		maxDigit := bank[i]
		for j := i + 1; j <= len(bank)-remaining; j++ {
			if bank[j] > maxDigit {
				maxDigit = bank[j]
			}
		}
		for bank[i] != maxDigit {
			i++
		}
		result += string(bank[i])
		remaining--
	}
	val, _ := strconv.Atoi(result)
	return val
}

func part2(lines []string) int {
	var total_output_joltage int
	for _, bank := range lines {
		total_output_joltage += findMaxCombination(bank)
	}
	return total_output_joltage
}

func main() {
	// Load data
	lines := utils.GetFileLines("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
