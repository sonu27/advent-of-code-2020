package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines, err := readLines("./02/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	valid1 := 0
	valid2 := 0
	r, _ := regexp.Compile("^(\\d+)\\-(\\d+) (\\w): (\\w+)$")
	for _, s := range lines {
		m := r.FindStringSubmatch(s)
		if validPassword1(m[4], s2i(m[1]), s2i(m[2]), []rune(m[3])[0]) {
			valid1++
		}
		if validPassword2(m[4], s2i(m[1]), s2i(m[2]), []rune(m[3])[0]) {
			valid2++
		}
	}

	fmt.Println(valid1, valid2)
}

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func validPassword1(pw string, min int, max int, char rune) bool {
	count := 0
	for _, s := range pw {
		if s == char {
			count++
			if count > max {
				return false
			}
		}
	}

	return count >= min
}

func validPassword2(pw string, pos1 int, pos2 int, char rune) bool {
	pwr := []rune(pw)
	l := len(pwr)

	validPos1 := false
	validPos2 := false

	if pos1 <= l {
		validPos1 = pwr[pos1-1] == char
	}

	if pos2 <= l {
		validPos2 = pwr[pos2-1] == char
	}

	return (validPos1 && !validPos2) || (!validPos1 && validPos2)
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
