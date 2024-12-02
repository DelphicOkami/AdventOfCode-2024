package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayTwoPartOneCasesPresented(t *testing.T) {
	input := []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}
	assert.Equal(t, 2, days.DayTwoPart1(input))
}

func TestDayTwoPartTwoCasesPresented(t *testing.T) {
	input := []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}
	assert.Equal(t, 4, days.DayTwoPart2(input))

	edges := []string{"48 46 47 49 51 54 56", "1 1 2 3 4 5", "1 2 3 4 5 5", "5 1 2 3 4 5", "1 4 3 2 1", "1 6 7 8 9", "1 2 3 4 3", "9 8 7 6 7", "8 9 10 11"}
	
	assert.Equal(t, 9, days.DayTwoPart2(edges))
}


func TestParseDay2Report(t *testing.T) {
	assert.Equal(t, []int{7,6,4,2,1}, days.ParseDay2Report("7 6 4 2 1"))
}

func TestFindDayTwoSafeReports(t *testing.T) {
	assert.Equal(t, []bool{true}, days.FindDayTwoSafeReports([]string{"7 6 4 2 1"}))
	assert.Equal(t, []bool{false}, days.FindDayTwoSafeReports([]string{"1 3 2 4 5"}))
	input := []string{"7 6 4 2 1", "1 2 7 8 9", "9 7 6 2 1", "1 3 2 4 5", "8 6 4 4 1", "1 3 6 7 9"}
	assert.Equal(t, []bool{true, false, false, false, false, true}, days.FindDayTwoSafeReports(input))
}

func TestIsReportSafe(t *testing.T) {
	assert.Equal(t, true, days.IsReportSafeWithDamnpener(days.ParseDay2Report("7 6 4 2 1")))
	assert.Equal(t, false, days.IsReportSafeWithDamnpener(days.ParseDay2Report("1 2 7 8 9")))
	assert.Equal(t, false, days.IsReportSafeWithDamnpener(days.ParseDay2Report("9 7 6 2 1")))
	assert.Equal(t, true, days.IsReportSafeWithDamnpener(days.ParseDay2Report("1 3 2 4 5")))
	assert.Equal(t, true, days.IsReportSafeWithDamnpener(days.ParseDay2Report("8 6 4 4 1")))
	assert.Equal(t, true, days.IsReportSafeWithDamnpener(days.ParseDay2Report("1 3 6 7 9")))
}