package days

import (
	"regexp"
	"strconv"
)

func DayThirteenPart1(input []string) int {
	games := ThirteenParseGames(input, 0)
	total := 0
	for _,g  := range games {
		winnable, aPresses, bPresses := g.CalculatePresses()
		if winnable {
			total += (aPresses * 3) + (bPresses)
		}
	}
	return int(total)
}

func DayThirteenPart2(input []string) int {
	games := ThirteenParseGames(input, 10000000000000)
	total := 0
	for _,g  := range games {
		winnable, aPresses, bPresses := g.CalculatePresses()
		if winnable {
			total += (aPresses * 3) + (bPresses)
		}
	}
	return total
}

func GetCoordsFromStringParam(x, y string) Coords {
	xi, err := strconv.Atoi(x)
	if err != nil {
		xi = -1
	}
	yi, err := strconv.Atoi(y)
	if err != nil {
		yi = -1
	}
	return Coords{X: xi, Y: yi}
}

func ThirteenParseGames(input []string, prizeInflation int) []ThirteenGame {
	MatchAllReg := regexp.MustCompile(`.+: X[\+=](\d+), Y[\+=](\d+)`)
	out := make([]ThirteenGame, 0)
	for i := 2; i < len(input); i += 4 {
		ALine := input[i-2]
		BLine := input[i-1]
		PrizeLine := input[i]

		ac := MatchAllReg.FindStringSubmatch(ALine)
		bc := MatchAllReg.FindStringSubmatch(BLine)
		pc := MatchAllReg.FindStringSubmatch(PrizeLine)
		game := ThirteenGame{
			ButtonA:           GetCoordsFromStringParam(ac[1], ac[2]),
			ButtonB:           GetCoordsFromStringParam(bc[1], bc[2]),
			Prize:             GetCoordsFromStringParam(pc[1], pc[2]),
		}
		game.Prize.X = game.Prize.X + prizeInflation
		game.Prize.Y = game.Prize.Y + prizeInflation
		out = append(out, game)
	}
	return out
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

type ThirteenGame struct {
	ButtonA           Coords
	ButtonB           Coords
	Prize             Coords
}

func (g *ThirteenGame) CalculatePresses() (bool, int, int) {
	ay := g.ButtonA.Y * g.ButtonB.X
	py := g.Prize.Y * g.ButtonB.X

	ax := g.ButtonA.X * g.ButtonB.Y
	bx := g.ButtonB.X * g.ButtonB.Y
	px := g.Prize.X * g.ButtonB.Y

	xDelta := abs(ax - ay)
	pDelta := abs(px - py)

	if pDelta % xDelta != 0 {
		return false, -1, -1
	}

	aPresses := pDelta / xDelta

	if (px-aPresses*ax)%bx != 0 {
		return false, -1, -1
	}

	bPresses := (px - aPresses * ax) / bx

	return true, aPresses, bPresses
}

