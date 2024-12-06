package days

import (
	"slices"
)

var directions = map[rune][]int{'^':{0, -1},'>':{1, 0},'v':{0, 1}, '<':{-1, 0}}
func DaySixPart1 (input []string) int {
	coords := make([]Coords, 0)
	mp := ParseDay6Map(input) 

	for {
		if !slices.Contains(coords, mp.Guard.Position) {
			coords = append(coords, mp.Guard.Position)
		}
		err := mp.Move()
		if err != nil { break }
	}
	return len(coords)
}

func ParseDay6Map (input []string) Day6Map {
	grid := GetDay4(input)
	guard := Guard{}
	for x, column := range grid.Grid {
		for y, rune := range column {
			if rune == '^' {
				guard.Position.X = x
				guard.Position.Y = y
				guard.Facing = '^'
			}
		}
	}
	return Day6Map{
		Grid: grid,
		Guard: guard,
	}
}

type Coords struct {
	X int
	Y int
}

type Guard struct {
	Position Coords
	Facing rune
}

func (g *Guard) Turn() {
	switch g.Facing {
		case '^':
			g.Facing = '>'
		case '>':
			g.Facing = 'v'
		case 'v':
			g.Facing = '<'
		case '<':
			g.Facing = '^'
	}
}

type Day6Map struct{
	Grid Day4
	Guard Guard
}

func (d *Day6Map) Move() error {
	direction := directions[d.Guard.Facing]
	newX := d.Guard.Position.X + direction[1]
	newY := d.Guard.Position.Y + direction[0]
	r, err := d.Grid.GetRune(newX, newY)
	if err != nil {
		return err
	}
	if r == '#' {
		d.Guard.Turn()
		return nil
	}
	d.Guard.Position.X = newX
	d.Guard.Position.Y = newY
	return nil
}