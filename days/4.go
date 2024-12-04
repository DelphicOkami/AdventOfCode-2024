package days

import "fmt"

func DayFourPart1 (input []string) int {
	d := GetDay4(input)
	return d.FindDirectionXmasCount(-1, 0) + d.FindDirectionXmasCount(-1, +1) + d.FindDirectionXmasCount(0, +1) + d.FindDirectionXmasCount(+1, +1) + d.FindDirectionXmasCount(+1, 0) + d.FindDirectionXmasCount(+1, -1) + d.FindDirectionXmasCount(0, -1) + d.FindDirectionXmasCount(-1, -1)
}
func DayFourPart2 (input []string) int {
	d := GetDay4(input)
	return d.FindXMasCount()
}

func GetDay4(input []string) Day4 {
	grid := make([][]rune, 0)
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	return Day4{Grid: grid}
}

type Day4 struct {
	Grid [][]rune
}

func (d *Day4) GetRune(x int, y int) (rune, error) {
	if (x < 0 || y < 0) {
		return '*', fmt.Errorf("out of bounds")
	}
	if (x >= len(d.Grid)) {
		return '*', fmt.Errorf("out of bounds")
	}
	row := d.Grid[x]
	if (y >= len(row)) {
		return '*', fmt.Errorf("out of bounds")
	}
	return row[y], nil
}
func (d *Day4) CoordMatches(x int, y int, r rune) bool {
	char, err := d.GetRune(x, y)
	if err != nil {
		return false
	}

	return char == r
}
func (d *Day4) FindDirectionXmasCount(xDir int, yDir int) int {
	count := 0
	var mx, ax, sx, my, ay, sy int
	for x, row := range d.Grid {
		for y, _ := range row {
			if (d.CoordMatches(x, y, 'X')) {
				mx = x + xDir
				my = y + yDir
				if (d.CoordMatches(mx, my, 'M')) {
					ax = mx + xDir
					ay = my + yDir
					if (d.CoordMatches(ax, ay, 'A')) {
						sx = ax + xDir
						sy = ay + yDir
						if (d.CoordMatches(sx, sy, 'S')) {
							count++
						}
					}
				}
			}
		}
	}
	return count
}
func (d *Day4) FindXMasCount() int {
	count := 0
	for x, row := range d.Grid {
		for y, _ := range row {
			if d.CoordMatches(x, y, 'A') {
				if ((d.CoordMatches(x + 1, y + 1, 'M') && d.CoordMatches(x - 1, y - 1, 'S')) || (d.CoordMatches(x + 1, y + 1, 'S') && d.CoordMatches(x - 1, y - 1, 'M'))) {
					if ((d.CoordMatches(x + 1, y - 1, 'M') && d.CoordMatches(x - 1, y + 1, 'S')) || (d.CoordMatches(x + 1, y - 1, 'S') && d.CoordMatches(x - 1, y + 1, 'M'))) {
						count ++
					}
				}
			}
		}
	}
	return count
}