package days

import (
	"fmt"
	"slices"
)

func DayEightPart1(input []string) int {
	grid := GetDay4(input)
	uniqueAntenodes := make([]Coords, 0)
	anteni := make(map[string][]Coords, 0)
	for i := 1; i <= 10; i++ {
		anteni[fmt.Sprint(i)] = make([]Coords, 0)
	}
	for i := 'A'; i <= 'Z'; i++ {
		anteni[string(i)] = make([]Coords, 0)
	}
	for i := 'a'; i <= 'z'; i++ {
		anteni[string(i)] = make([]Coords, 0)
	}
	for antenna, _ := range anteni {
		anteni[antenna] = grid.GetAntenodes(grid.GetCoordsFor([]rune(antenna)))
		for _, crd := range anteni[antenna] {
			if !slices.Contains(uniqueAntenodes, crd) {
				uniqueAntenodes = append(uniqueAntenodes, crd)
			}
		}
	}
	return len(uniqueAntenodes)
}

func DayEightPart2(input []string) int {
	return 0
}

func (g *Day4) GetCoordsFor(runes []rune) []Coords {
	out := make([]Coords, 0)
	for y, row := range g.Grid {
		for x, r := range row {
			if slices.Contains(runes, r) {
				out = append(out, Coords{X: x, Y: y})
			}
		}
	}
	return out
}

func (g *Day4) IsInGrid(c Coords) bool {
	_, err := g.GetRune(c.Y, c.X)
	return err == nil
}

func (g *Day4) GetAntenodes(c []Coords) []Coords {
	out := make([]Coords, 0)
	for pi, coordA := range c {
		for si, coordB := range c {
			if pi == si {
				continue
			}
			dX, dY := coordA.DistanceTo(coordB)
			ant := coordA.GetCoordAt(dX, dY)
			if g.IsInGrid(ant) {
				out = append(out, ant)
			}
		}
	}
	return out
}

func (c *Coords) DistanceTo(c2 Coords) (int, int) {
	return c.X - c2.X, c.Y - c2.Y
}

func (c *Coords) GetCoordAt(dx int, dy int) Coords {
	return Coords{X: c.X + dx, Y: c.Y + dy}
}
