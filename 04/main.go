package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var lines []map[string]string

func main() {
	lines, _ = readLines("./04/input.txt")

	count := 0
	for _, v := range lines {
		if validPassport1(v) {
			count++
		}
	}
	fmt.Println(count)

	count = 0
	for _, v := range lines {
		if validPassport2(v) {
			count++
		}
	}

	fmt.Println(count)
}

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func validPassport1(pp map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, r := range required {
		if _, ok := pp[r]; !ok {
			return false
		}
	}

	return true
}

func validPassport2(pp map[string]string) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	fn := map[string]func(string) bool{
		"byr": func(s string) bool {
			n := s2i(s)
			return 1920 <= n && n <= 2002
		},
		"iyr": func(s string) bool {
			n := s2i(s)
			return 2010 <= n && n <= 2020
		},
		"eyr": func(s string) bool {
			n := s2i(s)
			return 2020 <= n && n <= 2030
		},
		"hgt": func(s string) bool {
			if strings.HasSuffix(s, "cm") {
				t := strings.Replace(s, "cm", "", 1)
				i, err := strconv.Atoi(t)
				if err != nil {
					return false
				}
				if 150 <= i && i <= 193 {
					return true
				}
			}
			if strings.HasSuffix(s, "in") {
				t := strings.Replace(s, "in", "", 1)
				i, err := strconv.Atoi(t)
				if err != nil {
					return false
				}
				if 59 <= i && i <= 76 {
					return true
				}
			}
			return false
		},
		"hcl": func(s string) bool {
			matched, _ := regexp.MatchString(`^#[a-f0-9]{6}$`, s)
			return matched
		},
		"ecl": func(s string) bool {
			validValues := map[string]bool{
				"amb": true,
				"blu": true,
				"brn": true,
				"gry": true,
				"grn": true,
				"hzl": true,
				"oth": true,
			}

			_, ok := validValues[s]
			return ok
		},
		"pid": func(s string) bool {
			matched, _ := regexp.MatchString(`^[0-9]{9}$`, s)
			return matched
		},
	}

	for _, r := range required {
		if v, ok := pp[r]; ok {
			if !fn[r](v) {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func readLines(path string) ([]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []map[string]string
	var lastLine []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")

		if len(fields) == 1 && fields[0] == "" {
			m := map[string]string{}
			for _, v := range lastLine {
				if v == "" {
					continue
				}
				w := strings.Split(v, ":")
				m[w[0]] = w[1]
			}
			lines = append(lines, m)
			lastLine = []string{}
		}

		lastLine = append(lastLine, fields...)
	}

	m := map[string]string{}
	for _, v := range lastLine {
		if v == "" {
			continue
		}
		w := strings.Split(v, ":")
		m[w[0]] = w[1]
	}
	lines = append(lines, m)

	return lines, scanner.Err()
}
