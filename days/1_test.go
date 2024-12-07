package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayOnePartOneCasesPresented(t *testing.T) {
	input := []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
	assert.Equal(t, 11, days.DayOnePart1(input))
}

func TestDayOnePartTwoCasesPresented(t *testing.T) {
	input := []string{"3   4", "4   3", "2   5", "1   3", "3   9", "3   3"}
	assert.Equal(t, 31, days.DayOnePart2(input))
}

func TestDayOneParsing(t *testing.T) {
	lines := make([]string, 0)
	lines = append(lines, "81682   36089")
	lines = append(lines, "22289   41038")
	parsed := days.OneParseDistances(lines)
	expected1 := make([]int, 0)
	expected2 := make([]int, 0)
	expected1 = append(expected1, 81682)
	expected1 = append(expected1, 22289)
	expected2 = append(expected2, 36089)
	expected2 = append(expected2, 41038)

	assert.Equal(t, expected1, parsed[0])
	assert.Equal(t, expected2, parsed[1])
}

func TestSortAscending(t *testing.T) {
	one := []int{2, 7, 12, 9, 0}
	oneExpected := []int{0, 2, 7, 9, 12}
	oneSorted := days.SortSliceAscending(one)
	assert.Equal(t, oneExpected, oneSorted)
	assert.NotEqual(t, one, oneSorted)
}

func TestCountNeedlesInHaystack(t *testing.T) {
	haystack := []int{4, 3, 5, 3, 9, 3}
	assert.Equal(t, 0, days.CountNeedlesInHaystack(1, haystack))
	assert.Equal(t, 0, days.CountNeedlesInHaystack(2, haystack))
	assert.Equal(t, 3, days.CountNeedlesInHaystack(3, haystack))
	assert.Equal(t, 1, days.CountNeedlesInHaystack(4, haystack))
	assert.Equal(t, 1, days.CountNeedlesInHaystack(5, haystack))
	assert.Equal(t, 0, days.CountNeedlesInHaystack(6, haystack))
	assert.Equal(t, 0, days.CountNeedlesInHaystack(7, haystack))
	assert.Equal(t, 0, days.CountNeedlesInHaystack(8, haystack))
	assert.Equal(t, 1, days.CountNeedlesInHaystack(9, haystack))

}
