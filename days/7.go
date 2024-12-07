package days

import (
	"strconv"
	"strings"
)

func DaySevenPart1(input []string) int {
	d7 := ParseDaySevenInput(input)
	out := 0
	for _, line := range d7 {
		if line.DoesWork([]string{"*", "+"}) {
			out += line.Result
		}
	}
	return out
}

func DaySevenPart2(input []string) int {
	d7 := ParseDaySevenInput(input)
	out := 0
	for _, line := range d7 {
		if line.DoesWork([]string{"*", "+", "||"}) {
			out += line.Result
		}
	}
	return out
}

func ParseDaySevenInput(input []string) []DaySevenSum {
	out := make([]DaySevenSum, 0)
	for _, line := range input {
		total, parts, f := strings.Cut(line, ":")
		if !f {
			continue
		}
		sumParts := strings.Split(strings.Trim(parts, " "), " ")
		totalI, _ := strconv.Atoi(total)
		sumpartsI := make([]int, 0)
		for _, p := range sumParts {
			pi, _ := strconv.Atoi(p)
			sumpartsI = append(sumpartsI, pi)
		}
		out = append(out, DaySevenSum{
			Result: totalI,
			Inputs: sumpartsI,
		})
	}
	return out
}

type DaySevenSum struct {
	Result int
	Inputs []int
}

func (s *DaySevenSum) GetPossiblePatterns(ops []string) [][]string {
	patterns := make([][]string, 0)
	binaryRep := ""
	for i := 0; i < len(s.Inputs)-1; i++ {
		binaryRep += strconv.Itoa(len(ops) - 1)
	}
	possiblePatternCount, _ := strconv.ParseInt(binaryRep, len(ops), 64)
	for i := 0; i <= int(possiblePatternCount); i++ {
		patternBinRep := strconv.FormatInt(int64(i), len(ops))
		for len(binaryRep) > len(patternBinRep) {
			patternBinRep = "0" + patternBinRep
		}
		possiblePattern := make([]string, 0)
		for _, p := range patternBinRep {
			index, _ := strconv.Atoi(string(p))
			possiblePattern = append(possiblePattern, ops[index])
		}
		patterns = append(patterns, possiblePattern)
	}
	return patterns
}

func (s *DaySevenSum) GetWorkingPatterns(ops []string) [][]string {
	patterns := s.GetPossiblePatterns(ops)
	working := make([][]string, 0)
	for _, pattern := range patterns {
		prev := s.Inputs[0]
		for i, operation := range pattern {
			switch operation {
			case "*":
				prev = prev * s.Inputs[i+1]
			case "+":
				prev = prev + s.Inputs[i+1]
			case "||":
				prev, _ = strconv.Atoi(strconv.Itoa(prev) + strconv.Itoa(s.Inputs[i+1]))
			}
		}
		if prev == s.Result {
			working = append(working, pattern)
		}
	}
	return working
}
func (s *DaySevenSum) DoesWork(ops []string) bool {
	patterns := s.GetPossiblePatterns(ops)
	for _, pattern := range patterns {
		prev := s.Inputs[0]
		for i, operation := range pattern {
			switch operation {
			case "*":
				prev = prev * s.Inputs[i+1]
			case "+":
				prev = prev + s.Inputs[i+1]
			case "||":
				prev, _ = strconv.Atoi(strconv.Itoa(prev) + strconv.Itoa(s.Inputs[i+1]))
			}
		}
		if prev == s.Result {
			return true
		}
	}
	return false
}
