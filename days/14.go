package days

import (
	"regexp"
	"strconv"
)

func DayFourteenPart1(input []string, xSize, ySize int) int {
	bathroomMap := DayFourteenParse(input, xSize, ySize)
	bathroomMap.AdvanceTime(100)
	q0, q1, q2, q3 := bathroomMap.GetRobotsPerQuadrant()
	return q0 * q1 * q2 * q3
}

func DayFourteenPart2(input []string) int {
	return 0
}

func DayFourteenParse(input []string, xSize, ySize int) FourteenBathroomMap {
	out := FourteenBathroomMap{
		Robots: make(map[FourteenRobotStart]FourteenRobot, 0),
		XSize: xSize,
		YSize: ySize,
		Time: 0,
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
			SpeedX: sx,
			SpeedY: sy,	
		}
	}
	return out
}

type FourteenBathroomMap struct {
	Robots map[FourteenRobotStart]FourteenRobot
	XSize int
	YSize int
	Time int
}

func (b *FourteenBathroomMap) AdvanceTime(seconds int) {
	b.Time += seconds
	for k, r := range b.Robots {
		pos := k.Position.GetCoordAt(b.Time * r.SpeedX, b.Time * r.SpeedY)
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

func (b *FourteenBathroomMap) GetRobotsPerQuadrant() (int, int, int, int) {
	q0, q1, q2, q3 := 0, 0,0,0
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

type FourteenRobot struct {
	Position Coords
	SpeedX int
	SpeedY int
}

type FourteenRobotStart struct {
	Position Coords
	Index int
}
