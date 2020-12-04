package main

import (
	"fmt"
	"strconv"
)

func main() {
	start := 273025
	end := 767253

	valid := 0
	for number := start; number <= end; number++ {
		doubleNumber := false
		ascendingOnly := false
		for i := 0; i < 5; i++ {
			if numberAtIndex(number, i) == numberAtIndex(number, i+1) {
				doubleNumber = true
				continue
			}
		}
		if !doubleNumber {
			continue
		}
		ascendingOnly = true
		for i := 0; i < 5; i++ {
			if numberAtIndex(number, i) > numberAtIndex(number, i+1) {
				ascendingOnly = false
				break
			}

		}
		if !ascendingOnly {
			continue
		}
		valid++
	}
	fmt.Println(valid)

	// Pt 2
	valid = 0
	for number := start; number <= end; number++ {
		doubleNumber := false
		ascendingOnly := false
		groupOfTwo := false
		numFreq := map[string]int{}

		for i := 0; i < 5; i++ {
			if numberAtIndex(number, i) == numberAtIndex(number, i+1) {
				numFreq[string(strconv.Itoa(number)[i])]++
				doubleNumber = true
			}
		}
		for _, freq := range numFreq {
			if freq == 1 {
				groupOfTwo = true
				break
			}
		}
		if !doubleNumber {
			continue
		}

		if !groupOfTwo {
			continue
		}
		ascendingOnly = true
		for i := 0; i < 5; i++ {
			if numberAtIndex(number, i) > numberAtIndex(number, i+1) {
				ascendingOnly = false
				break
			}

		}
		if !ascendingOnly {
			continue
		}
		valid++
	}
	fmt.Println(valid)
}

func numberAtIndex(number, i int) int {
	stringNum := strconv.Itoa(number)
	num, _ := strconv.Atoi(string(stringNum[i]))
	return num
}
