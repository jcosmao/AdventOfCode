package main

import (
	aoc "aoc/lib"
	"container/heap"
	"fmt"
	l "log/slog"
	"os"
	"runtime/pprof"
	"strings"
)

type Day struct {
	config *aoc.Config
	grid   *aoc.Grid
	start  aoc.Position
	end    aoc.Position
	costs  map[aoc.Position]map[aoc.Direction]int
	val    int
}

type State struct {
	position  aoc.Position
	cost      int
	direction aoc.Direction
	index     int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func calculateCost(currentDirection, newDirection aoc.Direction) int {
	if currentDirection != newDirection {
		return 1000
	}
	return 0
}

func findShortestPath(day *Day) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &State{position: day.start, cost: 0, direction: aoc.E})

	day.costs = make(map[aoc.Position]map[aoc.Direction]int)
	day.costs[day.start] = make(map[aoc.Direction]int)
	day.costs[day.start][0] = 0

	directions := []aoc.Direction{aoc.N, aoc.E, aoc.S, aoc.W}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*State)
		l.Debug("Process", "node", current)

		if current.position == day.end {
			l.Info("costs", "", day.costs)
			return current.cost
		}

		for _, direction := range directions {
			if direction == current.direction.Opposite() {
				continue
			}

			newPos := day.grid.GetDirection(&current.position, direction)
			if newPos != nil && newPos != &current.position {
				if _, exists := day.grid.Node[*newPos]; exists {
					// wall
					continue
				}

				// cost
				newCost := current.cost + 1
				if current.direction != direction {
					newCost += 1000
				}

				if _, exists := day.costs[*newPos]; !exists {
					day.costs[*newPos] = make(map[aoc.Direction]int)
				}

				if _, exists := day.costs[*newPos][direction]; !exists || newCost < day.costs[*newPos][direction] {
					day.costs[*newPos][direction] = newCost
					heap.Push(pq, &State{position: *newPos, cost: newCost, direction: direction})
				}
			}
		}
	}

	return -1 // Si aucun chemin n'est trouvé
}

type nextCost struct {
	cost int
	dir  aoc.Direction
}

func tracePath(pos aoc.Position, day *Day, cost *int) {

	cursor := new(aoc.Cursor)
	cursor.Str = "O"

	day.grid.Node[pos] = make([]*aoc.Cursor, 1)
	day.grid.Node[pos][0] = cursor
	l.Debug("Mark position", "pos", pos)

	if pos == day.start {
		return
	}

	next := nextCost{
		cost: *cost,
		dir:  aoc.E,
	}

	togo := []nextCost{}

	directions := []aoc.Direction{aoc.N, aoc.E, aoc.S, aoc.W}

	for _, dir := range directions {
		if day.costs[pos][dir.Opposite()] <= *cost && day.costs[pos][dir.Opposite()] > 0 {
			next.dir = dir
			next.cost = *cost - (*cost - day.costs[pos][dir.Opposite()])

			l.Debug("Found next", "dir", dir, "next", next)
			togo = append(togo, next)
		}
	}

	for _, t := range togo {
		nextPos := day.grid.GetDirection(&pos, t.dir)

		if _, exists := day.grid.Node[*nextPos]; exists {
			l.Debug("Exists", "pos", nextPos, "v", day.grid.Node[*nextPos])
			continue
		}

		tracePath(*nextPos, day, &t.cost)
	}
}

func getResult(day *Day) int {
	total := 0
	shortestPathCost := findShortestPath(day)

	if day.config.Part == 1 {
		total = shortestPathCost

	} else if day.config.Part == 2 {
		tracePath(day.end, day, &shortestPathCost)
		for _, v := range day.grid.Node {
			if v != nil && v[0].Str == "O" {
				total += 1
			}
		}

		day.grid.Display()
	}

	return total
}

func parseLines(lines []string, day *Day) {

	day.grid = &aoc.Grid{
		Rows: len(lines),
		Node: make(map[aoc.Position][]*aoc.Cursor),
	}

	for i, line := range lines {
		sline := strings.Split(line, "")

		if i == 0 {
			day.grid.Cols = len(sline)
		}

		for j, char := range sline {
			position := aoc.Position{X: j, Y: i}

			if char == "." {
				continue
			} else if char == "S" {
				day.start = position
			} else if char == "E" {
				day.end = position
			} else {
				wall := &aoc.Cursor{Str: "#"}
				day.grid.Node[position] = []*aoc.Cursor{wall}
			}
		}
	}
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
	defer aoc.Timer("main")()

	lines := aoc.ReadFile(config.File)
	day := &Day{
		config: config,
	}

	parseLines(lines, day)
	l.Debug("Parsed input", "Day", day)

	res := getResult(day)
	fmt.Printf("part %d: %d\n", config.Part, res)
}
