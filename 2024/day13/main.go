package main

import (
	aoc "aoc/lib"
	"fmt"
	l "log/slog"
	"os"
	"regexp"
	"runtime/pprof"
)

// Day
type Day struct {
	Games []*Game
}

type Game struct {
	A     aoc.Position
	B     aoc.Position
	Prize aoc.Position
}

func Solve(g *Game, Acost float64, Bcost float64) (int, int) {
	// maxA := math.Min(float64(g.Prize.X/g.A.X), float64(g.Prize.Y/g.A.Y))
	// maxB := math.Min(float64(g.Prize.X/g.B.X), float64(g.Prize.Y/g.B.Y))

	// l.Debug("max", "A", int(maxA), "B", int(maxB))

	a := g.A
	b := g.B

	resA := (b.Y*g.Prize.X - b.X*g.Prize.Y) / (b.Y*g.A.X - b.X*a.Y)
	resB := (a.X*g.Prize.Y - a.Y*g.Prize.X) / (b.Y*a.X - b.X*a.Y)

	// C := g.A
	// maxC := int(maxA)
	// D := g.B
	// maxD := int(maxB)

	// if Acost > Bcost {
	// 	C = g.B
	// 	D = g.A
	// 	maxC = int(maxB)
	// 	maxD = int(maxA)
	// }

	// resC := 0
	// resD := 0

	// for c := maxC; c >= 0; c-- {
	// 	for d := 0; d <= maxD; d++ {
	// 		if (c*C.X+d*D.X == g.Prize.X) && (c*C.Y+d*D.Y == g.Prize.Y) {
	// 			resC = c
	// 			resD = d
	// 			goto end
	// 		}
	// 	}
	// }
	// end:
	// resA := resC
	// resB := resD

	// if Acost > Bcost {
	// 	resA = resD
	// 	resB = resC
	// }

	return resA, resB
}

func getResult(d *Day, part int) int {
	total := 0

	if part == 1 {
		for _, game := range d.Games {
			Acost := float64(game.Prize.X/game.A.X*3) + float64(game.Prize.X/game.A.Y*3)
			Bcost := float64(game.Prize.X/game.A.X) + float64(game.Prize.X/game.A.Y)

			a, b := Solve(game, Acost, Bcost)
			l.Debug("cost", "a", a, "b", b)
			if a <= 100 && b <= 100 {
				total += a*3 + b
			}
		}

	} else if part == 2 {
		add := 10000000000000
		for _, game := range d.Games {
			game.Prize.X += add
			game.Prize.Y += add

			Acost := float64(game.Prize.X/game.A.X*3) + float64(game.Prize.X/game.A.Y*3)
			Bcost := float64(game.Prize.X/game.A.X) + float64(game.Prize.X/game.A.Y)

			a, b := Solve(game, Acost, Bcost)
			l.Debug("cost", "a", a, "b", b)
			total += a*3 + b
		}

	}

	return total
}

func parseLines(lines []string) *Day {

	day := &Day{}

	r := regexp.MustCompile(`(.*): X(\+|=)(\d+), Y(\+|=)(\d+)`)

	game := new(Game)
	for _, line := range lines {
		f := r.FindStringSubmatch(line)
		l.Debug("Regex", "line", line, "match", f)
		if len(f) > 0 {
			x := aoc.StringToInt(f[3])
			y := aoc.StringToInt(f[5])

			if f[1] == "Button A" {
				game.A = aoc.Position{X: x, Y: y}
			} else if f[1] == "Button B" {
				game.B = aoc.Position{X: x, Y: y}
			} else if f[1] == "Prize" {
				game.Prize = aoc.Position{X: x, Y: y}
				day.Games = append(day.Games, game)
				game = new(Game)
			}
		}
	}

	return day
}

func main() {
	config := aoc.ParseFlags()

	if config.Profiling {
		_ = os.Remove("cpu.prof")
		cpuf, _ := os.Create("cpu.prof")
		pprof.StartCPUProfile(cpuf)
		defer cpuf.Close()
		defer pprof.StopCPUProfile()
		l.Info("[profiling] go tool pprof cpu.prof")
	}

	lines := aoc.ReadFile(config.File)
	d := parseLines(lines)
	l.Debug("Parsed input", "Day", d)

	res := getResult(d, config.Part)
	fmt.Printf("part %d: %d", config.Part, res)
}
