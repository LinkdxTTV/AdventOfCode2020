package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("~(UwU)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	processedData := strings.Split(string(data), "\n")

	ids := []int64{}
	idmap := map[int64]bool{}
	for _, data := range processedData {
		binaryRow := ""
		binaryColumn := ""
		for _, letter := range data {
			letterStr := string(letter)
			switch letterStr {
			case "B":
				binaryRow += "1"
			case "F":
				binaryRow += "0"
			case "L":
				binaryColumn += "0"
			case "R":
				binaryColumn += "1"
			}
		}

		row, _ := strconv.ParseInt(binaryRow, 2, 64)
		column, _ := strconv.ParseInt(binaryColumn, 2, 64)
		id := row*8 + column
		ids = append(ids, id)
		idmap[id] = true
	}
	min, max := findMinAndMax(ids)
	fmt.Println("min ID: ", min, "||", "max ID: ", max)
	// Find Missing ID in List
	for i := min; i <= max; i++ {
		_, ok := idmap[i]
		if !ok {
			fmt.Println("Missing Seat in List: ", i)
		}
	}
}

// Stole this from the internet because Go for some reason doesnt have this functionality natively thank you google
func findMinAndMax(a []int64) (min int64, max int64) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
