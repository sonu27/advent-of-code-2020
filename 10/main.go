package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	lines, err := readLines("./10/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	part1(lines)
	part2(lines)
}

func part1(nums []int) {
	v1 := 1
	v3 := 1
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 1 {
			v1++
		} else if diff == 3 {
			v3++
		} else if diff > 3 {
			break
		}
	}

	fmt.Println(v1 * v3)
}

func part2(nums []int) {
	nums = append([]int{0}, nums...)
	nums = append(nums, nums[len(nums)-1]+3)

	options := make(map[int][]int)
	for i := 0; i < len(nums); i++ {
		var next []int
		for j := i + 1; j < len(nums) && nums[j]-nums[i] <= 3; j++ {
			next = append(next, nums[j])
		}
		options[nums[i]] = next
	}

	last := nums[len(nums)-1]
	cache := make(map[int]int64)
	v := permutations(0, options, last, cache)
	fmt.Println(v)
}

func permutations(i int, options map[int][]int, last int, cache map[int]int64) int64 {
	var total int64

	if i == last {
		return 1
	}

	for _, v := range options[i] {
		if cache[v] == 0 {
			total += permutations(v, options, last, cache)
		} else {
			total += cache[v]
		}
	}
	cache[i] = total
	return total
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
	sort.Ints(lines)
	return lines, scanner.Err()
}
