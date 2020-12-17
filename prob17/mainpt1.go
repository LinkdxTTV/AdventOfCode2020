package main1

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

	ThreeDimCube := [][][]string{} // Fuck
	// [x][y][z]

	// Initialize
	for i := 0; i < len(processedData); i++ {
		iRow := [][]string{}
		for j := 0; j < len(processedData); j++ {
			jRow := []string{}
			for k := 0; k < len(processedData); k++ {
				jRow = append(jRow, ".")
			}
			iRow = append(iRow, jRow)
		}
		ThreeDimCube = append(ThreeDimCube, iRow)
	}

	// fmt.Println(ThreeDimCube)
	// Starting Data
	for index, line := range processedData {
		for i := 0; i < len(line); i++ {
			ThreeDimCube[i][index][1] = string(line[i])
		}
	}

	cubeOriginal := expandCube(ThreeDimCube)
	for iteration := 0; iteration < 6; iteration++ {
		cubeCopy := createDeepCopy(cubeOriginal)

		for i := 0; i < len(cubeOriginal); i++ {
			for j := 0; j < len(cubeOriginal); j++ {
				for k := 0; k < len(cubeOriginal); k++ {
					activeNeighors := checkNeighborsActivity(i, j, k, cubeOriginal)
					if cubeOriginal[i][j][k] == "." {
						if activeNeighors == 3 {
							cubeCopy[i][j][k] = "#"
						}
					} else if cubeOriginal[i][j][k] == "#" {
						if activeNeighors == 2 || activeNeighors == 3 {
							// Stay active
						} else {
							cubeCopy[i][j][k] = "."
						}
					}
				}
			}
		}
		cubeOriginal = expandCube(cubeCopy)
	}

	active := 0
	for i := 0; i < len(cubeOriginal); i++ {
		for j := 0; j < len(cubeOriginal); j++ {
			for k := 0; k < len(cubeOriginal); k++ {
				if cubeOriginal[i][j][k] == "#" {
					active++
				}
			}
		}
	}
	fmt.Println(active)
}

func createDeepCopy(input [][][]string) [][][]string {
	output := [][][]string{}
	for i := 0; i < len(input); i++ {
		iRow := [][]string{}
		for j := 0; j < len(input); j++ {
			jRow := []string{}
			for k := 0; k < len(input); k++ {
				jRow = append(jRow, input[i][j][k])
			}
			iRow = append(iRow, jRow)
		}
		output = append(output, iRow)
	}
	return output
}

func checkNeighborsActivity(x, y, z int, cube [][][]string) int {
	neighborsActive := 0
	rows := len(cube)
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			for k := -1; k <= 1; k++ {
				if x+i < 0 || y+j < 0 || z+k < 0 {
					continue
				}
				if x+i >= rows || y+j >= rows || z+k >= rows {
					continue
				}
				if i == 0 && j == 0 && k == 0 {
					continue
				}
				if cube[x+i][y+j][z+k] == "#" {
					neighborsActive++
				}
			}
		}
	}
	return neighborsActive
}

func expandCube(input [][][]string) [][][]string {
	currentSize := len(input)
	output := [][][]string{}

	newBlankPlane := [][]string{}
	for planeExpand := 0; planeExpand < currentSize+2; planeExpand++ {
		newBlankRow := []string{}
		for iExpand := 0; iExpand < currentSize+2; iExpand++ {
			newBlankRow = append(newBlankRow, ".")
		}
		newBlankPlane = append(newBlankPlane, newBlankRow)
	}
	output = append(output, newBlankPlane)

	for i := 0; i < len(input); i++ {
		iRow := [][]string{}
		newBlankRow := []string{}
		for iExpand := 0; iExpand < currentSize+2; iExpand++ {
			newBlankRow = append(newBlankRow, ".")
		}
		iRow = append(iRow, newBlankRow)
		for j := 0; j < len(input); j++ {
			jRow := []string{}
			jRow = append(jRow, ".")
			for k := 0; k < len(input); k++ {
				jRow = append(jRow, input[i][j][k])
			}
			jRow = append(jRow, ".")
			iRow = append(iRow, jRow)

		}
		newBlankRow = []string{}
		for iExpand := 0; iExpand < currentSize+2; iExpand++ {
			newBlankRow = append(newBlankRow, ".")
		}
		iRow = append(iRow, newBlankRow)
		output = append(output, iRow)
	}
	newBlankPlane = [][]string{}
	for planeExpand := 0; planeExpand < currentSize+2; planeExpand++ {
		newBlankRow := []string{}
		for iExpand := 0; iExpand < currentSize+2; iExpand++ {
			newBlankRow = append(newBlankRow, ".")
		}
		newBlankPlane = append(newBlankPlane, newBlankRow)
	}
	output = append(output, newBlankPlane)
	return output
}
