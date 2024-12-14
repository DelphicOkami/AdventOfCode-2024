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
	suite.ProvidedInput = []string{"p=0,4 v=3,-3", "p=6,3 v=-1,-3", "p=10,3 v=-1,2", "p=2,0 v=2,-1", "p=0,0 v=1,3", "p=3,0 v=-2,-2", "p=7,6 v=-1,-3", "p=3,0 v=-1,-2", "p=9,3 v=2,3", "p=7,3 v=-1,2", "p=2,4 v=2,-3", "p=9,5 v=-3,-3"}
}

func (suite *FourteenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 12, days.DayFourteenPart1(suite.ProvidedInput, 11, 7))
}

// func (suite *FourteenSuite) TestTwoCasesPresented() {
// 	assert.Equal(suite.T(), 0, days.DayFourteenPart2(suite.ProvidedInput))
// }

func (suite *FourteenSuite) TestParse() {
	BathroomMap := days.DayFourteenParse(suite.ProvidedInput, 11, 7)
	assert.Len(suite.T(), BathroomMap.Robots, 12)
	r, ok := BathroomMap.Robots[days.FourteenRobotStart{days.Coords{X: 0, Y: 4}, 0}]
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), 11, BathroomMap.XSize)
	assert.Equal(suite.T(), 7, BathroomMap.YSize)
	assert.Equal(suite.T(), days.Coords{X: 0, Y: 4}, r.Position)
	assert.Equal(suite.T(), 3, r.SpeedX)
	assert.Equal(suite.T(), -3, r.SpeedY)
}
func (suite *FourteenSuite) TestAdvanceTime() {
	BathroomMap := days.DayFourteenParse(suite.ProvidedInput, 11, 7)
	BathroomMap.AdvanceTime(0)
	assert.Len(suite.T(), BathroomMap.Robots, 12)
	r, ok := BathroomMap.Robots[days.FourteenRobotStart{days.Coords{X: 0, Y: 4}, 0}]
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), days.Coords{X: 0, Y: 4}, r.Position)
	BathroomMap.AdvanceTime(1)
	r, ok = BathroomMap.Robots[days.FourteenRobotStart{days.Coords{X: 0, Y: 4}, 0}]
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), days.Coords{X: 3, Y: 1}, r.Position)
	BathroomMap.AdvanceTime(1)
	r, ok = BathroomMap.Robots[days.FourteenRobotStart{days.Coords{X: 0, Y: 4}, 0}]
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), days.Coords{X: 6, Y: 5}, r.Position)
	BathroomMap.AdvanceTime(2)
	r, ok = BathroomMap.Robots[days.FourteenRobotStart{days.Coords{X: 0, Y: 4}, 0}]
	assert.True(suite.T(), ok)
	assert.Equal(suite.T(), days.Coords{X: 1, Y: 6}, r.Position)
}

func (suite *FourteenSuite) TestQuadrantTest() {
	BathroomMap := days.DayFourteenParse(suite.ProvidedInput, 11, 7)
	q1, q2, q3, q4 := BathroomMap.GetRobotsPerQuadrant()
	assert.Equal(suite.T(), 4, q1)
	assert.Equal(suite.T(), 0, q2)
	assert.Equal(suite.T(), 2, q3)
	assert.Equal(suite.T(), 2, q4)

	BathroomMap.AdvanceTime(100)
	q1, q2, q3, q4 = BathroomMap.GetRobotsPerQuadrant()
	assert.Equal(suite.T(), 1, q1)
	assert.Equal(suite.T(), 3, q2)
	assert.Equal(suite.T(), 4, q3)
	assert.Equal(suite.T(), 1, q4)

}
