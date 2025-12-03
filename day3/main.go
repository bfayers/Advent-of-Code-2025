package main

import (
	"fmt"
	"strconv"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

func findMaxCombination(bank string, max_length int) int {
	if len(bank) < max_length {
		return 0
	}
	var result string
	remaining := max_length
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

func part1(lines []string) int {
	var total_output_joltage int
	for _, bank := range lines {
		total_output_joltage += findMaxCombination(bank, 2)
	}
	return total_output_joltage
}

func part2(lines []string) int {
	var total_output_joltage int
	for _, bank := range lines {
		total_output_joltage += findMaxCombination(bank, 12)
	}
	return total_output_joltage
}

func main() {
	// Load data
	lines := utils.GetFileLines("input.txt")
	fmt.Printf("Part 1: %d\n", part1(lines))
	fmt.Printf("Part 2: %d\n", part2(lines))
}
