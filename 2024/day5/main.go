package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"slices"
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
			a := aoc.StringToInt(splitted[0])
			b := aoc.StringToInt(splitted[1])
			f.Rules[a] = append(f.Rules[a], b)
		} else {
			var update []int
			for _, v := range splitted {
				update = append(update, aoc.StringToInt(v))
			}
			f.Updates = append(f.Updates, update)
		}
	}

	return f
}

func processPart1(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	valid := [][]int{}
	for _, update := range file.Updates {
		if ok, _ := isUpdateValid(update, file.Rules); ok {
			valid = append(valid, update)
		}
	}
	l.Debug("", "valid", valid)

	var total int = 0
	for _, update := range valid {
		middleIndex := len(update) / 2
		total += update[middleIndex]
	}
	fmt.Println("part 1: ", int(total))
}

func processPart2(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	valid := [][]int{}
	for _, update := range file.Updates {
		corrected := false
		for {
			ok, updateCorrected := isUpdateValid(update, file.Rules)
			update = updateCorrected
			if ok {
				break
			} else {
				corrected = true
			}
		}

		if corrected {
			valid = append(valid, update)
		}
	}
	l.Debug("", "valid/corrected", valid)

	var total int = 0
	for _, update := range valid {
		middleIndex := len(update) / 2
		total += update[middleIndex]
	}

	fmt.Println("part 2: ", int(total))
}

func isUpdateValid(update []int, rules map[int][]int) (bool, []int) {
	l.Debug("Check update valid", "update", update)
	for i := 0; i < len(update)-1; i++ {
		// for each value in []update (except last one because it's the last...),
		// check if next values does not have priority over it
		if res, corrected := isRulesViolated(update[i], update[i+1:], rules); res {
			updateCorrected := slices.Concat(update[0:i], corrected)
			l.Debug("", "updateCorrected", updateCorrected)
			return false, updateCorrected
		}
	}
	return true, update
}

func isRulesViolated(ref int, tocheck []int, rules map[int][]int) (bool, []int) {
	for i, val := range tocheck {
		if slices.Contains(rules[val], ref) {
			l.Debug("Rule violated", "reason", fmt.Sprintf("%d must be before %d", val, ref))
			rule := []int{val, ref}
			corrected := slices.Concat(rule, tocheck[0:i], tocheck[i+1:])
			return true, corrected
		}
	}
	return false, []int{}
}
