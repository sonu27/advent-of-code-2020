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
	ins, err := readLines("./12/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	part1(ins)
	part2(ins)
}

type Ins struct {
	Ins   string
	Value int
}

func part1(ins []Ins) {
	dir := 90
	x, y := 0, 0

	for _, v := range ins {
		switch v.Ins {
		case "N":
			y += v.Value
		case "S":
			y -= v.Value
		case "E":
			x += v.Value
		case "W":
			x -= v.Value
		case "L":
			dir -= v.Value
		case "R":
			dir += v.Value
		case "F":
			switch dir {
			case 0:
				y += v.Value
			case 90:
				x += v.Value
			case 180:
				y -= v.Value
			case 270:
				x -= v.Value
			}
		}

		if dir < 0 {
			dir = 360 + dir
		}
		if dir > 270 {
			dir = dir - 360
		}
	}

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

func part2(ins []Ins) {
	x, y := 0, 0
	wx, wy := 10, 1

	for _, v := range ins {
		switch v.Ins {
		case "N":
			wy += v.Value
		case "S":
			wy -= v.Value
		case "E":
			wx += v.Value
		case "W":
			wx -= v.Value
		case "L":
			v.Value = mod(-v.Value, 360)
			fallthrough
		case "R":
			switch v.Value {
			case 90:
				wx, wy = +wy, -wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = -wy, +wx
			}
		case "F":
			x += wx * v.Value
			y += wy * v.Value
		}
	}

	fmt.Println(math.Abs(float64(x + y)))
}

func mod(val, m int) int {
	val = val % m
	if val < 0 {
		val = m + val
	}
	return val
}

func readLines(path string) ([]Ins, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ins []Ins
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		i, _ := strconv.Atoi(s[1:])
		ins = append(ins, Ins{
			Ins:   s[:1],
			Value: i,
		})
	}
	return ins, scanner.Err()
}
