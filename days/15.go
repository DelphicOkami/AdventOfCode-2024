package days

import (
	"fmt"
)

func DayFifteenPart1(input []string) int {
	w := FifteenParseInput(input)
	w.FollowMoveSequence()
	boxes := w.GetCoordsFor([]rune{'O'})
	boxTotal := 0
	for _, b := range boxes {
		boxTotal += (100 * b.Y) + b.X
	}
	return boxTotal
}

func DayFifteenPart2(input []string) int {
	w := FifteenParseInput(input)
	boxTotal := 0
	wt := w.EnThicken()
	wt.FollowMoveSequence()
	boxes := wt.GetCoordsFor([]rune{'['})

	for _, b := range boxes {
		boxTotal += (100 * b.Y) + b.X
	}
	return boxTotal
}

func FifteenParseInput(input []string) FifteenWarehouse {
	mapLines := make([]string, 0)
	dirSequence := make([]rune, 0)
	isMap := true
	for _, line := range input {
		if line == "" {
			isMap = false
			continue
		}
		if isMap {
			mapLines = append(mapLines, line)
		} else {
			dirSequence = append(dirSequence, []rune(line)...)
		}
	}

	chizu := GetChizu(mapLines)
	robotPos := chizu.GetCoordsFor([]rune{'@'})

	return FifteenWarehouse{
		Chizu:      chizu,
		Robot:      robotPos[0],
		Directions: dirSequence,
	}

}

type FifteenWarehouse struct {
	Chizu
	Robot      Coords
	Directions []rune
}

func (w *FifteenWarehouse) MoveRobot(X, Y int) {
	dest := w.Robot.GetCoordAt(X, Y)
	if w.Move(w.Robot, dest) {
		w.Robot = dest
	}
}

func (w *FifteenWarehouse) Move(start, end Coords) bool {
	s, _ := w.GetRuneFromCoords(start)
	e, _ := w.GetRuneFromCoords(end)
	if e == '.' {
		w.Chizu.Grid[end.Y][end.X] = s
		w.Chizu.Grid[start.Y][start.X] = e
		return true
	}
	if e == '#' {
		return false
	}
	dirX, dirY := end.GetDirection(start)
	chainEnd := end.GetCoordAt(dirX, dirY)
	if w.Move(end, chainEnd) {
		w.Chizu.Grid[end.Y][end.X] = s
		w.Chizu.Grid[start.Y][start.X] = '.'
		return true
	}
	return false
}

func (w *FifteenWarehouse) FollowMoveSequence() {
	for _, d := range w.Directions {
		switch d {
		case '^':
			w.MoveRobot(0, -1)
		case '>':
			w.MoveRobot(1, 0)
		case 'v':
			w.MoveRobot(0, 1)
		case '<':
			w.MoveRobot(-1, 0)
		}
	}
}

func (w *FifteenWarehouse) EnThicken() FifteenWarehouseThicc {
	var rob Coords
	tg := make([][]rune, 0)
	for y, row := range w.Grid {
		ty := make([]rune, 0)
		for x, r := range row {
			switch r {
			case '#':
				ty = append(ty, '#')
				ty = append(ty, '#')
			case '.':
				ty = append(ty, '.')
				ty = append(ty, '.')

			case 'O':
				ty = append(ty, '[')
				ty = append(ty, ']')

			case '@':
				ty = append(ty, '@')
				ty = append(ty, '.')
				rob.X = x * 2
				rob.Y = y
			}
		}
		tg = append(tg, ty)
	}

	return FifteenWarehouseThicc{Directions: w.Directions,
		Robot: rob,
		Chizu: Chizu{Grid: tg},
	}
}

func (c *Chizu) Render() {
	for _, row := range c.Grid {
		fmt.Println(string(row))
	}
}

type FifteenWarehouseThicc struct {
	Chizu
	Robot      Coords
	Directions []rune
}

