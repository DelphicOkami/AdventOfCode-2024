package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type SevenSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunSevenSuite(t *testing.T) {
	suite.Run(t, new(SevenSuite))
}

func (suite *SevenSuite) SetupTest() {
	suite.ProvidedInput = []string{"190: 10 19", "3267: 81 40 27", "83: 17 5", "156: 15 6", "7290: 6 8 6 15", "161011: 16 10 13", "192: 17 8 14", "21037: 9 7 18 13", "292: 11 6 16 20"}
}

func (suite *SevenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 3749, days.DaySevenPart1(suite.ProvidedInput))
}
func (suite *SevenSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 11387, days.DaySevenPart2(suite.ProvidedInput))
}

func (suite *SevenSuite) TestParse() {
	parsedInput := days.ParseDaySevenInput(suite.ProvidedInput)
	one := days.DaySevenSum{
		Result: 190,
		Inputs: []int{10, 19},
	}
	assert.Equal(suite.T(), one, parsedInput[0])
}

func (suite *SevenSuite) TestGetPossibleOptions() {
	parsedInput := days.ParseDaySevenInput(suite.ProvidedInput)
	zeroPatterns := parsedInput[0].GetPossiblePatterns([]string{"*", "+"})
	assert.Contains(suite.T(), zeroPatterns, []string{"*"})
	assert.Contains(suite.T(), zeroPatterns, []string{"+"})
	assert.Len(suite.T(), zeroPatterns, 2)
	onePatterns := parsedInput[1].GetPossiblePatterns([]string{"*", "+"})
	assert.Contains(suite.T(), onePatterns, []string{"*", "*"})
	assert.Contains(suite.T(), onePatterns, []string{"+", "+"})
	assert.Contains(suite.T(), onePatterns, []string{"+", "*"})
	assert.Contains(suite.T(), onePatterns, []string{"*", "+"})
	assert.Len(suite.T(), onePatterns, 4)
}
func (suite *SevenSuite) TestGetWorkingOptions() {
	ops := []string{"+", "*"}
	parsedInput := days.ParseDaySevenInput(suite.ProvidedInput)
	zeroPatterns := parsedInput[0].GetWorkingPatterns(ops)
	assert.Contains(suite.T(), zeroPatterns, []string{"*"})
	assert.Len(suite.T(), zeroPatterns, 1)
	onePatterns := parsedInput[1].GetWorkingPatterns(ops)
	assert.Contains(suite.T(), onePatterns, []string{"+", "*"})
	assert.Contains(suite.T(), onePatterns, []string{"*", "+"})
	assert.Len(suite.T(), onePatterns, 2)
	eightPatterns := parsedInput[8].GetWorkingPatterns(ops)
	assert.Equal(suite.T(), [][]string{{"+", "*", "+"}}, eightPatterns)

	twoPatterns := parsedInput[2].GetWorkingPatterns(ops)
	assert.Len(suite.T(), twoPatterns, 0)

	assert.True(suite.T(), parsedInput[0].DoesWork(ops))
	assert.True(suite.T(), parsedInput[1].DoesWork(ops))
	assert.False(suite.T(), parsedInput[2].DoesWork(ops))
	assert.False(suite.T(), parsedInput[3].DoesWork(ops))
	assert.False(suite.T(), parsedInput[4].DoesWork(ops))
	assert.False(suite.T(), parsedInput[5].DoesWork(ops))
	assert.False(suite.T(), parsedInput[6].DoesWork(ops))
	assert.False(suite.T(), parsedInput[7].DoesWork(ops))
	assert.True(suite.T(), parsedInput[8].DoesWork(ops))
}

func (suite *SevenSuite) TestGetTwoWorkingOptions() {
	ops := []string{"*", "+", "||"}
	parsedInput := days.ParseDaySevenInput(suite.ProvidedInput)
	zeroPatterns := parsedInput[0].GetWorkingPatterns(ops)
	assert.Contains(suite.T(), zeroPatterns, []string{"*"})
	assert.Len(suite.T(), zeroPatterns, 1)
	onePatterns := parsedInput[1].GetWorkingPatterns(ops)
	assert.Contains(suite.T(), onePatterns, []string{"+", "*"})
	assert.Contains(suite.T(), onePatterns, []string{"*", "+"})
	assert.Len(suite.T(), onePatterns, 2)
	eightPatterns := parsedInput[8].GetWorkingPatterns(ops)
	assert.Equal(suite.T(), [][]string{{"+", "*", "+"}}, eightPatterns)

	twoPatterns := parsedInput[2].GetWorkingPatterns(ops)
	assert.Len(suite.T(), twoPatterns, 0)

	threePatterns := parsedInput[3].GetWorkingPatterns(ops)
	assert.Len(suite.T(), threePatterns, 1)
	fourPatterns := parsedInput[4].GetWorkingPatterns(ops)
	assert.Len(suite.T(), fourPatterns, 1)
	sixPatterns := parsedInput[6].GetWorkingPatterns(ops)
	assert.Len(suite.T(), sixPatterns, 1)

	assert.True(suite.T(), parsedInput[0].DoesWork(ops))
	assert.True(suite.T(), parsedInput[1].DoesWork(ops))
	assert.False(suite.T(), parsedInput[2].DoesWork(ops))
	assert.True(suite.T(), parsedInput[3].DoesWork(ops))
	assert.True(suite.T(), parsedInput[4].DoesWork(ops))
	assert.False(suite.T(), parsedInput[5].DoesWork(ops))
	assert.True(suite.T(), parsedInput[6].DoesWork(ops))
	assert.False(suite.T(), parsedInput[7].DoesWork(ops))
	assert.True(suite.T(), parsedInput[8].DoesWork(ops))

}
