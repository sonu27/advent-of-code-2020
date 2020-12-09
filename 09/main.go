package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("./09/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(nums []int) int {
	p := 25

	for i := p; i < len(nums); i++ {
		if !find(nums[i], nums[i-p:i]) {
			return nums[i]
		}
	}
	return -1
}

func find(x int, nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == x {
				return true
			}
		}
	}
	return false
}

func part2(nums []int) int {
	x := 31161678

	for i := 0; i < len(nums)-1; i++ {
		n := nums[i]
		for j := i + 1; n < x && j < len(nums); j++ {
			n += nums[j]
			if n == x {
				low, high := getLowAndHigh(nums[i:j])
				return low + high
			}
		}
	}
	return -1
}

func getLowAndHigh(nums []int) (low int, high int) {
	low = math.MaxInt64
	high = math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] < low {
			low = nums[i]
		}
		if nums[i] > high {
			high = nums[i]
		}
	}
	return low, high
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
