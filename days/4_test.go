package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestD41CasesPresented(t *testing.T) {
	card1 := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	}
	card2 := []string{
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	}
	card3 := []string{
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	}
	card4 := []string{
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	}
	card5 := []string{
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	}
	card6 := []string{
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	mass := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}	
	assert.Equal(t, 8, days.DayFourPart1(card1))
	assert.Equal(t, 2, days.DayFourPart1(card2))
	assert.Equal(t, 2, days.DayFourPart1(card3))
	assert.Equal(t, 1, days.DayFourPart1(card4))
	assert.Equal(t, 0, days.DayFourPart1(card5))
	assert.Equal(t, 0, days.DayFourPart1(card6))
	assert.Equal(t, 13, days.DayFourPart1(mass))
}

func TestD41CardParsing(t *testing.T) {
	card1 := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	parsed := days.D4ParseCard(card1)
	expectedCard := days.D4card{
		CardNo: 1,
		Winning: []int{
			41,
			48,
			83,
			86,
			17,
		},
		Given: []int{
			83,
			86,
			6,
			31,
			17,
			9,
			48,
			53,
		},
	}
	assert.Equal(t, expectedCard, parsed)
}

func TestSpaceSeparatedIntListToSlice(t *testing.T) {
	input := " 41 48 83 86 17 "
	expected := []int{
		41,
		48,
		83,
		86,
		17,
	}
	assert.Equal(t, expected, days.SpaceSeparatedNumbersToIntSlice(input))
	input = " 83 86  6 31 17  9 48 53"
	expected = []int{
		83,
		86,
		6,
		31,
		17,
		9,
		48,
		53,
	}
	assert.Equal(t, expected, days.SpaceSeparatedNumbersToIntSlice(input))
}

func TestD4FindWinners(t *testing.T) {
	input := days.D4card{
		CardNo: 0,
		Winning: []int{
			41,
			48,
			83,
			86,
			17,
		},
		Given: []int{
			83,
			86,
			6,
			31,
			17,
			9,
			48,
			53,
		},
	}
	
	expected := []int{
		48,
		83,
		86,
		17,
	}
	assert.Equal(t, expected, input.GetWinningMatches())
}

func TestD42CasePresented(t *testing.T) {
	cards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	}
	assert.Equal(t, 7, days.DayFourPart2(cards))
	cards = []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	assert.Equal(t, 30, days.DayFourPart2(cards))
}

func TestCalculateGrantedD4Copies(t *testing.T) {
	cards := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}
	expected := []int{
		1,
		2, 
		4,
		8,
		14,
		1,
	}
	assert.Equal(t, expected, days.CalculateGrantedD4Copies(cards))
}