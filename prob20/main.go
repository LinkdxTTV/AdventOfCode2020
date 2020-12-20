package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type gridObject struct {
	literalGrid [][]string
	edges       [][]string
	edgesMap    map[string]int
}

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Data processing
	processedData := strings.Split(string(data), "\n")
	gridMap := map[int]gridObject{}
	grid := [][]string{}
	id := 0
	for _, line := range processedData {
		if line == "" { // Save this grid and start New Grid
			gridMap[id] = gridObject{
				literalGrid: grid,
			}
			grid = [][]string{}
			continue
		}
		if line[0:4] == "Tile" {
			number := strings.Split(line, " ")
			idInt, err := strconv.Atoi(number[1][0:4])
			if err != nil {
				panic(err)
			}
			id = idInt
			continue
		}
		newLine := []string{}
		for _, char := range line {
			newLine = append(newLine, string(char))
		}
		grid = append(grid, newLine)
	}

	realGridMap := map[int]gridObject{}
	for key, gridObj := range gridMap {
		rightEdge := returnRightEdge(gridObj.literalGrid)
		bottomEdge := returnRightEdge(rotate(gridObj.literalGrid))
		leftEdge := returnRightEdge(rotate(rotate(gridObj.literalGrid)))
		topEdge := returnRightEdge(rotate(rotate(rotate(gridObj.literalGrid))))
		edges := [][]string{rightEdge, bottomEdge, leftEdge, topEdge}

		edgesMap := map[string]int{}
		for _, edge := range edges {
			edgeStr := strings.Join(edge, "")
			_, alreadyExists := edgesMap[edgeStr]
			if !alreadyExists {
				edgeStrFlip := strings.Join(flippedEdge(edge), "")
				_, alreadyExistsFlipped := edgesMap[edgeStrFlip]
				if !alreadyExistsFlipped {
					edgesMap[edgeStr] = 1
				} else {
					edgesMap[edgeStrFlip]++
				}
			} else {
				edgesMap[edgeStr]++
			}
		}

		realGridMap[key] = gridObject{
			gridObj.literalGrid,
			edges,
			edgesMap,
		}
	}

	edgeMap := map[string][]int{}
	for matchingGridKey, gridObj := range realGridMap {
		for edgeStr, _ := range gridObj.edgesMap {
			_, exists := edgeMap[edgeStr]
			if !exists {
				edgeStrFlipped := flipString(edgeStr)
				_, exists2 := edgeMap[edgeStrFlipped]
				if !exists2 {
					edgeMap[edgeStr] = []int{matchingGridKey}
				} else {
					edgeMap[edgeStrFlipped] = append(edgeMap[edgeStrFlipped], matchingGridKey)
				}
			} else {
				edgeMap[edgeStr] = append(edgeMap[edgeStr], matchingGridKey)
			}
		}
	}

	// Run through every grid object and find the ones that have edges that have value 1
	// Find the four grid objects that have 2 objects with value 1
	corners := []int{}
	for key, grid := range realGridMap {
		uniqueEdges := 0
		// fmt.Println(len(grid.edges))
		for edgeStr, edgeVal := range grid.edgesMap {
			val, ok := edgeMap[edgeStr]
			if ok {
				// fmt.Println(key, val)
				if len(val) == 1 {
					uniqueEdges += edgeVal
				}
			}
		}
		if uniqueEdges == 2 {
			corners = append(corners, key)
		}
	}
	fmt.Println(corners)

	product := 1
	for _, num := range corners {
		product = product * num
	}
	fmt.Println("Part1", product)

	realRealGridMap := map[int][][]string{}
	// PART 2 PROCESSING STARTS HERE
	topLeft := corners[0]

	topLeftGrid := gridMap[topLeft]

	rotatedCorrectly := false
	gridLiteral := topLeftGrid.literalGrid
	for !rotatedCorrectly {
		topEdgeStr := strings.Join(returnTopEdge(gridLiteral), "")
		leftEdgeStr := strings.Join(returnLeftEdge(gridLiteral), "")

		if (len(edgeMap[topEdgeStr]) == 1 || len(edgeMap[flipString(topEdgeStr)]) == 1) && (len(edgeMap[leftEdgeStr]) == 1 || len(edgeMap[flipString(leftEdgeStr)]) == 1) {
			rotatedCorrectly = true
			realRealGridMap[topLeft] = gridLiteral
			break
		}
		gridLiteral = rotate(gridLiteral)
	}

	// fmt.Println(realRealGridMap)

	correctKeys := [][]int{}

	// Heh
	n := math.Sqrt(float64(len(realGridMap)))

	// Init
	for i := 0; i < int(n); i++ {
		row := []int{}
		for j := 0; j < int(n); j++ {
			row = append(row, 0)
		}
		correctKeys = append(correctKeys, row)
	}
	correctKeys[0][0] = topLeft

	var currentGridObjectKey int
	var currentGridObject [][]string

	for row := 1; row < int(n)+1; row++ {
		for i := 1; i < int(n); i++ {
			currentGridObjectKey = correctKeys[row-1][i-1]
			currentGridObject = realRealGridMap[currentGridObjectKey]
			correctRightEdge := returnRightEdge(currentGridObject)
			correctRightEdgeStr := strings.Join(correctRightEdge, "")

			val, ok := edgeMap[correctRightEdgeStr]
			if !ok {
				val, ok = edgeMap[flipString(correctRightEdgeStr)]
			}
			var nextKey int
			for _, key := range val {
				if len(val) > 2 {
					panic("wtf")
				}
				if key == currentGridObjectKey {
					continue
				}
				nextKey = key
				break
			}

			possibleGrid := realGridMap[nextKey].literalGrid
			correctlyRotated := false
			rotationCounter := 0

			for !correctlyRotated {
				if correctRightEdgeStr == strings.Join(returnLeftEdge(possibleGrid), "") {
					correctlyRotated = true
					break
				}

				possibleGrid = rotate(possibleGrid)
				rotationCounter++
				if rotationCounter%4 == 0 {
					possibleGrid = flip(possibleGrid)
				}
			}

			realRealGridMap[nextKey] = possibleGrid
			correctKeys[row-1][i] = nextKey
			currentGridObjectKey = nextKey
			currentGridObject = possibleGrid
		}
		if row == int(n) {
			break
		}
		// Find the one that matches on the top
		correctBottomEdge := returnBottomEdge(realRealGridMap[correctKeys[row-1][0]])
		correctBottomEdgeStr := strings.Join(correctBottomEdge, "")
		val, ok := edgeMap[correctBottomEdgeStr]
		if !ok {
			val, ok = edgeMap[flipString(correctBottomEdgeStr)]
		}
		var nextKey int
		for _, key := range val {
			if len(val) > 2 {
				panic("wtf")
			}
			if key == correctKeys[row-1][0] {
				continue
			}
			nextKey = key
			break
		}

		possibleGrid := realGridMap[nextKey].literalGrid
		correctlyRotated := false
		rotationCounter := 0
		for !correctlyRotated {
			if correctBottomEdgeStr == strings.Join(returnTopEdge(possibleGrid), "") {
				correctlyRotated = true
				break
			}

			possibleGrid = rotate(possibleGrid)
			rotationCounter++
			if rotationCounter%4 == 0 {
				possibleGrid = flip(possibleGrid)
			}
		}

		correctKeys[row][0] = nextKey
		realRealGridMap[nextKey] = possibleGrid
		currentGridObjectKey = nextKey
		currentGridObject = possibleGrid
	}

	fmt.Println(correctKeys)

	shavedGridMap := map[int][][]string{}
	for id, gridObj := range realRealGridMap {
		shavedGridMap[id] = shaveGridBorders(gridObj)
	}

	// Combine subgrids into gigaGrid
	superGrid := []string{}
	realSize := len(shavedGridMap[correctKeys[0][0]])
	// Initialize
	for i := 0; i < int(n)*realSize; i++ {
		row := ""
		for j := 0; j < int(n)*realSize; j++ {
			row += "0"
		}
		superGrid = append(superGrid, row)
	}

	fmt.Println(len(superGrid), len(superGrid[0]))

	// Fill supergrid

	for i := 0; i < int(n)*realSize; i++ {
		for j := 0; j < int(n)*realSize; j++ {
			// Find correct string to put in
			index1 := i / realSize
			remainder1 := i % realSize
			index2 := j / realSize
			remainder2 := j % realSize
			key := correctKeys[index1][index2]
			char := shavedGridMap[key][remainder1][remainder2]

			replacementString := superGrid[i][0:j] + char + superGrid[i][j+1:]
			superGrid[i] = replacementString
		}
	}
	fmt.Println(len(superGrid), len(superGrid[0]))
	for _, row := range superGrid {
		fmt.Println(row)
	}
	seaMonsters := []string{}
	for i := 0; i < int(n)*realSize; i++ {
		row := ""
		for j := 0; j < int(n)*realSize; j++ {
			row += "."
		}
		seaMonsters = append(seaMonsters, row)
	}
	// Look for sea monsters

	type seaMonsterCoordinate struct {
		x int
		y int
	}

	seaMonsterCoordinates := []seaMonsterCoordinate{
		{0, 0}, {0, 5}, {0, 6}, {0, 11}, {0, 12}, {0, 17}, {0, 18}, {0, 19}, {-1, 18}, {1, 1}, {1, 4}, {1, 7}, {1, 10}, {1, 13}, {1, 16},
	}

	seaMonsterCounter := 0
	rotationCounter := 0
	for seaMonsterCounter == 0 {

		for idx1, row := range superGrid {
			if idx1 == 0 || idx1 == len(superGrid)-1 {
				continue
			}
			for idx2, _ := range row {
				if idx2+19 > len(row) {
					continue
				}
				seaMonsterFound := true
				for _, coordinate := range seaMonsterCoordinates {
					if !(string(superGrid[idx1+coordinate.x][idx2+coordinate.y]) == "#") {
						seaMonsterFound = false
						break
					}
				}
				if seaMonsterFound {
					seaMonsterCounter++
					// Fill in output grid
					for _, coordinate := range seaMonsterCoordinates {
						seaMonsters[idx1+coordinate.x] = seaMonsters[idx1+coordinate.x][0:idx2+coordinate.y] + "0" + seaMonsters[idx1+coordinate.x][idx2+coordinate.y+1:]
					}
				}
			}
		}
		superGrid = rotateImage(superGrid)
		rotationCounter++
		if rotationCounter%4 == 0 {
			superGrid = flipImage(superGrid)
		}
		if rotationCounter == 8 {
			panic("sad")
		}
	}
	fmt.Println(seaMonsterCounter)

	for _, row := range seaMonsters {
		fmt.Println(row)
	}

	poundCount := 0
	zeroCount := 0
	for _, row := range superGrid {
		for _, char := range row {
			if string(char) == "#" {
				poundCount++
			}
		}
	}
	for _, row := range seaMonsters {
		for _, char := range row {
			if string(char) == "0" {
				zeroCount++
			}
		}
	}
	fmt.Println(poundCount, zeroCount, "ans", poundCount-zeroCount)

}

