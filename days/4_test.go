package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayFourPartOneCasesPresented(t *testing.T) {
	input := []string{"MMMSXXMASM", "MSAMXMSMSA", "AMXSXMAAMM", "MSAMASMSMX", "XMASAMXAMM", "XXAMMXXAMA", "SMSMSASXSS", "SAXAMASAAA", "MAMMMXMMMM", "MXMXAXMASX"}
	assert.Equal(t, 18, days.DayFourPart1(input))
}
func TestDayFourPartTwoCasesPresented(t *testing.T) {
	input := []string{"MMMSXXMASM", "MSAMXMSMSA", "AMXSXMAAMM", "MSAMASMSMX", "XMASAMXAMM", "XXAMMXXAMA", "SMSMSASXSS", "SAXAMASAAA", "MAMMMXMMMM", "MXMXAXMASX"}
	assert.Equal(t, 9, days.DayFourPart2(input))
}

func TestGetDay4StringToRuneSlice(t *testing.T) {
	input := []string{"MMMSXXMASM"}
	d4 := days.GetDay4(input)

	assert.Equal(t, 'M', d4.Grid[0][0])
	assert.Equal(t, 'M', d4.Grid[0][1])
	assert.Equal(t, 'M', d4.Grid[0][2])
	assert.Equal(t, 'S', d4.Grid[0][3])
	assert.Equal(t, 'X', d4.Grid[0][4])
}


func TestGetRunFromGrid(t *testing.T) {
	input := []string{"MMMSXXMASM", "MSAMXMSMSA", "AMXSXMAAMM", "MSAMASMSMX", "XMASAMXAMM", "XXAMMXXAMA", "SMSMSASXSS", "SAXAMASAAA", "MAMMMXMMMM", "MXMXAXMASX"}
	d4 := days.GetDay4(input)
	r, err := d4.GetRune(3, 4)
	assert.Equal(t, 'A', r)
	assert.Nil(t, err)

	_, err = d4.GetRune(3, 10)
	assert.NotNil(t, err)
	_, err = d4.GetRune(-1, 3)
	assert.NotNil(t, err)
	_, err = d4.GetRune(3, -1)
	assert.NotNil(t, err)
	_, err = d4.GetRune(10, 3)
	assert.NotNil(t, err)
}

func TestDayFourFindNXmas(t *testing.T) {
	input := []string{"MMMSXXMASM", "MSAMXMSMSA", "AMXSXMAAMM", "MSAMASMSMX", "XMASAMXAMM", "XXAMMXXAMA", "SMSMSASXSS", "SAXAMASAAA", "MAMMMXMMMM", "MXMXAXMASX"}
	d4 := days.GetDay4(input)
	assert.Equal(t, 2, d4.FindDirectionXmasCount(-1, 0))
}