package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TwelveSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunTwelveSuite(t *testing.T) {
	suite.Run(t, new(TwelveSuite))
}

func (suite *TwelveSuite) SetupTest() {
	suite.ProvidedInput = []string{}
}

func (suite *TwelveSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayTwelvePart1(suite.ProvidedInput))
}

func (suite *TwelveSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayTwelvePart2(suite.ProvidedInput))
}
