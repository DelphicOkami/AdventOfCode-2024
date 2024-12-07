package days

import (
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func DayFivePart1(input []string) int {
	d := Parse5Input(input)
	correctSets := make([][]int, 0)
	for _, pages := range d.PageSets {
		r := d.GetRuleSetFor(pages)
		if reflect.DeepEqual(FiveGetPageOrder(r, pages), pages) {
			correctSets = append(correctSets, pages)
		}
	}
	out := 0
	for _, i := range correctSets {
		out += i[(len(i)-1)/2]
	}
	return out
}

func DayFivePart2(input []string) int {
	d := Parse5Input(input)
	correctSets := make([][]int, 0)
	for _, pages := range d.PageSets {
		r := d.GetRuleSetFor(pages)
		sorted := FiveGetPageOrder(r, pages)
		if !reflect.DeepEqual(sorted, pages) {
			correctSets = append(correctSets, sorted)
		}
	}
	out := 0
	for _, i := range correctSets {
		out += i[(len(i)-1)/2]
	}
	return out
}

func Parse5Input(input []string) FivePagePrintOrder {
	rules := make([]FiveRule, 0)
	pages := make([][]int, 0)
	for _, line := range input {
		if strings.Contains(line, "|") {
			before, after, _ := strings.Cut(line, "|")
			beforei, _ := strconv.Atoi(before)
			afteri, _ := strconv.Atoi(after)
			rules = append(rules, FiveRule{First: beforei, Second: afteri})
			continue
		}
		if strings.Contains(line, ",") {
			is := make([]int, 0)
			p := strings.Split(line, ",")
			for _, ps := range p {
				i, _ := strconv.Atoi(ps)
				is = append(is, i)
			}
			pages = append(pages, is)
		}
	}
	return FivePagePrintOrder{
		Rules:    rules,
		PageSets: pages,
	}
}

func FiveGetPageOrder(rules []FiveRule, pages []int) []int {
	tmp := make([]int, len(pages))
	copy(tmp, pages)
	f := fiveGetRulesetSort(rules)
	slices.SortFunc(tmp, f)
	return tmp
}

func fiveGetRulesetSort(rules []FiveRule) func(i, j int) int {
	return func(i, j int) int {
		for _, r := range rules {
			if !r.DoesRuleApply([]int{i, j}) {
				continue
			}
			if r.First == i {
				return -1
			}
			return 1
		}
		return 0
	}
}

type FivePagePrintOrder struct {
	Rules    []FiveRule
	PageSets [][]int
}

func (po *FivePagePrintOrder) GetRuleSetFor(pages []int) []FiveRule {
	out := make([]FiveRule, 0)
	for _, r := range po.Rules {
		if r.DoesRuleApply(pages) {
			out = append(out, r)
		}
	}
	return out
}

type FiveRule struct {
	First  int
	Second int
}

func (r *FiveRule) DoesRuleApply(pages []int) bool {
	return slices.Contains(pages, r.First) && slices.Contains(pages, r.Second)
}
