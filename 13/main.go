package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("./13/input.txt")
	lines := strings.Split(string(b), "\n")
	num := s2i(lines[0])
	buses := strings.Split(lines[1], ",")

	part1(num, buses)
	part2(buses)
}

func part1(num int, s []string) {
	var buses []int
	for _, v := range s {
		if v != "x" {
			buses = append(buses, s2i(v))
		}
	}

	time := num
	for {
		for _, bus := range buses {
			if time%bus == 0 {
				fmt.Println(bus * (time - num))
				return
			}
		}
		time++
	}
}

func part2(s []string) {

}

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
