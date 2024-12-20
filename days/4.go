package days

import "fmt"

func DayFourPart1(input []string) int {
	d := GetChizu(input)
	return FindDirectionXmasCount(-1, 0, d) + FindDirectionXmasCount(-1, +1, d) + FindDirectionXmasCount(0, +1, d) + FindDirectionXmasCount(+1, +1, d) + FindDirectionXmasCount(+1, 0, d) + FindDirectionXmasCount(+1, -1, d) + FindDirectionXmasCount(0, -1, d) + FindDirectionXmasCount(-1, -1, d)
}
func DayFourPart2(input []string) int {
	d := GetChizu(input)
	return FindXMasCount(d)
}

func GetChizu(input []string) Chizu {
	grid := make([][]rune, 0)
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	return Chizu{Grid: grid}
}

type Chizu struct {
	Grid [][]rune
	maxX int
	maxY int
}

func (d *Chizu) GetRune(y int, x int) (rune, error) {
	if x < 0 || y < 0 || y > d.GetMaxY() || x > d.GetMaxX() {
		return '*', fmt.Errorf("out of bounds")
	}
	row := d.Grid[y]
	if x >= len(row) {
		return '*', fmt.Errorf("out of bounds")
	}
	return row[x], nil
}

func (g *Chizu) GetMaxX() int {
	if g.maxX == 0 {
		for _, row := range g.Grid {
			g.maxX = max(g.maxX, len(row)-1)
		}
	}
	return g.maxX
}

func (g *Chizu) GetMaxY() int {
	if g.maxY == 0 {
		g.maxY = len(g.Grid) - 1
	}
	return g.maxY
}

func (d *Chizu) CoordMatches(x int, y int, r rune) bool {
	char, err := d.GetRune(x, y)
	if err != nil {
		return false
	}

	return char == r
}
func FindDirectionXmasCount(xDir int, yDir int, d Chizu) int {
	count := 0
	var mx, ax, sx, my, ay, sy int
	for x, row := range d.Grid {
		for y := range row {
			if d.CoordMatches(x, y, 'X') {
				mx = x + xDir
				my = y + yDir
				if d.CoordMatches(mx, my, 'M') {
					ax = mx + xDir
					ay = my + yDir
					if d.CoordMatches(ax, ay, 'A') {
						sx = ax + xDir
						sy = ay + yDir
						if d.CoordMatches(sx, sy, 'S') {
							count++
						}
					}
				}
			}
		}
	}
	return count
}
func FindXMasCount(d Chizu) int {
	count := 0
	for x, row := range d.Grid {
		for y, _ := range row {
			if d.CoordMatches(x, y, 'A') {
				if (d.CoordMatches(x+1, y+1, 'M') && d.CoordMatches(x-1, y-1, 'S')) || (d.CoordMatches(x+1, y+1, 'S') && d.CoordMatches(x-1, y-1, 'M')) {
					if (d.CoordMatches(x+1, y-1, 'M') && d.CoordMatches(x-1, y+1, 'S')) || (d.CoordMatches(x+1, y-1, 'S') && d.CoordMatches(x-1, y+1, 'M')) {
						count++
					}
				}
			}
		}
	}
	return count
}
