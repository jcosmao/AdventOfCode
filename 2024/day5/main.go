package main

import (
	"fmt"
	l "log/slog"
	"os"
	"strconv"
	"strings"

	"aoc/cli"
)

func main() {
	config := cli.ParseFlags()

	lines, err := cli.ReadFile(config.File)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch config.Part {
	case 1:
		processPart1(lines)
	case 2:
		processPart2(lines)
	default:
		os.Exit(1)
	}
}

// FileFormat
type FileFormat struct {
	Rules   map[int][]int
	Updates [][]int
}

func parseLines(lines []string) FileFormat {
	f := FileFormat{}

	f.Rules = map[int][]int{}

	sep := "|"
	for _, line := range lines {
		if len(line) == 0 {
			sep = ","
			continue
		}

		splitted := strings.Split(line, sep)
		if len(splitted) == 2 {
			a, _ := strconv.Atoi(splitted[0])
			b, _ := strconv.Atoi(splitted[1])
			f.Rules[a] = append(f.Rules[a], b)
		} else {
			var update []int
			for _, v := range splitted {
				r, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				update = append(update, r)
			}
			f.Updates = append(f.Updates, update)
		}
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)
	l.Info("", "file", file)

	var total int = 0
	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	fmt.Println(file)

	var total int = 0
	fmt.Println("part 2: ", int(total))
}
