package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SixSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunSixSuite(t *testing.T) {
	suite.Run(t, new(SixSuite))
}

func (suite *SixSuite) SetupTest() {
	suite.ProvidedInput = []string{"....#.....", ".........#", "..........", "..#.......", ".......#..", "..........", ".#..^.....", "........#.", "#.........", "......#..."}
}


func (suite *SixSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 41, days.DaySixPart1(suite.ProvidedInput))
}