package aoc

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Part      int
	File      string
	Profiling bool
}

func ParseFlags() Config {
	partPtr := flag.Int("part", 1, "Part number (1 or 2)")
	filePtr := flag.String("file", "input.txt", "Input text file")
	debugPtr := flag.Bool("debug", false, "Toggle debug")
	profiling := flag.Bool("profiling", false, "Toggle profiling")

	flag.Parse()

	log.SetFlags(log.Lshortfile)
	if *debugPtr == true {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	return Config{
		Part: *partPtr,
		File: *filePtr,
		Profiling: *profiling,
	}
}

func ReadFile(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

// "1234" -> [1, 2, 3 ,4]
func StringToInts(input string) []int {
	fields := strings.Split(input, "")
	output := make([]int, len(fields))
	for i := 0; i < len(fields); i++ {
		output[i] = StringToInt(fields[i])
	}
	return output
}

func IntToSlice(n int64, sequence []int64) []int64 {
    if n != 0 {
        i := n % 10
        // sequence = append(sequence, i) // reverse order output
        sequence = append([]int64{i}, sequence...)
        return IntToSlice(n/10, sequence)
    }
    return sequence
}

func ConcatInts(a int, b int) int {
	// Calculate the number of digits in b
	bDigits := int(math.Log10(float64(b))) + 1

	// Shift a to the left by the number of digits in b
	aShifted := a * int(math.Pow(10, float64(bDigits)))

	// Add b to the shifted a
	return aShifted + b
}

func Timer(name string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("[Timer] %s took %v\n", name, time.Since(start))
    }
}
