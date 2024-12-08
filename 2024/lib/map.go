package lib

import (
	"fmt"
	"log/slog"
)

type Position struct {
	X int
	Y int
}

type T interface {
	Display() string
}

type Map struct {
	Rows  int
	Cols  int
	Nodes map[Position]T
}

type Cursor struct {
	Pos *Position
	Map   *Map
	Track []*Position
}

type Direction uint8

const (
	Up    Direction = iota
	Down
	Right
	Left
)

func (m *Map) Display() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			pos := Position{X: j, Y: i}
			if obj, exists := m.Nodes[pos]; exists {
				fmt.Print(obj.Display())
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

func (m *Map) OutOfBound(p Position) bool {
	if p.X < 0 || p.X >= m.Cols {
		return true
	}

	if p.Y < 0 || p.Y >= m.Rows {
		return true
	}

	return false
}

func (c *Cursor) Move(d Direction, recordTrack ...bool) bool {

	newPos := Position{}
	if d == Up {
		newPos.X = c.Pos.X
		newPos.Y = c.Pos.Y - 1
	} else if d == Down {
		newPos.X = c.Pos.X
		newPos.Y = c.Pos.Y + 1
	} else if d == Right {
		newPos.X = c.Pos.X + 1
		newPos.Y = c.Pos.Y
	} else if d == Left {
		newPos.X = c.Pos.X - 1
		newPos.Y = c.Pos.Y
	}

	slog.Debug("Move", "new", newPos)

	if !c.Map.OutOfBound(newPos) {
		c.Pos = &newPos

		if recordTrack == nil || recordTrack[0] {
			c.Track = append(c.Track, &newPos)
		}

		return true
	}

	return false
}
