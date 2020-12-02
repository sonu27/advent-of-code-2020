package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("./01/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []int) {
	for i := 0; i < len(lines)-1; i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i]+lines[j] == 2020 {
				fmt.Println(lines[i] * lines[j])
				return
			}
		}
	}
}

func part2(lines []int) {
	for i := 0; i < len(lines)-2; i++ {
		for j := i + 1; j < len(lines)-1; j++ {
			for k := j + 1; k < len(lines); k++ {
				if lines[i]+lines[j]+lines[k] == 2020 {
					fmt.Println(lines[i] * lines[j] * + lines[k])
					return
				}
			}
		}
	}
}

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, i)
	}
	return lines, scanner.Err()
}
