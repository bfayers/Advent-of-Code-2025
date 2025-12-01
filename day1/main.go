package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// func part1(data []string)

type Instruction struct {
	Direction string
	Amount    int
}

func part1(instructions []Instruction) int {
	// The dial has a lower bound of 0 and an upper bound of 99
	// Going below 0 wraps to 99, and above 99 wraps to 0
	// The dial starts pointed at 50
	// In part 1 we want to know how many times an instruction causes the dial to rest at 0
	position := 50
	zero_count := 0
	// fmt.Printf("The dial is at: %d\n", position)
	for _, instruction := range instructions {
		if instruction.Direction == "L" {
			position = position - instruction.Amount
			// fmt.Printf("The dial has rotated left by %d and is now at %d\n", instruction.Amount, position)
		}
		if instruction.Direction == "R" {
			position = position + instruction.Amount
			// fmt.Printf("The dial has rotated right by %d and is now at %d\n", instruction.Amount, position)
		}
		// Handle wrapping around
		// Keep adjusting position until it is within bounds
		for position < 0 || position > 99 {
			if position < 0 {
				position = 100 - (position * -1)
				// fmt.Printf("The dial wrapped around to %d\n", position)
			}
			if position > 99 {
				position = 0 + (position - 100)
				// fmt.Printf("The dial wrapped around to %d\n", position)
			}
		}
		// Check if we are at 0
		if position == 0 {
			zero_count++
			// fmt.Printf("The dial is at 0! Total zero count: %d\n", zero_count)
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
		for i := 0; i < instruction.Amount; i++ {
			if instruction.Direction == "L" {
				position = position - 1
				// fmt.Printf("The dial has rotated left by %d and is now at %d\n", instruction.Amount, position)
			}
			if instruction.Direction == "R" {
				position = position + 1
				// fmt.Printf("The dial has rotated right by %d and is now at %d\n", instruction.Amount, position)
			}
			// Handle wrapping around
			// Keep adjusting position until it is within bounds
			for position < 0 || position > 99 {
				if position < 0 {
					position = 100 - (position * -1)
					// fmt.Printf("The dial wrapped around to %d\n", position)
				}
				if position > 99 {
					position = 0 + (position - 100)
					// fmt.Printf("The dial wrapped around to %d\n", position)
				}
			}
			// Check if we are at 0
			if position == 0 {
				zero_count++
				// fmt.Printf("The dial is at 0! Total zero count: %d\n", zero_count)
			}
		}
	}
	return zero_count
}

func main() {
	// Load the data
	dat, _ := os.Open("input.txt")
	// Scan over the input to make an array of lines
	var instructions []Instruction
	scanner := bufio.NewScanner(dat)
	// Define regex to get the numbers
	number_re := regexp.MustCompile("[0-9]+")
	letter_re := regexp.MustCompile("[a-zA-Z]")
	for scanner.Scan() {
		this_line := scanner.Text()
		number, err := strconv.Atoi(number_re.FindString(this_line))
		if err != nil {
			fmt.Println("Error converting number:", err)
			continue
		}
		direction := letter_re.FindString(this_line)
		instructions = append(instructions, Instruction{Direction: direction, Amount: number})
	}

	// Part 1
	zero_count := part1(instructions)
	fmt.Printf("Part 1 - Zero count: %d\n", zero_count)
	// Part 2
	zero_count_2 := part2(instructions)
	fmt.Printf("Part 2 - Zero count: %d\n", zero_count_2)
}
