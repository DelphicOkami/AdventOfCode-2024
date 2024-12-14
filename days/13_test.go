package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ThirteenSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunThirteenSuite(t *testing.T) {
	suite.Run(t, new(ThirteenSuite))
}

func (suite *ThirteenSuite) SetupTest() {
	suite.ProvidedInput = []string{"Button A: X+94, Y+34", "Button B: X+22, Y+67", "Prize: X=8400, Y=5400", "", "Button A: X+26, Y+66", "Button B: X+67, Y+21", "Prize: X=12748, Y=12176", "", "Button A: X+17, Y+86", "Button B: X+84, Y+37", "Prize: X=7870, Y=6450", "", "Button A: X+69, Y+23", "Button B: X+27, Y+71", "Prize: X=18641, Y=10279"}
}

func (suite *ThirteenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 480, days.DayThirteenPart1(suite.ProvidedInput))
}

func (suite *ThirteenSuite) TestThirteenParsing() {
	thirteenGames := days.ThirteenParseGames(suite.ProvidedInput, 0)
	assert.Equal(suite.T(), 94, thirteenGames[0].ButtonA.X)
	assert.Equal(suite.T(), 34, thirteenGames[0].ButtonA.Y)
	assert.Equal(suite.T(), 22, thirteenGames[0].ButtonB.X)
	assert.Equal(suite.T(), 67, thirteenGames[0].ButtonB.Y)
	assert.Equal(suite.T(), 8400, thirteenGames[0].Prize.X)
	assert.Equal(suite.T(), 5400, thirteenGames[0].Prize.Y)
}

func (suite *ThirteenSuite) TestThirteenWinability() {
	thirteenGames := days.ThirteenParseGames(suite.ProvidedInput, 0)
	winnable, _, _ := thirteenGames[0].CalculatePresses()
	assert.True(suite.T(), winnable)

	winnable, _, _ = thirteenGames[1].CalculatePresses()
	assert.False(suite.T(), winnable)
	winnable, _, _ = thirteenGames[2].CalculatePresses()
	assert.True(suite.T(), winnable)
	winnable, _, _ = thirteenGames[3].CalculatePresses()
	assert.False(suite.T(), winnable)
}
