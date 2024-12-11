package days

import (
	"strconv"
	"strings"
)

var elevenPreMapped = map[elevenToBlink]int{}

func DayElevenPart1(input []string) int {
	singleLine := strings.Join(input, "")
	startStones := strings.Split(singleLine, " ")
	startStonesi := make([]int, 0)
	for _, s := range startStones {
		i, _ := strconv.Atoi(s)
		startStonesi = append(startStonesi, i)
	}
	return ElevenGetStoneCount(startStonesi, 25)

}

func DayElevenPart2(input []string) int {
	singleLine := strings.Join(input, "")
	startStones := strings.Split(singleLine, " ")
	startStonesi := make([]int, 0)
	for _, s := range startStones {
		i, _ := strconv.Atoi(s)
		startStonesi = append(startStonesi, i)
	}
	return ElevenGetStoneCount(startStonesi, 75)
}

func ElevenHardBlink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	stoneString := strconv.Itoa(stone)
	if len(stoneString)%2 == 0 {
		x, _ := strconv.Atoi(stoneString[0 : len(stoneString)/2])
		y, _ := strconv.Atoi(stoneString[len(stoneString)/2:])
		return []int{x, y}
	}
	return []int{stone * 2024}
}

func ElevenGetStoneCount(stones []int, blinks int) int {
	total := 0
	for _, s := range stones {
		total += ElevenGetStonesAfter(elevenToBlink{Number: s, RemainingBlinks: blinks})
	}
	return total
}

func ElevenGetStonesAfter(st elevenToBlink) int {
	if st.RemainingBlinks == 0 {
		return 1
	}
	val, ok := elevenPreMapped[st]
	if ok {
		return val
	}

	blunk := ElevenHardBlink(st.Number)

	elevenPreMapped[st] = ElevenGetStoneCount(blunk, st.RemainingBlinks - 1)
	return elevenPreMapped[st]
}

type elevenToBlink struct {
	Number          int
	RemainingBlinks int
}
