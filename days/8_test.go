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
	CoordsA       []days.Coords
	Coords0       []days.Coords
}

func TestRunEightSuite(t *testing.T) {
	suite.Run(t, new(EightSuite))
}

func (suite *EightSuite) SetupTest() {
	suite.ProvidedInput = []string{"............", "........0...", ".....0......", ".......0....", "....0.......", "......A.....", "............", "............", "........A...", ".........A..", "............", "............"}
	suite.Coords0 = []days.Coords{{X: 8, Y: 1}, {X: 5, Y: 2}, {X: 7, Y: 3}, {X: 4, Y: 4}}
	suite.CoordsA = []days.Coords{{X: 6, Y: 5}, {X: 8, Y: 8}, {X: 9, Y: 9}}
}

func (suite *EightSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 14, days.DayEightPart1(suite.ProvidedInput))
}

func (suite *EightSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 34, days.DayEightPart2(suite.ProvidedInput))
}

func (suite *EightSuite) TestGetAntennaLocations() {
	grid := days.GetChizu(suite.ProvidedInput)
	assert.Equal(suite.T(), suite.Coords0, grid.GetCoordsFor([]rune{'0'}))
	assert.Equal(suite.T(), suite.CoordsA, grid.GetCoordsFor([]rune{'A'}))
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
	grid := days.GetChizu(suite.ProvidedInput)

	ant0 := []days.Coords{{X: 11, Y: 0}, {X: 2, Y: 3}, {X: 3, Y: 1}, {X: 6, Y: 0}, {X: 6, Y: 5}, {X: 9, Y: 4}, {X: 10, Y: 2}, {X: 0, Y: 7}, {X: 3, Y: 6}, {X: 1, Y: 5}}
	antA := []days.Coords{{X: 4, Y: 2}, {X: 3, Y: 1}, {X: 10, Y: 11}, {X: 7, Y: 7}, {X: 10, Y: 10}}
	calcAnt0 := grid.GetAntenodes(suite.Coords0)
	assert.Len(suite.T(), calcAnt0, len(ant0))
	for _, ant := range ant0 {
		assert.Contains(suite.T(), calcAnt0, ant)
	}
	calcAntA := grid.GetAntenodes(suite.CoordsA)
	assert.Len(suite.T(), calcAntA, len(antA))
	for _, ant := range antA {
		assert.Contains(suite.T(), calcAntA, ant)
	}
}

func (suite *EightSuite) TestGetResonantAntenodes() {
	grid := days.GetChizu(suite.ProvidedInput)
	// crds := []days.Coords{{X: 4, Y: 4}, {X: 6, Y: 8}}

	// ant := []days.Coords{{X: 2, Y: 0}, {X: 3, Y: 2}, {X: 4, Y: 4}, {X: 5, Y: 6}, {X: 6, Y: 8}, {X: 7, Y: 10}}
	// calcAnt := grid.GetResonantAntenodes(crds)
	// assert.Len(suite.T(), calcAnt, len(ant))
	// for _, ant := range ant {
	// 	assert.Contains(suite.T(), calcAnt, ant)
	// }
	// AntB := []days.Coords{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}, {X: 3, Y: 1}, {X: 4, Y: 2}, {X: 4, Y: 4}, {X: 5, Y: 5}, {X: 6, Y: 6}, {X: 6, Y: 5}, {X: 7, Y: 7}, {X: 8, Y: 8}, {X: 9, Y: 9}, {X: 10, Y: 11}, {X: 10, Y: 10}, {X: 11, Y: 11}, {X: 12, Y: 13}}
	resAntA := []days.Coords{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 1, Y: 3}, {X: 3, Y: 3}, {X: 2, Y: 4}, {X: 4, Y: 4}, {X: 5, Y: 5}, {X: 6, Y: 5}, {X: 6, Y: 6}, {X: 7, Y: 7}, {X: 8, Y: 8}, {X: 9, Y: 9}, {X: 10, Y: 10}, {X: 11, Y: 10}, {X: 11, Y: 11}}
	calcResAntA := grid.GetResonantAntenodes(suite.CoordsA)
	assert.Len(suite.T(), calcResAntA, len(resAntA))

	// for _, anta := range resAntA {
	// 	assert.Contains(suite.T(), calcResAntA, anta)
	// }
}

func (suite *EightSuite) TestGetDirection() {
	cordA := days.Coords{X: 2, Y: 0}
	cordB := days.Coords{X: 6, Y: 8}

	dx, dy := cordA.GetDirection(cordB)
	assert.Equal(suite.T(), -1, dx)
	assert.Equal(suite.T(), -2, dy)

	dx, dy = cordB.GetDirection(cordA)
	assert.Equal(suite.T(), 1, dx)
	assert.Equal(suite.T(), 2, dy)
}
