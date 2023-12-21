package days_test

import (
	"testing"
	"aoc/days"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type D6Suite struct {
    suite.Suite
    ProvidedInput []string
	Races []days.D6Race
}

func TestRunD6Suite(t *testing.T) {
	suite.Run(t, new(D6Suite))
}

func (suite *D6Suite) SetupTest() {
	suite.ProvidedInput = []string{
		"Time:      7  15   30",
		"Distance:  9  40  200",
	}
	suite.Races = []days.D6Race{
		{
			Time: 7,
			TargetDistance: 9,
		},{
			Time: 15,
			TargetDistance: 40,
		},{
			Time: 30,
			TargetDistance: 200,
		},
	}
}

func (suite *D6Suite) TestRaceParsing() {
	assert.Equal(suite.T(), suite.Races, days.D6ParseRaces(suite.ProvidedInput))
}

func (suite *D6Suite) TestRaceWinOptions() {
	race0 := []int{2, 3, 4, 5}
	race1 := []int{4, 5, 6, 7, 8, 9, 10, 11}
	race2 := []int{11, 12, 13, 14, 15, 16, 17, 18, 19}
	assert.Equal(suite.T(), race0, suite.Races[0].GetWinOptions())
	assert.Equal(suite.T(), race1, suite.Races[1].GetWinOptions())
	assert.Equal(suite.T(), race2, suite.Races[2].GetWinOptions())
}

func (suite *D6Suite) TestD6P1CalculateTotal() {
	assert.Equal(suite.T(), 288, days.D6P1CalculateTotal(suite.Races))
}