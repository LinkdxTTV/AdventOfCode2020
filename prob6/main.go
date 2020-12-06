package main

import (
	"fmt"
	"io/ioutil"
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
	processedData = append(processedData, "") // Trailing whitespace to fix my shitty logic
	response := map[string]int{}
	groups := []map[string]int{}
	peoplePerGroup := []int{}
	numPeople := 0
	for _, line := range processedData {
		if line == "" {
			groups = append(groups, response)
			peoplePerGroup = append(peoplePerGroup, numPeople)
			response = map[string]int{}
			numPeople = 0
			continue
		}
		for _, letter := range line {
			response[string(letter)]++
		}

		numPeople++
	}
	count := 0
	for _, group := range groups {
		count += len(group)
	}
	fmt.Println("Part 1: ", count)

	count = 0
	for i, group := range groups {
		for _, v := range group {
			if v == peoplePerGroup[i] {
				count++
			}
		}
	}
	fmt.Println("Part 2: ", count)
}
