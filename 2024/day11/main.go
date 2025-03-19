package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"strings"
)

// Day11
type Day11 struct {
	Line        []int
	IterResults []int
}

// Rules:
// 0 -> 1
// even digit -> splitted in 2 numbers. Ex:  1000 -> 10 00 (00 -> 0)
// default ->  number muliplied by 2024. Ex:  1 -> 2024

func applyRules(n int) []int {
	result := []int{}

	toDigits := aoc.IntToDigits(n)
	if n == 0 {
		result = append(result, 1)
	} else if aoc.IsPair(len(toDigits)) {
		split := len(toDigits) / 2
		first := aoc.ConcatInts(toDigits[0:split]...)
		sec := aoc.ConcatInts(toDigits[split:]...)
		result = append(result, first)
		result = append(result, sec)
	} else {
		result = append(result, n*2024)
	}

	return result
}

func processPart2(lines []string) {
	file := parseLines(lines)

	ResultCounter := map[int]int{}
	for _, res := range file.Line {
		if _, exists := ResultCounter[res]; !exists {
			ResultCounter[res] = 0
		}
		ResultCounter[res] += 1
	}

	for i := 0; i < 75; i++ {
		l.Info("iter", "", i+1)
		newres := map[int]int{}

		for k, v := range ResultCounter {
			for _, res := range applyRules(k) {
				if _, exists := newres[res]; !exists {
					newres[res] = v
				} else {
					newres[res] += v
				}
			}
		}

		ResultCounter = newres
		l.Info("Blink result", "", ResultCounter)
	}

	total := 0
	for _, v := range ResultCounter {
		total += v
	}
	l.Info("Result", "part 2", total)
}

func processPart1(lines []string) {

	file := parseLines(lines)
	file.IterResults = []int{}

	l.Debug("", "file", file)
	toConvert := file.Line

	for i := 0; i < 25; i++ {
		l.Info("iter", "", i+1)
		file.IterResults = []int{}
		l.Debug("toconvert", "", toConvert)
		for _, v := range toConvert {
			for _, res := range applyRules(v) {
				file.IterResults = append(file.IterResults, res)
			}
		}
		toConvert = file.IterResults
	}

	l.Info("FileFormat", "", file)
	total := len(file.IterResults)
	fmt.Println("part 1: ", int(total))
}

func parseLines(lines []string) Day11 {

	f := Day11{}

	for _, line := range lines {
		splitted := strings.Fields(line)

		for i := 0; i < len(splitted); i++ {
			f.Line = append(f.Line, aoc.StringToInt(splitted[i]))
		}
	}

	return f
}

func main() {
	config := aoc.ParseFlags()
	lines := aoc.ReadFile(config.File)

	if config.Profiling {
		_ = os.Remove("cpu.prof")
		cpuf, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(cpuf)
		defer cpuf.Close()
		defer pprof.StopCPUProfile()
		l.Info("[profiling] go tool pprof cpu.prof")
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
