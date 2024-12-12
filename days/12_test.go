package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TwelveSuite struct {
	suite.Suite
	ProvidedInput       []string
	ProvidedSimpleInput []string
}

func TestRunTwelveSuite(t *testing.T) {
	suite.Run(t, new(TwelveSuite))
}

func (suite *TwelveSuite) SetupTest() {
	suite.ProvidedInput = []string{"RRRRIICCFF", "RRRRIICCCF", "VVRRRCCFFF", "VVRCCCJFFF", "VVVVCJJCFE", "VVIVCCJJEE", "VVIIICJJEE", "MIIIIIJJEE", "MIIISIJEEE", "MMMISSJEEE"}
	suite.ProvidedSimpleInput = []string{"AAAA", "BBCD", "BBCC", "EEEC"}
}

func (suite *TwelveSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 1930, days.DayTwelvePart1(suite.ProvidedInput))
}

func (suite *TwelveSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 1206, days.DayTwelvePart2(suite.ProvidedInput))
}

func (suite *TwelveSuite) TestGetPlot() {
	g := days.TweleveGetGarden(suite.ProvidedSimpleInput)
	plot := g.GetPlotFrom(days.Coords{X: 0, Y: 0})
	//A
	expectedPlot := []days.Coords{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}}
	assert.Len(suite.T(), plot.Spots, 4)
	assert.Equal(suite.T(), 10, plot.Perimiter)
	for _, c := range expectedPlot {
		assert.Contains(suite.T(), plot.Spots, c)
	}
	assert.Equal(suite.T(), 4, plot.GetFenceSegments())
	//B
	plot = g.GetPlotFrom(days.Coords{X: 0, Y: 1})
	expectedPlot = []days.Coords{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}}
	assert.Len(suite.T(), plot.Spots, 4)
	assert.Equal(suite.T(), 8, plot.Perimiter)
	for _, c := range expectedPlot {
		assert.Contains(suite.T(), plot.Spots, c)
	}
	assert.Equal(suite.T(), 4, plot.GetFenceSegments())
	//C
	plot = g.GetPlotFrom(days.Coords{X: 2, Y: 1})
	expectedPlot = []days.Coords{{X: 2, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 2}, {X: 3, Y: 3}}
	assert.Len(suite.T(), plot.Spots, 4)
	assert.Equal(suite.T(), 10, plot.Perimiter)
	for _, c := range expectedPlot {
		assert.Contains(suite.T(), plot.Spots, c)
	}
	assert.Equal(suite.T(), 8, plot.GetFenceSegments())
	//D
	plot = g.GetPlotFrom(days.Coords{X: 3, Y: 1})
	expectedPlot = []days.Coords{{X: 3, Y: 1}}
	assert.Len(suite.T(), plot.Spots, 1)
	assert.Equal(suite.T(), 4, plot.Perimiter)
	for _, c := range expectedPlot {
		assert.Contains(suite.T(), plot.Spots, c)
	}
	assert.Equal(suite.T(), 4, plot.GetFenceSegments())
	//E
	plot = g.GetPlotFrom(days.Coords{X: 0, Y: 3})
	expectedPlot = []days.Coords{{X: 0, Y: 3}, {X: 1, Y: 3}, {X: 0, Y: 3}}
	assert.Len(suite.T(), plot.Spots, 3)
	assert.Equal(suite.T(), 8, plot.Perimiter)
	for _, c := range expectedPlot {
		assert.Contains(suite.T(), plot.Spots, c)
	}
	assert.Equal(suite.T(), 4, plot.GetFenceSegments())
}
