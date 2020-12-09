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
	processedData := strings.Split(string(data), "\n")
	processedInts := []int{}
	for _, num := range processedData {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println(err)
		}
		processedInts = append(processedInts, numInt)
	}

	for idx, num := range processedInts {

		if idx < 25 {
			continue
		}
		check := twoSumInPastEntries(processedInts[(idx-25):idx], num)
		if !check {
			fmt.Println("Part 1: ", num)
			break
		}
	}

	target := 26134589

	found, left, right := windowFind(processedInts, target)
	if found {
		min, max := findMinAndMax(processedInts[left : right+1])
		fmt.Println("Part 2: ", min+max)
	}

}

func twoSumInPastEntries(array []int, sum int) bool {
	seen := map[int]bool{}
	for _, num := range array {
		complement := sum - num
		if num == complement {
			continue
		}
		if _, ok := seen[complement]; ok {
			return true
		}
		seen[num] = true
	}

	return false
}

func windowFind(array []int, target int) (bool, int, int) {
	windowLeft := 0
	windowRight := 0
	sum := 0
	for {
		if windowRight > len(array) {
			fmt.Println("didnt find any window")
			return false, 0, 0
		}
		if sum < target {
			sum += array[windowRight]
			windowRight++
			continue
		}
		if sum > target {
			sum -= array[windowLeft]
			windowLeft++
			continue
		}
		if sum == target {
			return true, windowLeft, windowRight
		}
	}
}

// Stole this from the internet because Go for some reason doesnt have this functionality natively thank you google
func findMinAndMax(a []int) (min int, max int) {
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
