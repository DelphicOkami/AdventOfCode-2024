package days_test

import (
	"aoc/days"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FiveSuite struct {
	suite.Suite
	ProvidedInput []string
}

func TestRunFiveSuite(t *testing.T) {
	suite.Run(t, new(FiveSuite))
}

func (suite *FiveSuite) SetupTest() {
	suite.ProvidedInput = []string{"47|53", "97|13", "97|61", "97|47", "75|29", "61|13", "75|53", "29|13", "97|29", "53|29", "61|53", "97|53", "61|29", "47|13", "75|47", "97|75", "47|61", "75|61", "47|29", "75|13", "53|13", "75,47,61,53,29", "97,61,53,29,13", "75,29,13", "75,97,47,61,53", "61,13,29", "97,13,75,29,47"}
}

func (suite *FiveSuite) TestOneCasesPresented() {
	assert.Equal(suite.T(), 143, days.DayFivePart1(suite.ProvidedInput))
}
func (suite *FiveSuite) TestTwoCasesPresented() {
	assert.Equal(suite.T(), 123, days.DayFivePart2(suite.ProvidedInput))
}

func (suite *FiveSuite) TestParse() {
	input := []string{"47|53", "97|13", "97|61", "75,47,61,53,29", "75,29,13"}
	expected := days.FivePagePrintOrder{
		Rules:    []days.FiveRule{{First: 47, Second: 53}, {First: 97, Second: 13}, {First: 97, Second: 61}},
		PageSets: [][]int{{75, 47, 61, 53, 29}, {75, 29, 13}},
	}
	assert.Equal(suite.T(), expected, days.Parse5Input(input))
}

func (suite *FiveSuite) TestRuleApplication() {
	rule := days.FiveRule{First: 47, Second: 53}
	assert.True(suite.T(), rule.DoesRuleApply([]int{75, 47, 61, 53, 29}))
	assert.False(suite.T(), rule.DoesRuleApply([]int{75, 61, 53, 29}))
	assert.False(suite.T(), rule.DoesRuleApply([]int{75, 61, 47, 29}))
	assert.False(suite.T(), rule.DoesRuleApply([]int{75, 61, 29}))
}

func (suite *FiveSuite) TestGetRuleList() {
	po := days.FivePagePrintOrder{
		Rules:    []days.FiveRule{{First: 47, Second: 53}, {First: 97, Second: 13}, {First: 97, Second: 61}},
		PageSets: [][]int{{75, 47, 61, 53, 29}, {75, 29, 13}},
	}
	assert.Equal(suite.T(), []days.FiveRule{{First: 47, Second: 53}}, po.GetRuleSetFor(po.PageSets[0]))
}

func (suite *FiveSuite) TestPageSorting() {
	d := days.Parse5Input(suite.ProvidedInput)

	assert.Equal(suite.T(), []int{75, 47, 61, 53, 29}, days.FiveGetPageOrder(d.Rules, []int{75, 47, 61, 53, 29}))
	assert.Equal(suite.T(), []int{61, 29, 13}, days.FiveGetPageOrder(d.Rules, []int{61, 13, 29}))
}
