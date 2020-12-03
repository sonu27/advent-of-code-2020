package main

import (
	"bufio"
	"os"
	"strings"
)

var lines [][]string

func main() {
	lines, _ = readLines("./03/input.txt")

	work := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	count := 1

	for _, v := range work {
		count *= treesForSlope(v[0], v[1])
	}

	println(count)
}

func treesForSlope(right int, down int) int {
	trees := 0
	x, y := right, down

	for y < len(lines) {
		if lines[y][x] == "#" {
			trees++
		}

		x += right
		if x+1 > len(lines[y]) {
			x = x - len(lines[y])
		}

		y += down
	}

	return trees
}

func readLines(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		lines = append(lines, line)
	}
	return lines, scanner.Err()
}
