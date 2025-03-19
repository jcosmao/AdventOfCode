package lib

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	Part      int
	File      string
	Profiling bool
	Display   bool
}

func ParseFlags() *Config {
	partPtr := flag.Int("part", 1, "Part number (1 or 2)")
	filePtr := flag.String("file", "input.txt", "Input text file")
	debugPtr := flag.Bool("debug", false, "Toggle debug")
	profiling := flag.Bool("profiling", false, "Toggle profiling")
	display := flag.Bool("display", false, "Display")

	flag.Parse()

	log.SetFlags(log.Lshortfile)
	if *debugPtr == true {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	return &Config{
		Part:      *partPtr,
		File:      *filePtr,
		Profiling: *profiling,
		Display:   *display,
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

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("[Timer] %s took %v\n", name, time.Since(start))
	}
}
