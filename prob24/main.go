package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Data processing
	processedData := strings.Split(string(data), "\n")

	directionsList := [][]string{}
	for _, line := range processedData {
		directions := []string{}
		idx := 0
		for len(line) > 0 {
			if string(line[idx]) == "n" || string(line[idx]) == "s" {
				if string(line[idx:idx+2]) == "ne" || string(line[idx:idx+2]) == "se" || string(line[idx:idx+2]) == "nw" || string(line[idx:idx+2]) == "sw" {
					directions = append(directions, line[idx:idx+2])
					line = line[idx+2:]
				}
			} else if string(line[idx]) == "w" || string(line[idx]) == "e" {
				directions = append(directions, string(line[idx]))
				line = line[idx+1:]
			} else {
				panic("monkaS")
			}
		}
		directionsList = append(directionsList, directions)
	}

	tiles := map[coordinate]string{}

	for _, directions := range directionsList {
		newCoordinate := &coordinate{0, 0}
		for _, direction := range directions {
			switch direction {
			case "se":
				newCoordinate.x += 1
				newCoordinate.y += -1
			case "ne":
				newCoordinate.x += 0
				newCoordinate.y += 1
			case "nw":
				newCoordinate.x += -1
				newCoordinate.y += 1
			case "sw":
				newCoordinate.x += 0
				newCoordinate.y += -1
			case "e":
				newCoordinate.x += 1
				newCoordinate.y += 0
			case "w":
				newCoordinate.x += -1
				newCoordinate.y += 0
			}
		}
		// fmt.Println(newCoordinate)
		color, placed := tiles[*newCoordinate]
		if !placed {
			tiles[*newCoordinate] = "black"
		} else {
			if color == "black" {
				tiles[*newCoordinate] = "white"
			} else {
				tiles[*newCoordinate] = "black"
			}
		}
	}
	blackTiles := 0
	for _, color := range tiles {
		if color == "black" {
			blackTiles++
		}
	}
	fmt.Println(blackTiles)

	// Part 2
	// Casually populate a bunch more tiles
	for i := -100; i <= 100; i++ {
		for j := -100; j <= 100; j++ {
			_, check := tiles[coordinate{i, j}]
			if !check {
				tiles[coordinate{i, j}] = "white"
			}
		}
	}

	for day := 1; day <= 100; day++ {
		newTiles := deepCopyMap(tiles)
		for coord, color := range tiles {
			adjacentBlackTiles := checkAdjacency(tiles, coord)
			if color == "black" {
				if adjacentBlackTiles == 0 || adjacentBlackTiles > 2 {
					newTiles[coord] = "white"
				}
			}
			if color == "white" {
				if adjacentBlackTiles == 2 {
					newTiles[coord] = "black"
				}
			}
		}
		tiles = newTiles
	}
	blackTiles = 0
	for _, color := range tiles {
		if color == "black" {
			blackTiles++
		}
	}
	fmt.Println(blackTiles)
}

func checkAdjacency(blackTiles map[coordinate]string, coord coordinate) int {
	adjacent := 0
	coordinateShift := []coordinate{{-1, 1}, {0, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, 0}}
	for _, shift := range coordinateShift {
		color := checkCoordinate(blackTiles, coordinate{coord.x + shift.x, coord.y + shift.y})
		if color == "black" {
			adjacent++
		}
	}
	return adjacent
}

func checkCoordinate(blackTiles map[coordinate]string, coordinate coordinate) string {
	color, ok := blackTiles[coordinate]
	if !ok {
		return "white"
	} else {
		return color
	}
}

func deepCopyMap(input map[coordinate]string) map[coordinate]string {
	output := map[coordinate]string{}
	for k, v := range input {
		output[k] = v
	}
	return output
}
