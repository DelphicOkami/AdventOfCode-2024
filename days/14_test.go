package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FourteenSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunFourteenSuite(t *testing.T) {
	suite.Run(t, new(FourteenSuite))
}

func (suite *FourteenSuite) SetupTest() {
	suite.ProvidedInput = []string{}
}

func (suite *FourteenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayFourteenPart1(suite.ProvidedInput))
}

func (suite *FourteenSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayFourteenPart2(suite.ProvidedInput))
}
