package main

import (
	"aoc"
	"fmt"
	l "log/slog"
	"os"
	"reflect"
	"strings"
)

func main() {
	config := aoc.ParseFlags()
	lines := aoc.ReadFile(config.File)

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
	Rows int
	Cols int
	data [][]string
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	f.Rows = len(lines)

	for i, line := range lines {
		splitted := strings.Split(line, "")
		f.data = append(f.data, splitted)

		if i == 0 {
			f.Cols = len(splitted)
		}
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)
	var total int = 0
	for row, line := range file.data {
		for col, letter := range line {
			if letter == "X" {
				total += findXMAS(row, col, file)
			}
		}
	}

	fmt.Println("part 1: ", int(total))
}

func findXMAS(r int, c int, f FileFormat) int {

	var word = []string{"X", "M", "A", "S"}
	var found = 0

	l.Debug("findXMAS", "row", r, "col", c)

	// find right
	if (c + 3) < f.Cols {
		w := []string{"X", f.data[r][c+1], f.data[r][c+2], f.data[r][c+3]}
		if reflect.DeepEqual(word, w) {
			found += 1
			l.Debug("found right")
		}
	}
	// find left
	if (c - 3) >= 0 {
		w := []string{"X", f.data[r][c-1], f.data[r][c-2], f.data[r][c-3]}
		if reflect.DeepEqual(word, w) {
			found += 1
			l.Debug("found left")
		}
	}
	// find down
	if (r + 3) < f.Rows {
		w := []string{"X", f.data[r+1][c], f.data[r+2][c], f.data[r+3][c]}
		if reflect.DeepEqual(word, w) {
			found += 1
			l.Debug("found down")
		}
	}
	// find up
	if r-3 >= 0 {
		w := []string{"X", f.data[r-1][c], f.data[r-2][c], f.data[r-3][c]}
		if reflect.DeepEqual(word, w) {
			found += 1
			l.Debug("found up")
		}
	}
	// find right/up
	if (c+3) < f.Cols && (r-3) >= 0 {
		w := []string{"X", f.data[r-1][c+1], f.data[r-2][c+2], f.data[r-3][c+3]}
		if reflect.DeepEqual(word, w) {
			found += 1
			l.Debug("found right/up")
		}
	}
	// find right/down
	if (c+3) < f.Cols && (r+3) < f.Rows {
		w := []string{"X", f.data[r+1][c+1], f.data[r+2][c+2], f.data[r+3][c+3]}
		if reflect.DeepEqual(word, w) {
			l.Debug("found right/down")
			found += 1
		}
	}
	// find left/up
	if (c-3) >= 0 && (r-3) >= 0 {
		w := []string{"X", f.data[r-1][c-1], f.data[r-2][c-2], f.data[r-3][c-3]}
		if reflect.DeepEqual(word, w) {
			l.Debug("found left/up")
			found += 1
		}
	}
	// find left/down
	if (c-3) >= 0 && (r+3) < f.Rows {
		w := []string{"X", f.data[r+1][c-1], f.data[r+2][c-2], f.data[r+3][c-3]}
		if reflect.DeepEqual(word, w) {
			l.Debug("found left/down")
			found += 1
		}
	}

	return found
}

func processPart2(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	var total int = 0

	for row, line := range file.data {
		for col, letter := range line {
			if letter == "A" {
				total += findXMAS2(row, col, file)
			}
		}
	}

	fmt.Println("part 2: ", int(total))
}

func findXMAS2(r int, c int, f FileFormat) int {

	word := []string{"M", "A", "S"}
	drow := []string{"S", "A", "M"}

	if c == 0 || c == f.Cols-1 || r == 0 || r == f.Rows-1 {
		return 0
	}

	cross1 := []string{f.data[r-1][c-1], "A", f.data[r+1][c+1]}
	cross2 := []string{f.data[r+1][c-1], "A", f.data[r-1][c+1]}

	if (reflect.DeepEqual(cross1, word) || reflect.DeepEqual(cross1, drow)) && (reflect.DeepEqual(cross2, word) || reflect.DeepEqual(cross2, drow)) {
		l.Debug("found")
		return 1
	}

	return 0
}
