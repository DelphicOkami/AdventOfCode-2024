package days_test

import (
	"aoc/days"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EightSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunEightSuite(t *testing.T) {
	suite.Run(t, new(EightSuite))
}

func (suite *EightSuite) SetupTest() {
	suite.ProvidedInput = []string{"............", "........0...", ".....0......", ".......0....", "....0.......", "......A.....", "............", "............", "........A...", ".........A..", "............", "............"}
}

func (suite *EightSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 14, days.DayEightPart1(suite.ProvidedInput))
}

func (suite *EightSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayEightPart2(suite.ProvidedInput))
}

func (suite *EightSuite) TestGetAntennaLocations() {
	crds0 := []days.Coords{{X: 8, Y: 1}, {X: 5, Y: 2}, {X: 7, Y: 3}, {X: 4, Y: 4}}
	crdsA := []days.Coords{{X: 6, Y: 5}, {X: 8, Y: 8}, {X: 9, Y: 9}}
	grid := days.GetDay4(suite.ProvidedInput)
	assert.Equal(suite.T(), crds0, grid.GetCoordsFor([]rune{'0'}))
	assert.Equal(suite.T(), crdsA, grid.GetCoordsFor([]rune{'A'}))
}

func (suite *EightSuite) TestGetDifference() {
	crds1 := days.Coords{X: 4, Y: 6}
	crds2 := days.Coords{X: 5, Y: 5}
	xCalc, yCalc := crds1.DistanceTo(crds2)
	assert.Equal(suite.T(), -1, xCalc)
	assert.Equal(suite.T(), 1, yCalc)
	xCalc, yCalc = crds2.DistanceTo(crds1)
	assert.Equal(suite.T(), 1, xCalc)
	assert.Equal(suite.T(), -1, yCalc)
}

func (suite *EightSuite) TestGetAntenodes() {
	crds0 := []days.Coords{{X: 8, Y: 1}, {X: 5, Y: 2}, {X: 7, Y: 3}, {X: 4, Y: 4}}
	crdsA := []days.Coords{{X: 6, Y: 5}, {X: 8, Y: 8}, {X: 9, Y: 9}}
	grid := days.GetDay4(suite.ProvidedInput)

	ant0 := []days.Coords{{X: 11, Y: 0}, {X: 2, Y: 3}, {X: 3, Y: 1}, {X: 6, Y: 0}, {X: 6, Y: 5}, {X: 9, Y: 4}, {X: 10, Y: 2}, {X: 0, Y: 7}, {X: 3, Y: 6}, {X: 1, Y: 5}}
	antA := []days.Coords{{X: 4, Y: 2}, {X: 3, Y: 1}, {X: 10, Y: 11}, {X: 7, Y: 7}, {X: 10, Y: 10}}
	calcAnt0 := grid.GetAntenodes(crds0)
	assert.Len(suite.T(), calcAnt0, len(ant0))
	for _, ant := range ant0 {
		assert.Contains(suite.T(), calcAnt0, ant)
	}
	calcAntA := grid.GetAntenodes(crdsA)
	assert.Len(suite.T(), calcAntA, len(antA))
	for _, ant := range antA {
		assert.Contains(suite.T(), calcAntA, ant)
	}
}
