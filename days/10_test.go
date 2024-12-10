package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TenSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunTenSuite(t *testing.T) {
	suite.Run(t, new(TenSuite))
}

func (suite *TenSuite) SetupTest() {
	suite.ProvidedInput = []string{"89010123", "78121874", "87430965", "96549874", "45678903", "32019012", "01329801", "10456732"}
}

func (suite *TenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 36, days.DayTenPart1(suite.ProvidedInput))
}

func (suite *TenSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayTenPart2(suite.ProvidedInput))
}

func (suite *TenSuite) TestRIConversion() {
	var i int
	var err error
	i, err = days.RuneToInt('0')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 0, i)
	i, err = days.RuneToInt('1')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 1, i)
	i, err = days.RuneToInt('2')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 2, i)
	i, err = days.RuneToInt('3')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 3, i)
	i, err = days.RuneToInt('4')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 4, i)
	i, err = days.RuneToInt('5')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 5, i)
	i, err = days.RuneToInt('6')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 6, i)
	i, err = days.RuneToInt('7')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 7, i)
	i, err = days.RuneToInt('8')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 8, i)
	i, err = days.RuneToInt('9')
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), 9, i)
	i, err = days.RuneToInt('a')
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), -1, i)

}

func (suite *TenSuite) TestGetIntDecreasingTrailHeads() {
	chizu := days.GetChizu(suite.ProvidedInput)
	peakCoords := days.Coords{Y: 6, X: 4}
	r, _ := chizu.GetRuneFromCoords(peakCoords)
	assert.Equal(suite.T(), '9', r)
	headsForPeak := []days.Coords{{Y: 5, X: 2}}
	assert.Equal(suite.T(), headsForPeak, chizu.GetIntDecreasingTrailHeadsFrom(peakCoords))
}

func (suite *TenSuite) TestGetIntIncreasingTrailHeads() {
	chizu := days.GetChizu(suite.ProvidedInput)
	headCoords := days.Coords{Y: 5, X: 2}
	r, _ := chizu.GetRuneFromCoords(headCoords)
	assert.Equal(suite.T(), '0', r)
	peaksForHead := []days.Coords{{Y: 6, X: 4}}
	assert.Equal(suite.T(), peaksForHead, chizu.GetIntIncreasingTrailHeadsFrom(headCoords))
	headCoords = days.Coords{Y: 4, X: 6}
	r, _ = chizu.GetRuneFromCoords(headCoords)
	assert.Equal(suite.T(), '0', r)
	peaksForHead = []days.Coords{{Y: 2, X: 5}, {Y: 4, X: 5}, {Y: 3, X: 4}}
	assert.Equal(suite.T(), peaksForHead, chizu.GetIntIncreasingTrailHeadsFrom(headCoords))
}

func (suite *TenSuite) TestGetPeaks() {
	chizu := days.GetChizu(suite.ProvidedInput)
	expected := []days.Coords{{Y: 0, X: 1}, {Y: 2, X: 5}, {Y: 3, X: 0}, {Y: 3, X: 4}, {Y: 4, X: 5}, {Y: 5, X: 4}, {Y: 6, X: 4}}
	instances := chizu.GetInstacesOf('9')
	for _, e := range expected {
		assert.Contains(suite.T(), instances, e)
	}
	assert.Len(suite.T(), instances, len(expected))
}
