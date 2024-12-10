package days

import (
	"fmt"
	"slices"
)

func DayTenPart1(input []string) int {
	m := GetChizu(input)
	peaks := make(map[Coords][]Coords, 0)
	heads := m.GetInstacesOf('0')
	for _, p := range heads {
		peaks[p] = m.GetIntIncreasingTrailHeadsFrom(p, true)
	}
	sum := 0
	// fmt.Println(peaks)
	for _, p := range peaks {
		sum += len(p)
	}
	return sum
}

func DayTenPart2(input []string) int {
	m := GetChizu(input)
	peaks := make(map[Coords][]Coords, 0)
	heads := m.GetInstacesOf('0')
	for _, p := range heads {
		peaks[p] = m.GetIntIncreasingTrailHeadsFrom(p, false)
	}
	sum := 0
	// fmt.Println(peaks)
	for _, p := range peaks {
		sum += len(p)
	}
	return sum
}

func RuneToInt(r rune) (int, error) {
	if int(r) < int('0') || int(r) > int('9') {
		return -1, fmt.Errorf("rune is out of 0-9 range")
	}
	return int(r - '0'), nil
}

func (m *Chizu) GetRuneFromCoords(c Coords) (rune, error) {
	return m.GetRune(c.Y, c.X)
}

func (c *Chizu) GetIntDecreasingTrailHeadsFrom(s Coords, unique bool) []Coords {
	currentRune, err := c.GetRuneFromCoords(s)
	if err != nil {
		return []Coords{}
	}
	if currentRune == '0' {
		return []Coords{s}
	}
	currentVal, _ := RuneToInt(currentRune)
	nextStep := currentVal - 1
	nsr := rune(fmt.Sprint(nextStep)[0])
	return c.findNextStepRecursion(s, nsr, unique, c.GetIntDecreasingTrailHeadsFrom)
}

func (c *Chizu) GetIntIncreasingTrailHeadsFrom(s Coords, unique bool) []Coords {
	currentRune, err := c.GetRuneFromCoords(s)
	if err != nil {
		return []Coords{}
	}
	if currentRune == '9' {
		return []Coords{s}
	}
	currentVal, _ := RuneToInt(currentRune)
	nextStep := currentVal + 1
	nsr := rune(fmt.Sprint(nextStep)[0])
	return c.findNextStepRecursion(s, nsr, unique, c.GetIntIncreasingTrailHeadsFrom)
}

func (c *Chizu) findNextStepRecursion(start Coords, search rune, unique bool, recur func(Coords, bool) []Coords) []Coords {

	heads := make([]Coords, 0)
	var err error
	var r rune
	var targ Coords
	targ = Coords{X: start.X, Y: start.Y - 1}
	r, err = c.GetRuneFromCoords(targ)
	if err == nil {
		if r == search {
			heads = append(heads, recur(targ, unique)...)
		}
	}
	targ = Coords{X: start.X + 1, Y: start.Y}
	r, err = c.GetRuneFromCoords(targ)
	if err == nil {
		if r == search {
			heads = append(heads, recur(targ, unique)...)
		}
	}
	targ = Coords{X: start.X, Y: start.Y + 1}
	r, err = c.GetRuneFromCoords(targ)
	if err == nil {
		if r == search {
			heads = append(heads, recur(targ, unique)...)
		}
	}
	targ = Coords{X: start.X - 1, Y: start.Y}
	r, err = c.GetRuneFromCoords(targ)
	if err == nil {
		if r == search {
			heads = append(heads, recur(targ, unique)...)
		}
	}
	if !unique {
		return heads
	}
	uniqueHeads := make([]Coords, 0)
	for _, h := range heads {
		if !slices.Contains(uniqueHeads, h) {
			uniqueHeads = append(uniqueHeads, h)
		}
	}

	return uniqueHeads
}

func (c *Chizu) GetInstacesOf(r rune) []Coords {
	out := make([]Coords, 0)
	for y, row := range c.Grid {
		for x, a := range row {
			if a == r {
				out = append(out, Coords{X: x, Y: y})
			}
		}
	}
	return out
}
