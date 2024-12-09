package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NineSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunNineSuite(t *testing.T) {
	suite.Run(t, new(NineSuite))
}

func (suite *NineSuite) SetupTest() {
	suite.ProvidedInput = []string{}
}

func (suite *NineSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayNinePart1(suite.ProvidedInput))
}

func (suite *NineSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayNinePart2(suite.ProvidedInput))
}
