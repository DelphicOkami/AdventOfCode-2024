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
					d5map := days.D5ParseInput(lines)
					out = strconv.Itoa(d5map.GetClosestSeedLocation())
				case 2: 
					fmt.Println("This is inefficient and takes a while to run, but it gets there")
					d5map := days.D5ParseInput(lines)
					out = strconv.Itoa(d5map.GetClosestSeedLocationFromRanges())
				}

			case 6:
				switch *part {
					case 1: 
						races := days.D6ParseRaces(lines)
						out = strconv.Itoa(days.D6P1CalculateTotal(races))
					case 2: 
						race := days.D6ParseRace(lines)
						out = strconv.Itoa(len(race.GetWinOptions()))
					}
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