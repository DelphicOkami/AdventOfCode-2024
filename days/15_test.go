package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FifteenSuite struct {
	suite.Suite
	ProvidedInputLarge []string
	ProvidedInputSmall []string
}

func TestRunFifteenSuite(t *testing.T) {
	suite.Run(t, new(FifteenSuite))
}

func (suite *FifteenSuite) SetupTest() {
	suite.ProvidedInputLarge = []string{"##########", "#..O..O.O#", "#......O.#", "#.OO..O.O#", "#..O@..O.#", "#O#..O...#", "#O..O..O.#", "#.OO.O.OO#", "#....O...#", "##########", "", "<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^", "vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v", "><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<", "<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^", "^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><", "^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^", ">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^", "<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>", "^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>", "v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^"}
	suite.ProvidedInputSmall = []string{"########", "#..O.O.#", "##@.O..#", "#...O..#", "#.#.O..#", "#...O..#", "#......#", "########", "", "<^^>>>vv<v>>v<<"}
}

func (suite *FifteenSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 10092, days.DayFifteenPart1(suite.ProvidedInputLarge))
}

func (suite *FifteenSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 0, days.DayFifteenPart2(suite.ProvidedInputLarge))
}

func (suite *FifteenSuite) TestParse() {
	expected := days.FifteenWarehouse{
		Chizu: days.Chizu{
			Grid: [][]rune{
				{'#', '#', '#', '#', '#', '#', '#', '#'},
				{'#', '.', '.', 'O', '.', 'O', '.', '#'},
				{'#', '#', '@', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', 'O', '.', '.', '#'},
				{'#', '.', '#', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', 'O', '.', '.', '#'},
				{'#', '.', '.', '.', '.', '.', '.', '#'},
				{'#', '#', '#', '#', '#', '#', '#', '#'},
			},
		},
		Robot:      days.Coords{X: 2, Y: 2},
		Directions: []rune{'<', '^', '^', '>', '>', '>', 'v', 'v', '<', 'v', '>', '>', 'v', '<', '<'},
	}

	assert.Equal(suite.T(), expected, days.FifteenParseInput(suite.ProvidedInputSmall))
}

func (suite *FifteenSuite) TestMoveRobot() {
	up := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '@', 'O', '.', 'O', '.', '#'},
		{'#', '#', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '#', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#'},
	}
	upRight := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '.', '@', 'O', 'O', '.', '#'},
		{'#', '#', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '#', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#'},
	}
	upRightRight := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '.', '.', '@', 'O', 'O', '#'},
		{'#', '#', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '#', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#'},
	}
	warehouse := days.FifteenParseInput(suite.ProvidedInputSmall)
	warehouse.MoveRobot(0, -1)
	assert.Equal(suite.T(), up, warehouse.Chizu.Grid)
	warehouse.MoveRobot(1, 0)
	assert.Equal(suite.T(), upRight, warehouse.Chizu.Grid)
	warehouse.MoveRobot(1, 0)
	assert.Equal(suite.T(), upRightRight, warehouse.Chizu.Grid)
	warehouse.MoveRobot(1, 0)
	assert.Equal(suite.T(), upRightRight, warehouse.Chizu.Grid)
}

func (suite *FifteenSuite) TestMoveSeq() {
	upRightRight := [][]rune{
		{'#', '#', '#', '#', '#', '#', '#', '#'},
		{'#', '.', '.', '.', '@', 'O', 'O', '#'},
		{'#', '#', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '#', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', 'O', '.', '.', '#'},
		{'#', '.', '.', '.', '.', '.', '.', '#'},
		{'#', '#', '#', '#', '#', '#', '#', '#'},
	}
	warehouse := days.FifteenParseInput(suite.ProvidedInputSmall)
	warehouse.Directions = []rune{'^', '>', '>', '>'}
	warehouse.FollowMoveSequence()
	assert.Equal(suite.T(), upRightRight, warehouse.Chizu.Grid)
}