func (w *FifteenWarehouseThicc) MoveRobot(X, Y int) {
	dest := w.Robot.GetCoordAt(X, Y)
	s, _ := w.GetRuneFromCoords(w.Robot)
	e, _ := w.GetRuneFromCoords(dest)
	if e == '.' {
		w.Chizu.Grid[dest.Y][dest.X] = s
		w.Chizu.Grid[w.Robot.Y][w.Robot.X] = '.'
		w.Robot = dest
		return
	}
	if e == '#' {
		return
	}
	if w.CanMoveBox(dest, X, Y) {
		w.MoveBox(dest, X, Y)
		w.Chizu.Grid[dest.Y][dest.X] = s
		w.Chizu.Grid[w.Robot.Y][w.Robot.X] = '.'
		w.Robot = dest
	}
}

func (w *FifteenWarehouseThicc) CanMoveBox(start Coords, X, Y int) bool {
	boxL, boxR := w.GetBox(start)
	boxLD := boxL.GetCoordAt(X, Y)
	boxRD := boxR.GetCoordAt(X, Y)
	boxLDR, _ := w.GetRuneFromCoords(boxLD)
	boxRDR, _ := w.GetRuneFromCoords(boxRD)
	if boxLDR == '#' || boxRDR == '#' {
		return false
	}
	boxLCan := boxLDR == '.' || boxLD == boxR
	boxRCan := boxRDR == '.' || boxRD == boxL
	if !boxLCan {
		boxLCan = w.CanMoveBox(boxLD, X, Y)
	}
	if !boxRCan {
		boxRCan = w.CanMoveBox(boxRD, X, Y)
	}

	return boxLCan && boxRCan
}

func (w *FifteenWarehouseThicc) MoveBox(start Coords, X, Y int) {
	boxL, boxR := w.GetBox(start)

	boxLR, _ := w.GetRuneFromCoords(boxL)
	boxRR, _ := w.GetRuneFromCoords(boxR)
	if boxLR == '.' || boxRR == '.' {
		//No longer a box
		return
	}

	boxLD := boxL.GetCoordAt(X, Y)
	boxRD := boxR.GetCoordAt(X, Y)
	boxLDR, _ := w.GetRuneFromCoords(boxLD)
	boxRDR, _ := w.GetRuneFromCoords(boxRD)
	if boxLDR == '#' || boxRDR == '#' {
		return
	}

	if boxLDR == '.' && boxRDR == '.' {
		w.Chizu.Grid[boxL.Y][boxL.X] = '.'
		w.Chizu.Grid[boxR.Y][boxR.X] = '.'
		w.Chizu.Grid[boxLD.Y][boxLD.X] = '['
		w.Chizu.Grid[boxRD.Y][boxRD.X] = ']'
		return
	}

	if (boxLDR == '[' || boxLDR == ']') && boxLD != boxR {
		w.MoveBox(boxLD, X, Y)
	}

	if (boxRDR == '[' || boxRDR == ']') && boxRD != boxL {
		w.MoveBox(boxRD, X, Y)
	}

	w.Chizu.Grid[boxL.Y][boxL.X] = '.'
	w.Chizu.Grid[boxR.Y][boxR.X] = '.'
	w.Chizu.Grid[boxLD.Y][boxLD.X] = '['
	w.Chizu.Grid[boxRD.Y][boxRD.X] = ']'
}

func (w *FifteenWarehouseThicc) GetBox(box Coords) (Coords, Coords) {
	b, _ := w.GetRuneFromCoords(box)
	if b == '[' {
		return box, box.GetCoordAt(1, 0)
	}
	return box.GetCoordAt(-1, 0), box
}

func (w *FifteenWarehouseThicc) FollowMoveSequence() {
	// fmt.Println("")
	// fmt.Println("SEQUENCE")
	for _, d := range w.Directions {
		switch d {
		case '^':
			w.MoveRobot(0, -1)
		case '>':
			w.MoveRobot(1, 0)
		case 'v':
			w.MoveRobot(0, 1)
		case '<':
			w.MoveRobot(-1, 0)
		}
		// fmt.Print(string(d))/
		// w.DetectBrokenBox()
	}
}

func (w *FifteenWarehouseThicc) DetectBrokenBox() {
	prev := '.'
	for _, row := range w.Chizu.Grid {
		for _, c := range row {
			if c == '[' || c == ']' {
				if prev == c {
					fmt.Println("")
					w.Render()
					panic("broken box detected")
				}
			}
			prev = c
		}
	}
}
