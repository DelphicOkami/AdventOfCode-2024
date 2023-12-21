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
					
				}
	default:
		fmt.Println("Invalid day provided")
		os.Exit(1)
	}
	if out == "" {
		fmt.Println("invalid Day / part combo")
	} else {
		fmt.Printf("Day %d Part %d result: %s\n", *day, *part, out)
	}

	
}