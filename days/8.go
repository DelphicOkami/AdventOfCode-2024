package days

import (
	"fmt"
	"slices"
)

func DayEightPart1(input []string) int {
	grid := GetDay4(input)
	uniqueAntenodes := make([]Coords, 0)
	anteni := make(map[string][]Coords, 0)
	for i := 0; i <= 9; i++ {
		anteni[fmt.Sprint(i)] = make([]Coords, 0)
	}
	for i := 'A'; i <= 'Z'; i++ {
		anteni[string(i)] = make([]Coords, 0)
	}
	for i := 'a'; i <= 'z'; i++ {
		anteni[string(i)] = make([]Coords, 0)
	}
	for antenna := range anteni {
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
	grid := GetDay4(input)
	uniqueAntenodes := make([]Coords, 0)
	anteni := make(map[string][]Coords, 0)
	for i := 0; i <= 9; i++ {
		anteni[fmt.Sprint(i)] = make([]Coords, 0)
	}
	for i := 'A'; i <= 'Z'; i++ {
		anteni[string(i)] = make([]Coords, 0)
	}
	for i := 'a'; i <= 'z'; i++ {
		anteni[string(i)] = make([]Coords, 0)
	}
	for antenna := range anteni {
		anteni[antenna] = grid.GetResonantAntenodes(grid.GetCoordsFor([]rune(antenna)))
		for _, crd := range anteni[antenna] {
			if !slices.Contains(uniqueAntenodes, crd) {
				uniqueAntenodes = append(uniqueAntenodes, crd)
			}
		}
	}
	// fmt.Println(anteni)
	return len(uniqueAntenodes)
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
				if !slices.Contains(out, ant) {out = append(out, ant)}
			}
		}
	}
	return out
}

func (g *Day4) GetResonantAntenodes(c []Coords) []Coords {
	out := make([]Coords, 0)
	if len(c) <= 1 {
		return out
	}
	for pi, coordA := range c {
		if !slices.Contains(out, coordA) {
			out = append(out, coordA)
		}
		for si, coordB := range c {
			if pi <= si {
				continue
			}
			dX, dY := coordA.GetDirection(coordB)
			coord := coordA.GetCoordAt(dX, dY)
			for ok := true; ok; ok = g.IsInGrid(coord) {
				if (!slices.Contains(out, coord) && g.IsInGrid(coord)) {
					out = append(out, coord)
				}
				coord = coord.GetCoordAt(dX, dY)
			}
			coord = coordA.GetCoordAt(dX *-1, dY *-1)
			for ok := true; ok; ok = g.IsInGrid(coord) {
				if (!slices.Contains(out, coord) && g.IsInGrid(coord)) {
					out = append(out, coord)
				}
				coord = coord.GetCoordAt(dX*-1, dY*-1)
			}
		}
	}
	// Sort and print out coords for debugging
	// slices.SortFunc(out, func(i, j Coords) int {
	// 	if i.Y == j.Y {
	// 		if i.X == j.X { return 0 }
	// 		if i.X > j.X { return 1 }
	// 		return -1
	// 	}
	// 	if i.X > j.X { return 1 }
	// 	return -1
	// })

	// for _, o := range out {
	// 	fmt.Printf("{X: %d, Y: %d}, ", o.X, o.Y)
	// }
	return out
}

func (c *Coords) DistanceTo(c2 Coords) (int, int) {
	return c.X - c2.X, c.Y - c2.Y
}

func (c *Coords) GetCoordAt(dx int, dy int) Coords {
	return Coords{X: c.X + dx, Y: c.Y + dy}
}

func (c *Coords) GetDirection(c2 Coords) (int, int) {
	dX, dY := c.DistanceTo(c2)
	cdX := dX
	cdY := dY
	if (dX < 0) {
		cdX = dX * - 1
	}
	if (dY < 0) {
		cdY = dY * - 1
	}
	XFactors := getFactors(cdX)
	YFactors := getFactors(cdY)
	biggestCommon := 1
	for _, f := range XFactors {
		if slices.Contains(YFactors, f) {
			biggestCommon = max(biggestCommon, f)
		}
	}
	return dX / biggestCommon, dY / biggestCommon
}

func getFactors(x int) []int {
	out := make([]int, 0)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			out = append(out, i)
		}
	 }
	 return out
}