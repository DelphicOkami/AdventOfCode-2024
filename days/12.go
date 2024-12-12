package days

import (
	"slices"
)

func DayTwelvePart1(input []string) int {
	sum := 0
	g := TweleveGetGarden(input)
	for _, p := range g.GetAllPlots() {
		sum += p.Perimiter * len(p.Spots)
	}
	return sum
}

func DayTwelvePart2(input []string) int {
	return 0
}

func TweleveGetGarden(input []string) Garden {
	return Garden{
		GetChizu(input),
		make([]TwelvePlot, 0),
		make([]Coords, 0),
	}
}

func TwelveGetPerimiter(c []Coords) int {
	var maxX, maxY, minX, minY int
	for i, tc := range c {
		if i == 0 {
			maxX = tc.X
			minX = tc.X
			maxY = tc.Y
			minY = tc.Y
			continue
		}
		if maxX < tc.X {
			maxX = tc.X
		}
		if maxY < tc.Y {
			maxY = tc.Y
		}
		if minX > tc.X {
			minX = tc.X
		}
		if minY > tc.Y {
			minY = tc.Y
		}
	}

	return (maxX - minX + 1) + (maxY-minY+1)*2
}

type Garden struct {
	Layout   Chizu
	plots    []TwelvePlot
	Assigned []Coords
}

func (g *Garden) GetAllPlots() []TwelvePlot {
	if len(g.plots) == 0 {
		for y, row := range g.Layout.Grid {
			for x := range row {
				coord := Coords{X: x, Y: y}
				if !slices.Contains(g.Assigned, coord) {
					plot := g.GetPlotFrom(coord)
					g.plots = append(g.plots, plot)
					g.Assigned = append(g.Assigned, plot.Spots...)

				}
			}
		}
	}
	return g.plots
}

func (g *Garden) GetPlotFrom(c Coords) TwelvePlot {
	r, _ := g.Layout.GetRuneFromCoords(c)
	previous := make([]Coords, 0)
	plot, _, walls := g.consider(r, c, previous, 0)
	return TwelvePlot{plot, walls}
}

func (g *Garden) consider(r rune, c Coords, previous []Coords, walls int) ([]Coords, []Coords, int) {
	plot := make([]Coords, 0)
	if !g.Layout.CoordMatches(c.Y, c.X, r) {
		walls = walls + 1
	}
	if slices.Contains(previous, c) {
		return plot, previous, walls
	}
	previous = append(previous, c)
	if !g.Layout.CoordMatches(c.Y, c.X, r) {

		return plot, previous, walls
	}
	plot = append(plot, c)
	north := c.GetCoordAt(0, -1)
	east := c.GetCoordAt(1, 0)
	south := c.GetCoordAt(0, 1)
	west := c.GetCoordAt(-1, 0)
	pl, previous, nwalls := g.consider(r, north, previous, 0)
	plot = append(plot, pl...)
	pl, previous, ewalls := g.consider(r, east, previous, 0)
	plot = append(plot, pl...)
	pl, previous, swalls := g.consider(r, south, previous, 0)
	plot = append(plot, pl...)
	pl, previous, wwalls := g.consider(r, west, previous, 0)
	plot = append(plot, pl...)
	return plot, previous, (walls + nwalls + ewalls + swalls + wwalls)
}

type TwelvePlot struct {
	Spots     []Coords
	Perimiter int
}
