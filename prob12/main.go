package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Data processing
	processedData := strings.Split(string(data), "\n")

	type command struct {
		instruction string
		value       int
	}

	commands := []command{}
	for _, line := range processedData {
		letter := string(line[0:1])
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		commands = append(commands, command{letter, value})
	}
	x := 0
	y := 0
	currentAngle := 0
	for _, command := range commands {
		switch command.instruction {
		case "N":
			y += command.value
		case "S":
			y -= command.value
		case "E":
			x += command.value
		case "W":
			x -= command.value
		case "L":
			incrementAngle(&currentAngle, -command.value)
		case "R":
			incrementAngle(&currentAngle, command.value)
		case "F":
			switch currentAngle {
			case 0:
				x += command.value
			case 90:
				y -= command.value
			case 180:
				x -= command.value
			case 270:
				y += command.value
			default:
				fmt.Println("WTF?", currentAngle)
			}
		default:
			fmt.Println("WTF?", command.instruction)
		}
	}
	fmt.Println(abs(x) + abs(y))

	x = 0
	y = 0
	xWP := 10
	yWP := 1
	for _, command := range commands {

		switch command.instruction {
		case "N":
			yWP += command.value
		case "S":
			yWP -= command.value
		case "E":
			xWP += command.value
		case "W":
			xWP -= command.value
		case "L":
			rotateWaypoint(&xWP, &yWP, -command.value)
		case "R":
			rotateWaypoint(&xWP, &yWP, command.value)
		case "F":
			x += (xWP * command.value)
			y += (yWP * command.value)
		}

	}

	fmt.Println(abs(x) + abs(y))
}

// Clockwise is positive
func incrementAngle(currentAngle *int, angle int) {
	*currentAngle += angle
	if *currentAngle >= 360 {
		*currentAngle -= 360
	}
	if *currentAngle < 0 {
		*currentAngle += 360
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func rotateWaypoint(xWP, yWP *int, angleIncrement int) {
	oldX := *xWP
	oldY := *yWP
	if angleIncrement < 0 {
		angleIncrement += 360
	}
	if angleIncrement >= 360 {
		angleIncrement -= 360
	}
	switch angleIncrement {
	case 90:
		*xWP = oldY
		*yWP = -oldX
	case 180:
		*xWP = -oldX
		*yWP = -oldY
	case 270:
		*xWP = -oldY
		*yWP = oldX
	case 0:
		// Do Nothing
	default:
		fmt.Println("weird angle", angleIncrement)
	}

}
