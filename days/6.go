package days

import "strings"

type D6Race struct{
	Time int
	TargetDistance int
	winOptions []int
}

func D6ParseRaces(input []string) []D6Race {
	var times []int
	var distances []int
	for _, line := range input {
		if strings.Contains(line, "Time:") {
			times = spaceSeparatedNumbersToIntSlice(strings.Replace(strings.ToLower(line), "time:", "", 1))
		} else if strings.Contains(line, "Distance:") {
			distances = spaceSeparatedNumbersToIntSlice(strings.Replace(strings.ToLower(line), "distances:", "", 1))
		}
	}
	races := make([]D6Race, 0)
	for i, time := range times {
		races = append(races, D6Race{
			Time: time,
			TargetDistance: distances[i],
		})
	}
	return races
}

func (r D6Race) GetWinOptions() []int {
	if len(r.winOptions) == 0 {
		r.winOptions = make([]int, 0)
		last := -1
		for ms := 1; ms < r.Time; ms++ {
			dist := (r.Time - ms) * ms
			if dist > r.TargetDistance {
				r.winOptions = append(r.winOptions, ms)
			} else if dist < last {
				break
			}
			last = dist
		} 
	}
	return r.winOptions
}

func D6P1CalculateTotal(races []D6Race) int {
	total := 1
	for _, race := range races {
		total = total * len(race.GetWinOptions())
	}
	return total
}