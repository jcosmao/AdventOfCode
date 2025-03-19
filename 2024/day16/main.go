// package main

// import (
// 	aoc "aoc/lib"
// 	"fmt"
// 	l "log/slog"
// 	"os"
// 	"runtime/pprof"
// 	"slices"
// 	"strings"

// 	"math/rand"
// )

// // Day
// type Day struct {
// 	config      *aoc.Config
// 	grid        *aoc.Grid
// 	start       aoc.Position
// 	end         aoc.Position
// 	paths       []*aoc.Cursor
// 	best        int
// 	nodeMinCost map[aoc.Position]int
// }

// const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// const (
// 	letterIdxBits = 6                    // 6 bits to represent a letter index
// 	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
// 	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
// )

// func RandStringBytesMaskImpr(n int) string {
// 	b := make([]byte, n)
// 	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
// 	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
// 		if remain == 0 {
// 			cache, remain = rand.Int63(), letterIdxMax
// 		}
// 		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
// 			b[i] = letterBytes[idx]
// 			i--
// 		}
// 		cache >>= letterIdxBits
// 		remain--
// 	}

// 	return string(b)
// }

// func RunRunRun(day *Day, c *aoc.Cursor, dir aoc.Direction) {

// 	if *c.Pos == day.end {
// 		day.paths = append(day.paths, c)
// 		if day.best == 0 || c.Val < day.best {
// 			day.best = c.Val
// 		}
// 		l.Info("End", "score", day.best)

// 		return
// 	}

// 	if day.best != 0 && c.Val > day.best {
// 		return
// 	}

// 	if _, exists := day.nodeMinCost[*c.Pos]; exists {
// 		if day.nodeMinCost[*c.Pos] > 0 && c.Val > day.nodeMinCost[*c.Pos] {
// 			return
// 		}
// 	} else {
// 		day.nodeMinCost[*c.Pos] = c.Val
// 		l.Info("NodeMinCost update", "pos", c.Pos, "val", c.Val)
// 	}

// 	l.Debug("HERE", "direction", dir, "cursor", c)

// 	for _, dirNext := range []aoc.Direction{aoc.N, aoc.E, aoc.S, aoc.W} {

// 		if dirNext == dir.Opposite() {
// 			continue
// 		}

// 		posNext := c.GetDirection(dirNext)
// 		if c.Pos == posNext {
// 			l.Debug("Cannot move", "from", c.Pos, "to", posNext)
// 			continue
// 		}

// 		if slices.Contains(c.Track, *posNext) {
// 			l.Debug("Already visited", "pos", posNext, "cursor", c)
// 			continue
// 		}

// 		nodeNext := day.grid.GetFirstNode(*posNext)

// 		if nodeNext == nil || nodeNext.Str != "#" {

// 			csplit := *c
// 			csplit.Str = RandStringBytesMaskImpr(1)

// 			c.MoveTo(posNext)

// 			if dirNext != dir {
// 				c.Val += 1000
// 			}

// 			c.Val += 1

// 			RunRunRun(day, c, dirNext)

// 			c = &csplit
// 		}
// 	}
// }

// func getResult(day *Day) int {
// 	total := 0

// 	cursor := aoc.Cursor{
// 		Str:  "X",
// 		Pos:  &day.start,
// 		Grid: day.grid,
// 		Val:  0,
// 	}

// 	startDir := aoc.E

// 	RunRunRun(day, &cursor, startDir)

// 	l.Debug("Cursor", "", &cursor)
// 	l.Debug("paths", "", day.paths)

// 	if day.config.Part == 1 {
// 		total = day.best

// 		// for _, cursors := range day.paths {

// 		// 	if len(cursors.Track) == 36 && cursors.Val < 10000 {

// 		// 		l.Debug("Path", "", cursors.Track)
// 		// 		l.Debug("Path", "", len(cursors.Track))
// 		// 		l.Debug("Cost", "", cursors.Val)

// 		// 		for _, step := range cursors.Track {
// 		// 			day.grid.Node[step] = []*aoc.Cursor{&aoc.Cursor{Str: "X"}}
// 		// 			// time.Sleep(time.Second)
// 		// 			// day.grid.Display()
// 		// 		}
// 		// 	}
// 		// }

// 	} else if day.config.Part == 2 {
// 	}

// 	return total
// }

// func parseLines(lines []string, day *Day) {

// 	day.grid = &aoc.Grid{
// 		Rows: len(lines),
// 		Node: make(map[aoc.Position][]*aoc.Cursor),
// 	}

// 	for i, line := range lines {
// 		sline := strings.Split(line, "")

// 		if i == 0 {
// 			day.grid.Cols = len(sline)
// 		}

// 		for j, char := range sline {
// 			position := aoc.Position{X: j, Y: i}

// 			if char == "." {
// 				continue
// 			} else if char == "S" {
// 				day.start = position
// 			} else if char == "E" {
// 				day.end = position
// 			} else {
// 				wall := &aoc.Cursor{Str: "#"}
// 				day.grid.Node[position] = []*aoc.Cursor{wall}
// 			}
// 		}
// 	}
// }

// func main() {
// 	config := aoc.ParseFlags()

// 	if config.Profiling {
// 		_ = os.Remove("cpu.prof")
// 		cpuf, _ := os.Create("cpu.prof")
// 		pprof.StartCPUProfile(cpuf)
// 		defer cpuf.Close()
// 		defer pprof.StopCPUProfile()
// 		l.Info("[profiling] go tool pprof cpu.prof")
// 	}
// 	defer aoc.Timer("main")()

// 	lines := aoc.ReadFile(config.File)
// 	day := &Day{
// 		config: config,
// 		nodeMinCost: make(map[aoc.Position]int),
// 	}

// 	parseLines(lines, day)
// 	l.Debug("Parsed input", "Day", day)

// 	res := getResult(day)
// 	fmt.Printf("part %d: %d\n", config.Part, res)
// }
