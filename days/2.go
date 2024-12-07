package days

import (
	"regexp"
	"strconv"
	"strings"
)

func DayTwoPart1(input []string) int {
	safeCount := 0
	reportSafety := FindDayTwoSafeReports(input)
	for _, safe := range reportSafety {
		if safe {
			safeCount++
		}
	}
	return safeCount
}

func DayTwoPart2(input []string) int {
	safeCount := 0
	reportSafety := FindDayTwoSafeReportsWithDampener(input)
	for _, safe := range reportSafety {
		if safe {
			safeCount++
		}
	}
	return safeCount
}

func ParseDay2Report(input string) []int {
	report := make([]int, 0)
	reg := regexp.MustCompile(`[\s]+`)
	reportCsv := reg.ReplaceAllString(input, ",")
	reportParts := strings.Split(reportCsv, ",")
	for _, val := range reportParts {
		intval, _ := strconv.Atoi(val)
		report = append(report, intval)
	}
	return report
}

func FindDayTwoSafeReports(reports []string) []bool {
	safety := make([]bool, len(reports))
	for repi, report := range reports {
		parsedReport := ParseDay2Report(report)
		safety[repi] = IsReportSafeWithoutDamnpener(parsedReport)
	}
	return safety
}

func FindDayTwoSafeReportsWithDampener(reports []string) []bool {
	safety := make([]bool, len(reports))
	for repi, report := range reports {
		parsedReport := ParseDay2Report(report)
		safety[repi] = IsReportSafeWithDamnpener(parsedReport)
	}
	return safety
}

func IsReportSafeWithDamnpener(report []int) bool {
	reportClone := make([]int, len(report))
	copy(reportClone, report)
	safe := IsReportSafeWithoutDamnpener(report)
	if safe {
		return true
	}
	unsafeIndexes := make([]int, 0)
	for i, _ := range report {
		safeReport := RemoveIntIndex(reportClone, i)
		if IsReportSafeWithoutDamnpener(safeReport) {
			unsafeIndexes = append(unsafeIndexes, i)
		}
	}
	if len(unsafeIndexes) == 0 {
		return false
	}
	if len(unsafeIndexes) == 1 {
		return true
	}
	if len(unsafeIndexes) > 2 {
		return false
	} else {
		return (unsafeIndexes[1]-unsafeIndexes[0] == 1)
	}
}

func IsReportSafeWithoutDamnpener(report []int) bool {
	var prevValue = 0
	var diff = 0
	var increase = true
	var decrease = true

	for i, val := range report {
		if i == 0 {
			prevValue = val
			continue
		}
		diff = prevValue - val
		if diff == 0 {
			return false
		}
		if increase {
			if diff > 3 {
				return false
			} else if diff < 0 {
				increase = false
			}
		}

		if decrease {
			if diff < -3 {
				return false
			} else if diff > 0 {
				decrease = false
			}
		}

		if !increase && !decrease {
			return false
		}

		prevValue = val
	}
	return true
}

func RemoveIntIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	ret = append(ret, s[index+1:]...)
	return ret
}
