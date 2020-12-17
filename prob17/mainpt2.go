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

	FourDimCube := [][][][]string{} // Fuck hypers
	// [x][y][z]

	// Initialize
	for h := 0; h < len(processedData); h++ {
		hRow := [][][]string{}
		for i := 0; i < len(processedData); i++ {
			iRow := [][]string{}
			for j := 0; j < len(processedData); j++ {
				jRow := []string{}
				for k := 0; k < len(processedData); k++ {
					jRow = append(jRow, ".")
				}
				iRow = append(iRow, jRow)
			}
			hRow = append(hRow, iRow)
		}
		FourDimCube = append(FourDimCube, hRow)
	}

	// fmt.Println(ThreeDimCube)
	// Starting Data
	for index, line := range processedData {
		for i := 0; i < len(line); i++ {
			FourDimCube[0][i][index][1] = string(line[i])
		}
	}

	cubeOriginal := expandCube(FourDimCube)
	for iteration := 0; iteration < 6; iteration++ {
		cubeCopy := createDeepCopy(cubeOriginal)
		for h := 0; h < len(cubeOriginal); h++ {
			for i := 0; i < len(cubeOriginal); i++ {
				for j := 0; j < len(cubeOriginal); j++ {
					for k := 0; k < len(cubeOriginal); k++ {
						activeNeighors := checkNeighborsActivity(h, i, j, k, cubeOriginal)
						if cubeOriginal[h][i][j][k] == "." {
							if activeNeighors == 3 {
								cubeCopy[h][i][j][k] = "#"
							}
						} else if cubeOriginal[h][i][j][k] == "#" {
							if activeNeighors == 2 || activeNeighors == 3 {
								// Stay active
							} else {
								cubeCopy[h][i][j][k] = "."
							}
						}
					}
				}
			}
		}
		cubeOriginal = expandCube(cubeCopy)
	}

	active := 0
	for h := 0; h < len(cubeOriginal); h++ {
		for i := 0; i < len(cubeOriginal); i++ {
			for j := 0; j < len(cubeOriginal); j++ {
				for k := 0; k < len(cubeOriginal); k++ {
					if cubeOriginal[h][i][j][k] == "#" {
						active++
					}
				}
			}
		}
	}
	fmt.Println(active)
}

func createDeepCopy(input [][][][]string) [][][][]string {
	output := [][][][]string{}
	for h := 0; h < len(input); h++ {
		hRow := [][][]string{}
		for i := 0; i < len(input); i++ {
			iRow := [][]string{}
			for j := 0; j < len(input); j++ {
				jRow := []string{}
				for k := 0; k < len(input); k++ {
					jRow = append(jRow, input[h][i][j][k])
				}
				iRow = append(iRow, jRow)
			}
			hRow = append(hRow, iRow)
		}
		output = append(output, hRow)
	}
	return output
}

func checkNeighborsActivity(w, x, y, z int, cube [][][][]string) int {
	neighborsActive := 0
	rows := len(cube)
	for h := -1; h <= 1; h++ {
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				for k := -1; k <= 1; k++ {
					if x+i < 0 || y+j < 0 || z+k < 0 || w+h < 0 {
						continue
					}
					if x+i >= rows || y+j >= rows || z+k >= rows || w+h >= rows {
						continue
					}
					if i == 0 && j == 0 && k == 0 && h == 0 {
						continue
					}
					if cube[w+h][x+i][y+j][z+k] == "#" {
						neighborsActive++
					}
				}
			}
		}
	}
	return neighborsActive
}

func expandCube(input [][][][]string) [][][][]string {
	currentSize := len(input)
	output := [][][][]string{}

	newBlankSpace := [][][]string{}
	for spaceExpand := 0; spaceExpand < currentSize+2; spaceExpand++ {
		newBlankPlane := [][]string{}
		for planeExpand := 0; planeExpand < currentSize+2; planeExpand++ {
			newBlankRow := []string{}
			for iExpand := 0; iExpand < currentSize+2; iExpand++ {
				newBlankRow = append(newBlankRow, ".")
			}
			newBlankPlane = append(newBlankPlane, newBlankRow)
		}
		newBlankSpace = append(newBlankSpace, newBlankPlane)
	}
	output = append(output, newBlankSpace)

	for h := 0; h < len(input); h++ {
		hSpace := [][][]string{}
		newBlankPlane := [][]string{}
		for planeExpand := 0; planeExpand < currentSize+2; planeExpand++ {
			newBlankRow := []string{}
			for iExpand := 0; iExpand < currentSize+2; iExpand++ {
				newBlankRow = append(newBlankRow, ".")
			}
			newBlankPlane = append(newBlankPlane, newBlankRow)
		}
		hSpace = append(hSpace, newBlankPlane)

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
					jRow = append(jRow, input[h][i][j][k])
				}
				jRow = append(jRow, ".")
				iRow = append(iRow, jRow)

			}
			newBlankRow = []string{}
			for iExpand := 0; iExpand < currentSize+2; iExpand++ {
				newBlankRow = append(newBlankRow, ".")
			}
			iRow = append(iRow, newBlankRow)
			hSpace = append(hSpace, iRow)
		}
		newBlankPlane = [][]string{}
		for planeExpand := 0; planeExpand < currentSize+2; planeExpand++ {
			newBlankRow := []string{}
			for iExpand := 0; iExpand < currentSize+2; iExpand++ {
				newBlankRow = append(newBlankRow, ".")
			}
			newBlankPlane = append(newBlankPlane, newBlankRow)
		}
		hSpace = append(hSpace, newBlankPlane)
		output = append(output, hSpace)
	}

	newBlankSpace = [][][]string{}
	for spaceExpand := 0; spaceExpand < currentSize+2; spaceExpand++ {
		newBlankPlane := [][]string{}
		for planeExpand := 0; planeExpand < currentSize+2; planeExpand++ {
			newBlankRow := []string{}
			for iExpand := 0; iExpand < currentSize+2; iExpand++ {
				newBlankRow = append(newBlankRow, ".")
			}
			newBlankPlane = append(newBlankPlane, newBlankRow)
		}
		newBlankSpace = append(newBlankSpace, newBlankPlane)
	}
	output = append(output, newBlankSpace)

	return output
}
