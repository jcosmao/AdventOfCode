package lib

import (
	"fmt"
	l "log/slog"
)

type Position struct {
	X int
	Y int
}

type T interface {
	Repr() string
	GetValue() int
}

type Node struct {
	Value  int
	String string
	Edges  []*Node
}

type Grid struct {
	Rows int
	Cols int
	Nodes map[Position]T
	Node map[Position][]*Cursor
}

func NewGrid(rows int, cols int) *Grid {
	m := new(Grid)
	m.Cols = cols
	m.Rows = rows
	m.Node = map[Position][]*Cursor{}
	return m
}

func (g *Grid) String() string {
	return fmt.Sprintf("&Grid{Cols:%d, Rows:%d}", g.Cols, g.Rows)
}

type Cursor struct {
	Str   string
	Val   int
	Pos   *Position
	Grid   *Grid
	Track []Position
	Edges []*Cursor
}

func (c *Cursor) Repr() string {
	return c.Str
}

func (c *Cursor) GetValue() int {
	return c.Val
}

func (c *Cursor) String() string {
	return fmt.Sprintf("&Cursor{Str:%s, Val:%d, Pos:%s, Grid:%s, Track(len):%d, Edges(len):%d}", c.Str, c.Val, c.Pos, c.Grid, len(c.Track), len(c.Edges))
}

type Direction uint8

const (
	N Direction = iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

func (d Direction) Opposite() Direction {
	return (d + 4) % 8
}

func (p *Position) String() string {
	return fmt.Sprintf("[x:%d y:%d]", p.X, p.Y)
}

func (p *Position) ToString() string {
	return fmt.Sprintf("[x:%d y:%d]", p.X, p.Y)
}

func (m *Grid) GetFirstNode(p Position) *Cursor {
	if cursors, exists := m.Node[p]; exists {
		if len(cursors) > 1 {
			l.Debug("GetFirstCursor", "len", len(cursors))
		}

		if len(cursors) > 0 {
			return cursors[0]
		}
	}
	return nil
}

func (m *Grid) Display() {
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			pos := Position{X: j, Y: i}
			cursors, exists := m.Node[pos]
			if exists && len(cursors) > 0 {
				fmt.Print(cursors[0].Str)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println("")
	}
}

func (m *Grid) OutOfBound(p Position) bool {
	if p.X < 0 || p.X >= m.Cols {
		return true
	}

	if p.Y < 0 || p.Y >= m.Rows {
		return true
	}

	return false
}

func (c *Cursor) GetStart() *Position {
	if len(c.Track) > 0 {
		return &c.Track[0]
	}
	return nil
}

func (c *Cursor) GetDirection(d Direction) *Position {
	return c.Grid.GetDirection(c.Pos, d)
}

func (m *Grid) GetDirection(p *Position, d Direction) *Position {
	newPos := Position{}
	if d == N {
		newPos.X = p.X
		newPos.Y = p.Y - 1
	} else if d == S {
		newPos.X = p.X
		newPos.Y = p.Y + 1
	} else if d == E {
		newPos.X = p.X + 1
		newPos.Y = p.Y
	} else if d == W {
		newPos.X = p.X - 1
		newPos.Y = p.Y
	}

	if m.OutOfBound(newPos) {
		l.Debug("OutOfBound")
		return p
	}

	return &newPos
}

func (m *Grid) GetAllDirections(p *Position, dirs ...Direction) []*Position {
	if dirs == nil {
		dirs = []Direction{N, E, S, W}
	}
	allPos := []*Position{}
	for _, d := range dirs {
		newPos := m.GetDirection(p, d)
		if newPos != p {
			allPos = append(allPos, newPos)
		}
	}
	return allPos
}

func (c *Cursor) CanMove(d Direction) bool {
	newPos := c.Grid.GetDirection(c.Pos, d)
	if newPos == c.Pos {
		return false
	}
	return true
}


func (c *Cursor) Update(pos Position, target Position) bool {

	if _, exists := c.Grid.Node[pos]; exists {
		for i, cursor := range c.Grid.Node[pos] {
			if cursor == c {
				c.Grid.Node[pos] = append(c.Grid.Node[pos][:i], c.Grid.Node[pos][i+1:]...)
			}
		}
	}

	c.Pos = &target

	c.Grid.Node[target] = append(c.Grid.Node[target], c)

	return true
}



func (c *Cursor) GetNexts(d Direction) []*Position {

	res := []*Position{}
	pos := c.Pos
	for {
		newpos := c.Grid.GetDirection(pos, d)
		if pos == newpos {
			// we reached a border
			break
		}

		res = append(res, newpos)
		pos = newpos
	}

	return res
}

func (c *Cursor) Move(d Direction, recordTrack ...bool) bool {

	newPos := c.Grid.GetDirection(c.Pos, d)
	return c.MoveTo(newPos)
}

func (c *Cursor) MoveTo(pos *Position, recordTrack ...bool) bool {

	currentPos := c.Pos

	if pos == currentPos {
		return false
	}

	c.Update(*currentPos, *pos)

	if recordTrack == nil || recordTrack[0] {
		c.Track = append(c.Track, *pos)
	}

	return true
}



func (c *Cursor) MoveVector(v Position) {

	if _, exists := c.Grid.Node[*c.Pos]; exists {
		for i, cursor := range c.Grid.Node[*c.Pos] {
			if cursor == c {
				c.Grid.Node[*c.Pos] = append(c.Grid.Node[*c.Pos][:i], c.Grid.Node[*c.Pos][i+1:]...)
			}
		}
	}

	c.Pos.X += v.X
	c.Pos.Y += v.Y

	if c.Grid.OutOfBound(*c.Pos) {
		if c.Pos.X >= c.Grid.Cols {
			c.Pos.X = c.Pos.X - c.Grid.Cols
		}
		if c.Pos.Y >= c.Grid.Rows {
			c.Pos.Y = c.Pos.Y - c.Grid.Rows
		}
		if c.Pos.X < 0 {
			c.Pos.X = c.Grid.Cols + c.Pos.X
		}
		if c.Pos.Y < 0 {
			c.Pos.Y = c.Grid.Rows + c.Pos.Y
		}
	}

	c.Grid.Node[*c.Pos] = append(c.Grid.Node[*c.Pos], c)
}
