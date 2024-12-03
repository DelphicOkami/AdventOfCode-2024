package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDayThreePartOneCasesPresented(t *testing.T) {
	input := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}
	assert.Equal(t, 161, days.DayThreePart1(input))
}

func TestDayThreePartTwoCasesPresented(t *testing.T) {
	input := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}
	assert.Equal(t, 48, days.DayThreePart2(input))
}

func TestGetMultiplicationList(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	mul := days.GetThreeMultiplicationList(input, true)
	assert.Equal(t, 4, len(mul))
	assert.Equal(t, 2, mul[0].GetLeftValue())
	assert.Equal(t, 4, mul[0].GetRightValue())
	assert.Equal(t, 11, mul[2].GetLeftValue())
	assert.Equal(t, 8, mul[2].GetRightValue())
}