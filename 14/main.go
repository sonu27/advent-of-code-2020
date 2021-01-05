package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memory map[int64]string

func main() {
	memory = make(map[int64]string)
	_ = readLines("./14/input.txt")
}

func readLines(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var mask []string
	r, _ := regexp.Compile(`^mem\[(\d+)\]`)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " = ")
		if fields[0] == "mask" {
			mask = strings.Split(fields[1], "")
		} else {
			memAddr := s2i(r.FindStringSubmatch(fields[0])[1])
			num := fmt.Sprintf("%036b", s2i(fields[1]))
			nums := strings.Split(num, "")
			for i := 0; i < 36; i++ {
				if mask[i] == "0" {
					nums[i] = "0"
				} else if mask[i] == "1" {
					nums[i] = "1"
				}
			}

			memory[memAddr] = strings.Join(nums, "")
		}
	}

	var sum int64
	for _, v := range memory {
		i, _ := strconv.ParseInt(v, 2, 64)
		sum += i
	}

	fmt.Println(sum)

	return scanner.Err()
}

func s2i(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
