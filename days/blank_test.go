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
	suite.ProvidedInput = []string{}
}

func (suite *BlankSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayBlankPart1(suite.ProvidedInput))
}

func (suite *BlankSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayBlankPart2(suite.ProvidedInput))
}
