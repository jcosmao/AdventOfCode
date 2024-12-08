package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"math"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"sync"
)

type Operation struct {
	result   int
	products []int
}

const (
	ADD    string = "0"
	MUL    string = "1"
	CONCAT string = "2"
)

// FileFormat
type FileFormat struct {
	lines []Operation
}

func CanSolve(op Operation, Operators []string) bool {
	nops := len(Operators)
	neededOperators := len(op.products) - 1
	// nb operators available ** neededOperators
	possibility := int(math.Pow(float64(nops), float64(neededOperators)))

	l.Debug("CanSolve", "products", op.products, "result", op.result, "possibility", possibility)

	for i := 0; i < possibility; i++ {
		// convert i to a mask representing operator sequence
		// i=4  with 2 operators (convert 4 to base 2)  -> "11" -> ["1","1"]
		// i=7  with 3 operators (convert 7 to base 3)  -> "21" -> ["0","2","1"]
		mask := strconv.FormatInt(int64(i), nops)
		seq := strings.Split(mask, "")
		missing := neededOperators - len(seq)
		for i := 0; i < missing; i++ {
			seq = append([]string{"0"}, seq...)
		}

		l.Debug("try", "mask", mask, "operators sequence", seq)

		test := op.products[0]
		for j := 1; j < len(op.products); j++ {
			if seq[j-1] == ADD {
				test = test + op.products[j]
			} else if seq[j-1] == MUL {
				test = test * op.products[j]
			} else if seq[j-1] == CONCAT {
				test = aoc.ConcatInts(test, op.products[j])
			}

			// abort fastly -> next possibility
			if test > op.result {
				l.Debug("Abort result higher than expected", "test result", test)
				break
			}
		}
		if test == op.result {
			l.Debug("YEAH")
			return true
		}
	}
	return false
}

func processPart1(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	var total int = 0

	for _, op := range file.lines {
		ok := CanSolve(op, []string{ADD, MUL})
		if ok {
			total += op.result
		}
	}

	fmt.Println("part 1: ", int(total))
}

// single thread
// func processPart2(lines []string) {
// 	file := parseLines(lines)
// 	l.Debug("", "file", file)

// 	var total int = 0

// 	for _, op := range file.lines {
// 		ok := CanSolve(op, []string{ADD, MUL, CONCAT})
// 		if ok {
// 			total += op.result
// 		}
// 	}

// 	fmt.Println("part 2: ", int(total))
// }

func processPart2(lines []string) {
	file := parseLines(lines)
	l.Debug("", "file", file)

	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, op := range file.lines {
		wg.Add(1)
		go func(op Operation) {
			defer wg.Done()
			if CanSolve(op, []string{ADD, MUL, CONCAT}) {
				mu.Lock()
				total += op.result
				mu.Unlock()
			}
		}(op)
	}

	wg.Wait()

	fmt.Println("part 2: ", total)
}

func parseLines(lines []string) FileFormat {

	f := FileFormat{}

	for _, line := range lines {
		s := strings.Split(line, ":")
		op := Operation{}
		op.result = aoc.StringToInt(s[0])

		strProduct := strings.Fields(s[1])

		for _, v := range strProduct {
			op.products = append(op.products, aoc.StringToInt(v))
		}
		f.lines = append(f.lines, op)
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

	defer aoc.Timer("main")()

	switch config.Part {
	case 1:
		processPart1(lines)
	case 2:
		processPart2(lines)
	default:
		os.Exit(1)
	}
}
