package days

import "fmt"

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
	return 0
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

func (c *Chizu) render() {
	for _, row := range c.Grid {
		fmt.Println(string(row))
	}
}