// Rotates counter clockwise
func rotate(grid [][]string) [][]string {
	output := [][]string{}

	//Init output
	for i := 0; i < len(grid[0]); i++ {
		newRow := []string{}
		for j := 0; j < len(grid); j++ {
			newRow = append(newRow, "0")
		}
		output = append(output, newRow)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			output[len(grid[0])-1-j][i] = grid[i][j]
		}
	}
	return output
}

func flip(grid [][]string) [][]string {
	//Init output
	output := [][]string{}
	for i := 0; i < len(grid); i++ {
		newRow := []string{}
		for j := 0; j < len(grid[0]); j++ {
			newRow = append(newRow, "0")
		}
		output = append(output, newRow)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			output[i][len(grid[0])-1-j] = grid[i][j]
		}
	}
	return output
}

func flippedEdge(edge []string) []string {
	output := []string{}
	for i := 0; i < len(edge); i++ {
		output = append(output, string(edge[len(edge)-1-i]))
	}
	return output
}

func flipString(input string) string {
	var output string
	for i := 0; i < len(input); i++ {
		output += string(input[len(input)-1-i])
	}
	return output
}

// Edge shit

func returnRightEdge(grid1 [][]string) []string {
	// Right Edge
	rightEdgeIndex := len(grid1[0]) - 1
	edge1 := []string{}
	for _, row := range grid1 {
		edge1 = append(edge1, string(row[rightEdgeIndex]))
	}
	return edge1
}

