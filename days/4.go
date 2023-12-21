package days

import (
	"strconv"
	"strings"
)

type D4card struct {
	CardNo int
	Winning []int
	Given []int
}

func DayFourPart1(input []string) int {
	gpSum := 0
	for _, cl := range input {
		card := d4ParseCard(cl)
		winners := card.GetWinningMatches()
		gpSum += calculateGeometricProgressionValue(len(winners))
	}

	return gpSum
}

func d4ParseCard(cardLine string) D4card {
	colonIndex := strings.Index(cardLine, ":")
	pipeIndex := strings.Index(cardLine, "|")
	cardNo, _ := strconv.Atoi(strings.Replace(strings.ToLower(cardLine[0:colonIndex]), "card ", "", 1))
	parsed := D4card{
		CardNo: cardNo,
		Winning: spaceSeparatedNumbersToIntSlice(cardLine[colonIndex+1:pipeIndex]),
		Given: spaceSeparatedNumbersToIntSlice(cardLine[pipeIndex+1:]),
	}
	return parsed
}

func spaceSeparatedNumbersToIntSlice(input string) []int {
	input = strings.Trim(input, " ")
	for hasDS := true; hasDS; hasDS = strings.Contains(input, "  ") {
		input = strings.ReplaceAll(input, "  ", " ")
	}
	numbers := make([]int, 0)
	for _, no := range strings.Split(input, " ") {
		atoiNo, err := strconv.Atoi(no)
		if err != nil {
			continue
		}
		numbers = append(numbers, atoiNo)
	}
	return numbers
}

func (c D4card) GetWinningMatches() []int {
	winners := make([]int, 0)
	for _, winner := range c.Winning {
		if intInSlice(winner, c.Given) {
			winners = append(winners, winner)
		}
	}
	return winners
}

func intInSlice(a int, list []int) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func calculateGeometricProgressionValue (term int) int {
	if term == 0 {
		return 0
	}
	out := 1
	for i := 1; i < term; i++ {
		out = out * 2
	}
	return out
}