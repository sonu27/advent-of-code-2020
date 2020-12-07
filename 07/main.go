package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLines("./07/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	data := make(map[string][]string)
	for _, v := range lines {
		a := strings.Split(v, " contain ")

		b := strings.Split(a[0], " ")
		bag := fmt.Sprintf("%s-%s", b[0], b[1])

		c := strings.Split(a[1], ", ")
		for _, w := range c {
			cs := strings.Split(w, " ")
			bb := fmt.Sprintf("%s-%s", cs[1], cs[2])

			if _, ok := data[bb]; ok {
				data[bb] = append(data[bb], bag)
			} else {
				data[bb] = []string{bag}
			}
		}
	}

	bags := make(map[string]int)
	find("shiny-gold", data, bags)
	fmt.Println(len(bags))
}

func find(bag string, data map[string][]string, bags map[string]int) {
	for _, v := range data[bag] {
		if _, ok := bags[v]; !ok {
			bags[v] = 1
		}
		find(v, data, bags)
	}
}

func part2(lines []string) {
	data := make(map[string][]string)
	for _, v := range lines {
		a := strings.Split(v, " contain ")

		b := strings.Split(a[0], " ")
		bag := fmt.Sprintf("%s-%s", b[0], b[1])
		data[bag] = []string{}

		c := strings.Split(a[1], ", ")
		for _, w := range c {
			cs := strings.Split(w, " ")
			q := s2i(cs[0])
			bb := fmt.Sprintf("%s-%s", cs[1], cs[2])

			for i := 0; i < q; i++ {
				data[bag] = append(data[bag], bb)
			}
		}
	}

	bags := make(map[string]int)
	find2("shiny-gold", data, bags)

	count := 0
	for _, v := range bags {
		count += v
	}
	fmt.Println(count)
}

func find2(bag string, data map[string][]string, bags map[string]int) {
	for _, v := range data[bag] {
		if _, ok := bags[v]; ok {
			bags[v]++
		} else {
			bags[v] = 1
		}
		find2(v, data, bags)
	}
}

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
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
