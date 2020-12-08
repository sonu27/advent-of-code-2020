package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name string
	val  int
	line int
}

func main() {
	lines, err := readLines("./08/input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	fmt.Println(part1(lines))
	fmt.Println(part2(lines))
}

func part1(ins []Instruction) int {
	acc, _, _ := Prog(ins)

	return acc
}

func part2(ins []Instruction) int {
	acc, infiniteLoop, _ := Prog(ins)

	for i := 0; i < len(ins); i++ {
		tmp := make([]Instruction, len(ins))
		copy(tmp, ins)

		if ins[i].name == "jmp" {
			tmp[i].name = "nop"
		} else if ins[i].name == "nop" {
			tmp[i].name = "jmp"
		} else {
			continue
		}

		acc, infiniteLoop, _ = Prog(tmp)
		if !infiniteLoop {
			break
		}
	}

	return acc

}

func Prog(ins []Instruction) (int, bool, []Instruction) {
	var stack []Instruction
	done := make(map[int]Instruction)
	acc := 0
	i := 0
	for i < len(ins) {
		if _, ok := done[i]; ok && stack[len(stack)-1].line != len(stack)-1 {
			return acc, true, stack
		}
		stack = append(stack, ins[i])
		done[i] = ins[i]

		switch ins[i].name {
		case "acc":
			acc += ins[i].val
		case "jmp":
			i += ins[i].val
			continue
		case "nop":
		}
		i++
	}

	return acc, false, stack
}

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func readLines(path string) ([]Instruction, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []Instruction
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		lines = append(lines, Instruction{
			name: s[0],
			val:  s2i(s[1]),
			line: i,
		})
		i++
	}
	return lines, scanner.Err()
}
