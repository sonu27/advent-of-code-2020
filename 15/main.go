package main

import "fmt"

func main() {
	input := []int{12, 1, 16, 3, 11, 0}

	fmt.Printf("Part 1: %d", solve(input, 2020))
	fmt.Printf("Part 2: %d", solve(input, 30000000))
}

func solve(input []int, untilTurn int) int {
	spoken := make([]int, untilTurn)
	turn := 1

	for _, v := range input {
		spoken[v] = turn
		turn++
	}

	var last int
	speak := input[len(input)-1]

	for ; turn <= untilTurn; turn++ {
		last = speak

		if t := spoken[speak]; t != 0 {
			speak = turn - t
		} else {
			speak = 0
		}

		spoken[last] = turn
	}

	return last
}
