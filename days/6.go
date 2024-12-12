package days

import (
	"slices"
)

var directions = map[rune][]int{'^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0}}

func DaySixPart1(input []string) int {
	coords := make([]Coords, 0)
	mp := ParseDay6Map(input)

	for {
		if !slices.Contains(coords, mp.Guard.Position) {
			coords = append(coords, mp.Guard.Position)
		}
		err := mp.Move()
		if err != nil {
			break
		}
	}
	return len(coords)
}
func DaySixPart2(input []string) int {
	coords := make([]Coords, 0)
	mp := ParseDay6Map(input)
	for y, row := range mp.Grid.Grid {
		for x, _ := range row {
			if mp.Grid.Grid[y][x] == '#' {
				continue
			}
			thisMap := mp.Copy()
			thisMap.Grid.Grid[y][x] = '#'
			if thisMap.DetectLoop() {
				coords = append(coords, Coords{X: x, Y: y})
			}
			thisMap = Day6Map{}
		}
	}
	return len(coords)
}

func ParseDay6Map(input []string) Day6Map {
	grid := GetChizu(input)
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
		Grid:  grid,
		Guard: guard,
	}
}

type Coords struct {
	X int
	Y int
}

type Guard struct {
	Position Coords
	Facing   rune
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

type Day6Map struct {
	Grid  Chizu
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

func (d *Day6Map) DetectLoop() bool {
	coords := make(map[Coords]rune, 0)
	for {

		val, ok := coords[d.Guard.Position]
		if !ok {
			coords[d.Guard.Position] = d.Guard.Facing
		} else if val == d.Guard.Facing {
			return true
		}
		err := d.Move()
		if err != nil {
			return false
		}
	}
}

func (d *Day6Map) Copy() Day6Map {
	grid := make([][]rune, 0)
	for _, row := range d.Grid.Grid {
		r := make([]rune, 0)
		r = append(r, row...)
		grid = append(grid, r)
	}
	g := Guard{
		Position: Coords{X: d.Guard.Position.X, Y: d.Guard.Position.Y},
		Facing:   d.Guard.Facing,
	}
	return Day6Map{
		Grid:  Chizu{Grid: grid},
		Guard: g,
	}
}