func returnTopEdge(grid [][]string) []string {
	return grid[0]
}

func returnLeftEdge(grid [][]string) []string {

	edge1 := []string{}
	for _, row := range grid {
		edge1 = append(edge1, string(row[0]))
	}
	return edge1
}

func returnBottomEdge(grid [][]string) []string {
	return grid[len(grid)-1]
}

func shaveGridBorders(grid [][]string) [][]string {
	output := [][]string{}

	for jdx, row := range grid {
		if jdx == 0 || jdx == len(grid)-1 {
			continue
		}
		newRow := []string{}
		for idx, char := range row {
			if idx == 0 || idx == len(row)-1 {
				continue
			}
			newRow = append(newRow, char)
		}
		output = append(output, newRow)
	}
	return output
}

// Image shit

// Rotates counterclockwise
func rotateImage(image []string) []string {
	output := []string{}
	for i := 0; i < len(image); i++ {
		newStr := ""
		for j := 0; j < len(image[0]); j++ {
			newStr += string(image[j][len(image)-1-i])
		}
		output = append(output, newStr)
	}
	return output
}

func flipImage(image []string) []string {
	output := []string{}
	for i := 0; i < len(image); i++ {
		newStr := ""
		for j := 0; j < len(image[0]); j++ {
			newStr += string(image[i][len(image)-1-j])
		}
		output = append(output, newStr)
	}
	return output
}
