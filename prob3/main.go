package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("OwO")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	processedData := strings.Split(string(data), "\n")
	// Part 1
	trees := lookForTrees(processedData, 3, 1)
	fmt.Println(trees)

	// Part 2
	treeArray := []int{}

	type incPairs struct {
		xInt int
		yInt int
	}
	incrementArray := []incPairs{
		{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2},
	}
	for _, incPair := range incrementArray {
		treeArray = append(treeArray, lookForTrees(processedData, incPair.xInt, incPair.yInt))
	}
	totalProduct := 1
	for _, treesProd := range treeArray {
		totalProduct = totalProduct * treesProd
	}
	fmt.Println(totalProduct)
}

func lookForTrees(data []string, xInc, yInc int) int {
	x := 0
	y := 0
	trees := 0
	for {
		if y >= len(data) {
			break
		}
		if x > len(data[0])-1 {
			x = x - len(data[0])
		}
		if string(data[y][x]) == "#" {
			trees++
		}
		x += xInc
		y += yInc
	}
	return trees
}

// 90 274 82 68 44
