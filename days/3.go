package days

import (
	"regexp"
	"strconv"
)

func DayThreePart1(input []string) int {
	result := 0
	for _, str := range input {
		mulList := GetMultiplicationList(str)
		for _, mul := range mulList {
			result += mul.Multiply()
		}
	}
	return result
}

func GetMultiplicationList(input string) []D3Mul {
	out := make([]D3Mul, 0)
	reg := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	matches := reg.FindAllString(input, -1)
	if matches == nil {
		return out
	}

	for _, match := range matches {
		out = append(out, D3Mul{
			Original: match,
			Active: true,
		})
	}
	return out
}

type D3Mul struct {
	Original string
	leftValue int
	rightValue int
	Active bool
}

func (d *D3Mul) GetLeftValue() int {
	if d.leftValue == 0 {
		d.matchValues()
	}

	return d.leftValue
}

func (d *D3Mul) GetRightValue() int {
	if d.rightValue == 0 {
		d.matchValues()
	}

	return d.rightValue
}

func (d *D3Mul) matchValues() {
	reg := regexp.MustCompile(`([\d]{1,3}),([\d]{1,3})`)
	matches := reg.FindStringSubmatch(d.Original)
	if len(matches) != 3 {
		return
	}
	d.leftValue, _ = strconv.Atoi(matches[1])
	d.rightValue, _ = strconv.Atoi(matches[2])
}

func (d *D3Mul) Multiply() int {
	if !d.Active {
		return 0
	}
	return d.GetLeftValue() * d.GetRightValue()
}