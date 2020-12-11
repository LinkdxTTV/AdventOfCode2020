package main

import (
	"fmt"
	"io/ioutil"
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
	fullGrid := [][]string{}
	for _, line := range processedData {
		lineString := []string{}
		for _, char := range line {
			lineString = append(lineString, string(char))
		}
		fullGrid = append(fullGrid, lineString)
	}
	output := createDataCopy(fullGrid)

	// Part 1
	for {
		input := createDataCopy(output)
		output = goThroughGridOnce(input)
		equal := areSlicesAreEqual(input, output)
		if equal {
			break
		}
	}

	count := countOccupiedSeats(output)
	fmt.Println(count)

	// Part 2
	fmt.Println("part 2")
	output = createDataCopy(fullGrid)

	for {
		input := createDataCopy(output)
		output = goThroughGridOncePart2(input)
		equal := areSlicesAreEqual(input, output)
		if equal {
			break
		}

	}

	count = countOccupiedSeats(output)
	fmt.Println(count)

}

func countOccupiedSeats(data [][]string) int {
	count := 0
	for _, line := range data {
		for _, char := range line {
			if char == "#" {
				count++
			}
		}
	}
	return count
}

func areSlicesAreEqual(one, two [][]string) bool {
	for i := range one {
		for j := range one[0] {
			if one[i][j] != two[i][j] {
				return false
			}
		}
	}
	return true
}

func adjacentSeats(x, y int, data [][]string) int {

	count := 0
	for xTest := x - 1; xTest <= x+1; xTest++ {
		for yTest := y - 1; yTest <= y+1; yTest++ {
			if xTest == x && y == yTest {
				continue
			}
			if xTest < 0 || yTest < 0 || xTest >= len(data[0]) || yTest >= len(data) {
				continue
			}
			if data[yTest][xTest] == "#" {
				count++
			}
		}
	}
	return count
}

func createDataCopy(input [][]string) [][]string {
	output := [][]string{}
	for i := range input {
		slice := make([]string, len(input[0]))
		copy(slice, input[i])
		output = append(output, slice)
	}
	return output
}

func goThroughGridOnce(input [][]string) [][]string {
	output := createDataCopy(input)
	for y, xList := range input {
		for x, char := range xList {
			if char == "." {
				continue
			}
			adjSeats := adjacentSeats(x, y, input)

			if adjSeats == 0 && char == "L" {
				output[y][x] = "#"
			}
			if char == "#" && adjSeats >= 4 {
				output[y][x] = "L"
			}

		}
	}
	return output
}

func goThroughGridOncePart2(input [][]string) [][]string {
	output := createDataCopy(input)
	for y, xList := range input {
		for x, char := range xList {
			if char == "." {
				continue
			}
			adjSeats := adjacentSeatsPart2(x, y, input)

			if adjSeats == 0 && char == "L" {
				output[y][x] = "#"
			}
			if char == "#" && adjSeats >= 5 {
				output[y][x] = "L"
			}

		}
	}
	return output
}

// Define sight directions
type sight struct {
	xInc int
	yInc int
}

func adjacentSeatsPart2(x, y int, data [][]string) int {
	// Starting at up and going clockwise
	lookingDirections := []sight{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}
	count := 0
	for _, direction := range lookingDirections {
		xTest, yTest := x, y
		for xTest >= 0 && xTest < len(data[0]) && yTest >= 0 && yTest < len(data) {

			if x == xTest && y == yTest {
				xTest, yTest = xTest+direction.xInc, yTest+direction.yInc
				continue
			}

			if data[yTest][xTest] == "L" || data[yTest][xTest] == "#" {
				if data[yTest][xTest] == "#" {
					count++
				}
				break
			}
			xTest, yTest = xTest+direction.xInc, yTest+direction.yInc
		}
	}
	return count
}
