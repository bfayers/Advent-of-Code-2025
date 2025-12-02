package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

// func part1(data []string)

type Instruction struct {
	Direction string
	Amount    int
}

func calculate_new_position(position int, move int) int {
	return (position + move + 100) % 100
}

func part1(instructions []Instruction) int {
	// The dial has a lower bound of 0 and an upper bound of 99
	// Going below 0 wraps to 99, and above 99 wraps to 0
	// The dial starts pointed at 50
	// In part 1 we want to know how many times an instruction causes the dial to rest at 0
	position := 50
	zero_count := 0
	for _, instruction := range instructions {
		position = calculate_new_position(position, instruction.Amount)
		// Check if we are at 0
		if position == 0 {
			zero_count++
		}
	}
	return zero_count
}

func part2(instructions []Instruction) int {
	// The dial has a lower bound of 0 and an upper bound of 99
	// Going below 0 wraps to 99, and above 99 wraps to 0
	// The dial starts pointed at 50
	// In part 2 we want to know how many times any 'click' causes the dial to point at 0
	position := 50
	zero_count := 0
	// fmt.Printf("The dial is at: %d\n", position)
	for _, instruction := range instructions {
		if instruction.Amount < 0 {
			instruction.Amount = instruction.Amount * -1
		}
		for i := 0; i < instruction.Amount; i++ {
			if instruction.Direction == "L" {
				position = calculate_new_position(position, -1)
			}
			if instruction.Direction == "R" {
				position = calculate_new_position(position, 1)
			}
			// Check if we are at 0
			if position == 0 {
				zero_count++
			}
		}
	}
	return zero_count
}

func main() {
	// Load the data
	lines := utils.GetFileLines("input.txt")
	var instructions []Instruction
	// Define regex to get the numbers
	number_re := regexp.MustCompile("[0-9]+")
	letter_re := regexp.MustCompile("[a-zA-Z]")
	for _, line := range lines {
		number, err := strconv.Atoi(number_re.FindString(line))
		if err != nil {
			fmt.Println("Error converting number:", err)
			continue
		}
		direction := letter_re.FindString(line)
		if direction == "L" {
			instructions = append(instructions, Instruction{Direction: "L", Amount: -number})
		} else {
			instructions = append(instructions, Instruction{Direction: "R", Amount: number})
		}
	}

	// Part 1
	zero_count := part1(instructions)
	fmt.Printf("Part 1 - Zero count: %d\n", zero_count)
	// Part 2
	zero_count_2 := part2(instructions)
	fmt.Printf("Part 2 - Zero count: %d\n", zero_count_2)
}
