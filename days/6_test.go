package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
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
func (suite *SixSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 6, days.DaySixPart2(suite.ProvidedInput))
}

func (suite *SixSuite) TestLoopDetection() {
	d6 := days.ParseDay6Map(suite.ProvidedInput)
	d61 := d6.Copy()
	assert.False(suite.T(), d61.DetectLoop(), "Found a loop where there isn't")
	d62 := d6.Copy()
	d62.Grid.Grid[6][3] = '#'
	assert.True(suite.T(), d62.DetectLoop(), "Found no loop where there is")
	d63 := d6.Copy()
	d63.Grid.Grid[4][7] = '#'
	assert.False(suite.T(), d63.DetectLoop(), "Found a loop where there isn't")
}
