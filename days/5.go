package days

import "strings"


type D5MapRange struct {
	SourceStart int
	SourceEnd int
	DestStart int
	DestEnd int
}

type D5Map struct {
	Seeds []int
	SeedToSoil []D5MapRange
	SoilToFertilizer []D5MapRange
	FertilizerToWater []D5MapRange
	WaterToLight []D5MapRange
	LightToTemperature []D5MapRange
	TemperatureToHumidity []D5MapRange
	HumidityToLocation []D5MapRange
}

func D5ParseInput(input []string) D5Map {
	seedsLine, input := input[0], input[1:]
	blockName := ""
	seedToSoil:= make([]D5MapRange, 0)
	soilToFertilizer:= make([]D5MapRange, 0)
	fertilizerToWater:= make([]D5MapRange, 0)
	waterToLight:= make([]D5MapRange, 0)
	lightToTemperature:= make([]D5MapRange, 0)
	temperatureToHumidity:= make([]D5MapRange, 0)
	humidityToLocation:= make([]D5MapRange, 0)
	for _, line := range input {
		if line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			blockName = strings.Trim(line, " ")
			continue
		}
		rng := D5ParseMapRange(line)
		switch blockName {
		case "seed-to-soil map:":
			seedToSoil = append(seedToSoil, rng)
		case "soil-to-fertilizer map:":
			soilToFertilizer = append(soilToFertilizer, rng)
		case "fertilizer-to-water map:":
			fertilizerToWater = append(fertilizerToWater, rng)
		case "water-to-light map:":
			waterToLight = append(waterToLight, rng)
		case "light-to-temperature map:":
			lightToTemperature = append(lightToTemperature, rng)
		case "temperature-to-humidity map:":
			temperatureToHumidity = append(temperatureToHumidity, rng)
		case "humidity-to-location map:":
			humidityToLocation = append(humidityToLocation, rng)
		}
	}

	return D5Map{
		Seeds: spaceSeparatedNumbersToIntSlice(strings.Replace(strings.ToLower(seedsLine), "seeds:", "", 1)),
		SeedToSoil: seedToSoil,
		SoilToFertilizer: soilToFertilizer,
		FertilizerToWater: fertilizerToWater,
		WaterToLight: waterToLight,
		LightToTemperature: lightToTemperature,
		TemperatureToHumidity: temperatureToHumidity,
		HumidityToLocation: humidityToLocation,
	}
}

func D5ParseMapRange(rang string) D5MapRange {
	bits := spaceSeparatedNumbersToIntSlice(rang)
	return D5MapRange{
		SourceStart: bits[1],
		DestStart: bits[0],
		SourceEnd: bits[1] + bits[2] - 1,
		DestEnd: bits[0] + bits[2] - 1,
	}
}

func (mr D5MapRange) IsInSourceRange(test int) bool {
	return test >= mr.SourceStart && test <= mr.SourceEnd
}

func D5ConvertSourceToDest(source int, rangs []D5MapRange) int {
	for _, mr := range rangs {
		if mr.IsInSourceRange(source) {
			offset := source - mr.SourceStart
			return mr.DestStart + offset
		}
	}
	
	return source
}

func (m D5Map) SeedToLocation(seed int) int {
	soil := D5ConvertSourceToDest(seed, m.SeedToSoil)
	fertilizer := D5ConvertSourceToDest(soil, m.SoilToFertilizer)
	water := D5ConvertSourceToDest(fertilizer, m.FertilizerToWater)
	light := D5ConvertSourceToDest(water, m.WaterToLight)
	temperature := D5ConvertSourceToDest(light, m.LightToTemperature)
	humitidiy := D5ConvertSourceToDest(temperature, m.TemperatureToHumidity)
	location := D5ConvertSourceToDest(humitidiy, m.HumidityToLocation)
	return location
}

func (m D5Map) GetClosestSeedLocation() int {
	location := -1
	for _, seed := range m.Seeds {
		seedLoc := m.SeedToLocation(seed)
		if seedLoc < location || location == -1 {
			location = seedLoc
		}
	}
	return location
}