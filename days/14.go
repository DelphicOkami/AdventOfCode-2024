package days

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"

	"github.com/manifoldco/promptui"
)

func DayFourteenPart1(input []string, xSize, ySize int) int {
	bathroomMap := DayFourteenParse(input, xSize, ySize)
	bathroomMap.AdvanceTime(100)
	q0, q1, q2, q3 := bathroomMap.GetRobotsPerQuadrant()
	return q0 * q1 * q2 * q3
}

func DayFourteenPart2(input []string, xSize, ySize int) int {
	bathroomMap := DayFourteenParse(input, xSize, ySize)
	tree := false
	for !tree {
		bathroomMap.AdvanceTime(1)
		if bathroomMap.HasABotLine(7) {
			fmt.Print(bathroomMap.Render())
			fmt.Println("Is there a tree?")
			tree = yesNo()
		}
	}
	return bathroomMap.Time
}
func yesNo() bool {
	prompt := promptui.Select{
		Label: "Select[Yes/No]",
		Items: []string{"Yes", "No"},
	}
	_, result, _ := prompt.Run()
	return result == "Yes"
}

func DayFourteenParse(input []string, xSize, ySize int) FourteenBathroomMap {
	out := FourteenBathroomMap{
		Robots: make(map[FourteenRobotStart]FourteenRobot, 0),
		XSize:  xSize,
		YSize:  ySize,
		Time:   0,
	}

	botRegex := regexp.MustCompile(`p=([\d]+),([\d]+) v=(\-?[\d]+),(\-?[\d]+)`)
	for i, line := range input {
		botParts := botRegex.FindStringSubmatch(line)
		px, _ := strconv.Atoi(botParts[1])
		py, _ := strconv.Atoi(botParts[2])
		sx, _ := strconv.Atoi(botParts[3])
		sy, _ := strconv.Atoi(botParts[4])
		startPos := Coords{X: px, Y: py}
		out.Robots[FourteenRobotStart{startPos, i}] = FourteenRobot{
			Position: startPos,
			SpeedX:   sx,
			SpeedY:   sy,
		}
	}
	return out
}

type FourteenBathroomMap struct {
	Robots map[FourteenRobotStart]FourteenRobot
	XSize  int
	YSize  int
	Time   int
}

func (b *FourteenBathroomMap) AdvanceTime(seconds int) {
	b.Time += seconds
	for k, r := range b.Robots {
		pos := k.Position.GetCoordAt(b.Time*r.SpeedX, b.Time*r.SpeedY)
		pos.X = pos.X % b.XSize
		pos.Y = pos.Y % b.YSize
		if pos.X < 0 {
			pos.X = b.XSize + pos.X
		}
		if pos.Y < 0 {
			pos.Y = b.YSize + pos.Y
		}
		r.Position = pos
		b.Robots[k] = r
	}
}

func (b *FourteenBathroomMap) HasABotLine(size int) bool {
	yArray := make(map[int][]int)
	for _, r := range b.Robots {
		_, ok := yArray[r.Position.Y]
		if !ok {
			yArray[r.Position.Y] = make([]int, 0)
		}
		yArray[r.Position.Y] = append(yArray[r.Position.Y], r.Position.X)
	}
	for _, row := range yArray {
		slices.Sort(row)
		if len(row) < size {
			continue
		}
		for i, x := range row {
			if i <= size {
				continue
			}
			consec := true
			for j := 1; j <= size; j++ {
				if x != row[i-j]+j {
					consec = false
					break
				}
			}
			if consec {
				return true
			}
		}
	}

	return false
}

func (b *FourteenBathroomMap) GetRobotsPerQuadrant() (int, int, int, int) {
	q0, q1, q2, q3 := 0, 0, 0, 0
	centreX := b.XSize / 2
	centreY := b.YSize / 2
	for _, r := range b.Robots {
		if r.Position.X == centreX || r.Position.Y == centreY {
			// This robot is in the deadzone
			continue
		}
		if r.Position.X < centreX && r.Position.Y < centreY {
			q0++
		} else if r.Position.X > centreX && r.Position.Y < centreY {
			q1++
		} else if r.Position.X < centreX && r.Position.Y > centreY {
			q2++
		} else {
			q3++
		}
	}
	return q0, q1, q2, q3
}
func (b *FourteenBathroomMap) Render() string {
	out := ""
	positions := make([]Coords, 0)
	for _, r := range b.Robots {
		positions = append(positions, r.Position)
	}
	for y := 0; y < b.YSize; y++ {
		for x := 0; x < b.XSize; x++ {
			pos := Coords{X: x, Y: y}
			if slices.Contains(positions, pos) {
				out += "@"
			} else {
				out += " "
			}
		}
		out += "\n"
	}
	return out
}

type FourteenRobot struct {
	Position Coords
	SpeedX   int
	SpeedY   int
}

type FourteenRobotStart struct {
	Position Coords
	Index    int
}
