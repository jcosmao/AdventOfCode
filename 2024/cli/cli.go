package cli

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
)

type Config struct {
	Part int
	File string
}

func ParseFlags() Config {
	partPtr := flag.Int("part", 1, "Part number (1 or 2)")
	filePtr := flag.String("file", "input.txt", "Input text file")
	debugPtr := flag.Bool("debug", false, "Toggle debug")

	flag.Parse()

	log.SetFlags(log.Lshortfile)
	if *debugPtr == true {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	return Config{
		Part: *partPtr,
		File: *filePtr,
	}
}

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	return lines, nil
}
