package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ElevenSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunElevenSuite(t *testing.T) {
	suite.Run(t, new(ElevenSuite))
}

func (suite *ElevenSuite) SetupTest() {
	suite.ProvidedInput = []string{"125 17"}
}

func (suite *ElevenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 55312, days.DayElevenPart1(suite.ProvidedInput))
}

func (suite *ElevenSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 65601038650482, days.DayElevenPart2(suite.ProvidedInput))
}
