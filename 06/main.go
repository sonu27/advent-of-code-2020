package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines, err := readLines("./06/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(lines []string) int {
	total := 0

	count := 0
	uniqueChars := make(map[string]bool)

	for _, v := range lines {
		c := strings.Split(v, "")

		if len(c) == 0 {
			total += count
			count = 0
			uniqueChars = make(map[string]bool)
		}

		for _, w := range c {
			if _, ok := uniqueChars[w]; !ok {
				count++
				uniqueChars[w] = true
			}
		}
	}

	total += count

	return total
}

func part2(lines []string) int {
	total := 0

	groupNum := 0
	uniqueChars := make(map[string]int)

	for _, v := range lines {
		c := strings.Split(v, "")

		if len(c) > 0 {
			groupNum++
			for _, w := range c {
				if _, ok := uniqueChars[w]; !ok {
					uniqueChars[w] = 1
				} else {
					uniqueChars[w]++
				}
			}
		} else {
			for _, x := range uniqueChars {
				if x == groupNum {
					total++
				}
			}

			groupNum = 0
			uniqueChars = make(map[string]int)
		}
	}

	for _, x := range uniqueChars {
		if x == groupNum {
			total++
		}
	}

	return total
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
