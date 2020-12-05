package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	lines, err := readLines("./05/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var seats []int
	high := 0
	for _, line := range lines {
		a := get1(line)
		if a > high {
			high = a
		}

		seats = append(seats, a)
	}

	fmt.Printf("Part 1: %d\n", high)

	sort.Ints(seats)

	for i := 1; i < len(seats); i++ {
		if seats[i] != seats[i-1]+1 {
			fmt.Printf("Part 2: %d\n", seats[i-1]+1)
			break
		}
	}
}

func get1(s string) int {
	ss := strings.Split(s, "")
	low, high := 0, 127
	low2, high2 := 0, 7
	for i, v := range ss {
		if i < 7 {
			if v == "F" {
				high = (high + low) / 2
			} else {
				low = ((high + low) / 2) + 1
			}
		} else {
			if v == "L" {
				high2 = (high2 + low2) / 2
			} else {
				low2 = ((high2 + low2) / 2) + 1
			}
		}
	}

	return high*8 + high2
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
