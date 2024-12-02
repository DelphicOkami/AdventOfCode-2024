package days

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func DayOnePart1(input []string) int {
	gpSum := 0
	lists := OneParseDistances(input)
	leftList := SortSliceAscending(lists[0])
	rightList := SortSliceAscending(lists[1])

	for i, left := range leftList {
		right := rightList[i]
		if (left >= right) {
			gpSum += left - right
		} else {
			gpSum += right - left
		}
	}
	return gpSum
}
func DayOnePart2(input []string) int {
	similarityScore := 0
	lists := OneParseDistances(input)
	needles := SortSliceAscending(lists[0])
	haystack := SortSliceAscending(lists[1])

	for _, needle := range needles {
		similarityScore += needle * CountNeedlesInHaystack(needle, haystack)
	}
	return similarityScore
}

func OneParseDistances(input []string) [][]int {
	leftList := make([]int, 0)
	rightList := make([]int, 0)
	reg := regexp.MustCompile(`[\s]+`)
	for _, line := range input {
		lp := reg.ReplaceAllString(line, ",")
		lpArts := strings.Split(lp, ",")
		atoiL, _ := strconv.Atoi(lpArts[0])
		atoiR, _ := strconv.Atoi(lpArts[1])
		leftList = append(leftList, atoiL)
		rightList = append(rightList, atoiR)
	}
	output := make([][]int, 0)
	output = append(output, leftList)
	output = append(output, rightList)
	return output
}

func SortSliceAscending(input []int) []int {
	output := make([]int, len(input))
	copy(output, input)
	sort.Slice(output, func(i, j int) bool {
		return output[i] < output[j]
	})
	return output
}

func CountNeedlesInHaystack(needle int, haystack []int) int {
	out := 0
	for _, straw := range haystack {
		if (straw == needle) {
			out++
		}
	}
	return out
}