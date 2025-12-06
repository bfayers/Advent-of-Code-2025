package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/bfayers/Advent-of-Code-2025/utils"
)

type Problem struct {
	numbers []int
	action  string
}

func process_problems(problems []Problem) int {
	// Execute each problem and sum the results
	var total int
	for _, problem := range problems {
		var sum int
		switch problem.action {
		case "+":
			for _, num := range problem.numbers {
				sum += num
			}
		case "*":
			sum = 1
			for _, num := range problem.numbers {
				sum *= num
			}
		}
		total += sum
	}
	return total
}

func parse_problems(lines []string) []Problem {
	// Process the input
	// Regex to find any number
	number_re := regexp.MustCompile(`\d+`)
	action_re := regexp.MustCompile(`[+\-/*]`)
	var problems []Problem
	for _, line := range lines {
		// Find all numbers in the line
		numbers_results := number_re.FindAllString(line, -1)

		if numbers_results != nil {
			// Process the numbers into the problems array
			for idx, num_str := range numbers_results {
				num, _ := strconv.Atoi(num_str)
				// If there is no problem at this index then create one
				if idx >= len(problems) {
					var new_problem Problem
					new_problem.numbers = []int{num}
					problems = append(problems, new_problem)
				} else {
					problems[idx].numbers = append(problems[idx].numbers, num)
				}
			}
		} else {
			// If there are no number sin the line then this line includes the action
			action_result := action_re.FindAllString(line, -1)
			for idx, action_str := range action_result {
				problems[idx].action = action_str
			}
		}
	}
	return problems
}

func parse_problems_p2(lines []string) []Problem {
	// In part 2 the numbers are actually written in columns rather than rows

	// Process the input
	action_re := regexp.MustCompile(`[+\-/*]`)
	var problems []Problem
	var problem_index int = 0
	// Iterate backwards through the lines to build the problems column-wise
	for char_idx := len(lines[0]) - 1; char_idx >= 0; char_idx-- {
		// char := string(line[char_idx])
		// Go down all the lines (except last one) at this character index to build the number
		var num_str string
		for scan_line_idx := 0; scan_line_idx < len(lines)-1; scan_line_idx++ {
			scan_line := lines[scan_line_idx]
			if char_idx < len(scan_line) {
				scan_char := string(scan_line[char_idx])
				if scan_char != " " {
					num_str += scan_char
				}
			}
		}
		num, _ := strconv.Atoi(num_str)
		// Add the number to the problems array
		if len(num_str) == 0 {
			problem_index++
			continue
		}
		if problem_index >= len(problems) {
			var new_problem Problem
			new_problem.numbers = []int{num}

			// Find the action from the last line (RTL)
			action_results := action_re.FindAllString(lines[len(lines)-1], -1)
			new_problem.action = action_results[len(action_results)-1-problem_index]

			problems = append(problems, new_problem)
		} else {
			problems[problem_index].numbers = append(problems[problem_index].numbers, num)
		}
	}
	return problems

}

func main() {
	// Load the files
	lines := utils.GetFileLines("input.txt")

	p1_problems := parse_problems(lines)
	p2_problems := parse_problems_p2(lines)

	fmt.Printf("Part 1: %d\n", process_problems(p1_problems))
	fmt.Printf("Part 2: %d\n", process_problems(p2_problems))
}
