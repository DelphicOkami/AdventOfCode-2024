package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type BlankSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunBlankSuite(t *testing.T) {
	suite.Run(t, new(BlankSuite))
}

func (suite *BlankSuite) SetupTest() {
	suite.ProvidedInput = []string{"190: 10 19", "3267: 81 40 27", "83: 17 5", "156: 15 6", "7290: 6 8 6 15", "161011: 16 10 13", "192: 17 8 14", "21037: 9 7 18 13", "292: 11 6 16 20"}
}

func (suite *BlankSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 14, days.DayBlankPart1(suite.ProvidedInput))
}