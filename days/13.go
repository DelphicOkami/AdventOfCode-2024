package days

import (
	"regexp"
	"strconv"
)

func DayThirteenPart1(input []string) int {
	games := ThirteenParseGames(input, 0)
	total := 0
	for _, g := range games {
		if g.IsWinnable(100) {
			win := g.GetCheapestWin(100)
			total += win.GetTotalCost()
		}
	}
	return total
}

func DayThirteenPart2(input []string) int {
	return 0
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
			possibleSolutions: make(map[int][]ThirteenSolution, 0),
		}
		game.Prize.X = game.Prize.X + prizeInflation
		game.Prize.Y = game.Prize.Y + prizeInflation
		out = append(out, game)
	}
	return out
}

type ThirteenGame struct {
	ButtonA           Coords
	ButtonB           Coords
	Prize             Coords
	possibleSolutions map[int][]ThirteenSolution
}

func (g *ThirteenGame) IsWinnable(maxPresses int) bool {
	mp, ok := g.possibleSolutions[maxPresses]
	if ok {
		return len(mp) > 0
	}
	g.possibleSolutions[maxPresses] = make([]ThirteenSolution, 0)
	maxAPresses := min(maxPresses, g.Prize.X/g.ButtonA.X, g.Prize.Y/g.ButtonA.Y)
	maxBPresses := min(maxPresses, g.Prize.X/g.ButtonB.X, g.Prize.Y/g.ButtonB.Y)
	for i := maxBPresses; i >= 0; i-- {
		bX := i * g.ButtonB.X
		bY := i * g.ButtonB.Y
		for j := 0; j <= maxAPresses; j++ {
			x := bX + (j * g.ButtonA.X)
			y := bY + (j * g.ButtonA.Y)
			if x == g.Prize.X && y == g.Prize.Y {
				g.possibleSolutions[maxPresses] = append(g.possibleSolutions[maxPresses], ThirteenSolution{APresses: j, BPresses: i})
				break
			}
			if x > g.Prize.X || y > g.Prize.Y {
				break
			}
		}
	}
	return len(g.possibleSolutions[maxPresses]) > 0
}

func (g *ThirteenGame) GetCheapestWin(maxPresses int) ThirteenSolution {
	if !g.IsWinnable(maxPresses) {
		return ThirteenSolution{}
	}
	var cheapest ThirteenSolution
	for i, s := range g.possibleSolutions[maxPresses] {
		if i == 0 {
			cheapest = s
			continue
		}
		if s.GetTotalCost() < cheapest.GetTotalCost() {
			cheapest = s
		}
	}
	return cheapest
}

type ThirteenSolution struct {
	APresses  int
	BPresses  int
	totalCost int
}

func (s *ThirteenSolution) GetTotalCost() int {
	if s.totalCost > 0 {
		return s.totalCost
	}

	s.totalCost = (s.APresses * 3) + (s.BPresses)
	return s.totalCost
}
