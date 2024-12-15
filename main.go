package main

import (
	"aoc/days"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	day := flag.Int("day", 0, "Specify the Advent of Code day to run")
	part := flag.Int("part", 0, "Advent of Code puzzles have multiple parts, select the part to run")
	inputFile := flag.String("input", "", "Specify the input for the puzzle")

	flag.Parse()
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("You must specify a valid input file")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	var out string
	switch *day {
	case 1:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayOnePart1(lines))
		case 2:
			out = strconv.Itoa(days.DayOnePart2(lines))
		}
	case 2:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayTwoPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayTwoPart2(lines))
		}
	case 3:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayThreePart1(lines))
		case 2:
			out = strconv.Itoa(days.DayThreePart2(lines))
		}
	case 4:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayFourPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayFourPart2(lines))
		}
	case 5:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayFivePart1(lines))
		case 2:
			out = strconv.Itoa(days.DayFivePart2(lines))
		}
	case 6:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DaySixPart1(lines))
		case 2:
			out = strconv.Itoa(days.DaySixPart2(lines))
		}
	case 7:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DaySevenPart1(lines))
		case 2:
			out = strconv.Itoa(days.DaySevenPart2(lines))
		}
	case 8:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayEightPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayEightPart2(lines))
		}
	case 9:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayNinePart1(lines))
		case 2:
			out = strconv.Itoa(days.DayNinePart2(lines))
		}
	case 10:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayTenPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayTenPart2(lines))
		}
	case 11:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayElevenPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayElevenPart2(lines))
		}
	case 12:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayTwelvePart1(lines))
		case 2:
			out = strconv.Itoa(days.DayTwelvePart2(lines))
		}
	case 13:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayThirteenPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayThirteenPart2(lines))
		}
	case 14:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayFourteenPart1(lines, 101, 103))
		case 2:
			out = strconv.Itoa(days.DayFourteenPart2(lines, 101, 103))
		}
	case 15:
		switch *part {
		case 1:
			out = strconv.Itoa(days.DayFifteenPart1(lines))
		case 2:
			out = strconv.Itoa(days.DayFifteenPart2(lines))
		}
		//DAY PLACEHOLDER//
	default:
		fmt.Printf("Invalid day %d provided\n", *day)
		os.Exit(1)
	}
	if out == "" {
		fmt.Printf("Day %d Part %d not implemented yet", *day, *part)
	} else {
		fmt.Printf("Day %d Part %d result: %s\n", *day, *part, out)
	}
}
