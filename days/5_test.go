package days_test

import (
	"testing"
	"aoc/days"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type D5Suite struct {
    suite.Suite
    ProvidedInput []string
	SoilToFertilizer []days.D5MapRange
}

func TestRunD5Suite(t *testing.T) {
	suite.Run(t, new(D5Suite))
}

func (suite *D5Suite) SetupTest() {
    suite.ProvidedInput = []string{
		"seeds: 79 14 55 13",
		"",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	} 

	suite.SoilToFertilizer = []days.D5MapRange{
		{
			SourceStart: 15,
			SourceEnd: 51,
			DestStart: 0,
			DestEnd: 36,
		},
		{
			SourceStart: 52,
			SourceEnd: 53,
			DestStart: 37,
			DestEnd: 38,
		},
		{
			SourceStart: 0,
			SourceEnd: 14,
			DestStart: 39,
			DestEnd: 53,
		},
	}
}

func (suite *D5Suite) TestMapParsing() {
	expected := days.D5Map{
		Seeds: []int{79, 14, 55, 13},
		SeedToSoil: []days.D5MapRange{
			{
				SourceStart: 98,
				SourceEnd: 99,
				DestStart: 50,
				DestEnd: 51,
			},
			{
				SourceStart: 50,
				SourceEnd: 97,
				DestStart: 52,
				DestEnd: 99,
			},
		},
		SoilToFertilizer: suite.SoilToFertilizer,
		FertilizerToWater: []days.D5MapRange{
			{
				SourceStart: 53,
				SourceEnd: 60,
				DestStart: 49,
				DestEnd: 56,
			},
			{
				SourceStart: 11,
				SourceEnd: 52,
				DestStart: 0,
				DestEnd: 41,
			},
			{
				SourceStart: 0,
				SourceEnd: 6,
				DestStart: 42,
				DestEnd: 48,
			},
			{
				SourceStart: 7,
				SourceEnd: 10,
				DestStart: 57,
				DestEnd: 60,
			},
		},
		WaterToLight: []days.D5MapRange{
			{
				SourceStart: 18,
				SourceEnd: 24,
				DestStart: 88,
				DestEnd: 94,
			},
			{
				SourceStart: 25,
				SourceEnd: 94,
				DestStart: 18,
				DestEnd: 87,
			},
		},
		LightToTemperature: []days.D5MapRange{
			{
				SourceStart: 77,
				SourceEnd: 99,
				DestStart: 45,
				DestEnd: 67,
			},
			{
				SourceStart: 45,
				SourceEnd: 63,
				DestStart: 81,
				DestEnd: 99,
			},
			{
				SourceStart: 64,
				SourceEnd: 76,
				DestStart: 68,
				DestEnd: 80,
			},
		},
		TemperatureToHumidity: []days.D5MapRange{
			{
				SourceStart: 69,
				SourceEnd: 69,
				DestStart: 0,
				DestEnd: 0,
			},
			{
				SourceStart: 0,
				SourceEnd: 68,
				DestStart: 1,
				DestEnd: 69,
			},
		},
		HumidityToLocation: []days.D5MapRange{
			{
				SourceStart: 56,
				SourceEnd: 92,
				DestStart: 60,
				DestEnd: 96,
			},
			{
				SourceStart: 93,
				SourceEnd: 96,
				DestStart: 56,
				DestEnd: 59,
			},
		},
	}

	assert.Equal(suite.T(), expected, days.D5ParseInput(suite.ProvidedInput))
}

func (suite *D5Suite) TestMapRangeParsing() {
	input := "56 93 4"
	expected := days.D5MapRange{
		DestStart: 56,
		DestEnd: 59,
		SourceStart: 93,
		SourceEnd: 96,
	}
	assert.Equal(suite.T(), expected, days.D5ParseMapRange(input))
}

func (suite *D5Suite) TestRangeConversion() {
	assert.Equal(suite.T(), 100, days.D5ConvertSourceToDest(100, suite.SoilToFertilizer))
	assert.Equal(suite.T(), 0, days.D5ConvertSourceToDest(15, suite.SoilToFertilizer))
	assert.Equal(suite.T(), 36, days.D5ConvertSourceToDest(51, suite.SoilToFertilizer))
	assert.Equal(suite.T(), 37, days.D5ConvertSourceToDest(52, suite.SoilToFertilizer))
	assert.Equal(suite.T(), 38, days.D5ConvertSourceToDest(53, suite.SoilToFertilizer))
	assert.Equal(suite.T(), 39, days.D5ConvertSourceToDest(0, suite.SoilToFertilizer))
	assert.Equal(suite.T(), 53, days.D5ConvertSourceToDest(14, suite.SoilToFertilizer))
}

func (suite *D5Suite) TestSeedToLocationConversion() {
	d5map := days.D5ParseInput(suite.ProvidedInput)
	assert.Equal(suite.T(), 82, d5map.SeedToLocation(79))
	assert.Equal(suite.T(), 43, d5map.SeedToLocation(14))
	assert.Equal(suite.T(), 86, d5map.SeedToLocation(55))
	assert.Equal(suite.T(), 35, d5map.SeedToLocation(13))
}

func (suite *D5Suite) TestGetClosestSeedLocation() {
	d5map := days.D5ParseInput(suite.ProvidedInput)
	assert.Equal(suite.T(), 35, d5map.GetClosestSeedLocation())
}