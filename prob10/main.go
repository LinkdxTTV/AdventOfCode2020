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
	joltDifferences := map[int]int{}
	intMap := map[int]bool{}
	intArray := []int{}
	for _, num := range processedData {
		numInt, err := strconv.Atoi(string(num))
		fmt.Println(numInt)
		if err != nil {
			fmt.Println(err)
			return
		}
		intMap[numInt] = true
		intArray = append(intArray, numInt)
	}

	_, max := findMinAndMax(intArray)
	intMap[max+3] = true
	intArray = append(intArray, max+3)
	currentJolt := 0
	for i := 0; i <= max+3; i++ {
		_, ok := intMap[i]
		if !ok {
			fmt.Println("didnt see ", i)
			continue
		}
		diff := i - currentJolt
		currentJolt += diff
		if diff > 3 {
			return
		}
		joltDifferences[diff]++
	}

	fmt.Println(joltDifferences)
	fmt.Println(joltDifferences[1] * joltDifferences[3])

	memo := map[int]int{}
	recursiveFindWithMemo(0, max+3, intMap, memo)
	fmt.Println("valid", memo[0])

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

func recursiveFindWithMemo(start, target int, intMap map[int]bool, memo map[int]int) int {
	if start == target {
		return 1
	}
	value, ok := memo[start]
	if ok {
		return value
	}
	children := 0
	for i := start + 1; i <= start+3; i++ {
		_, ok := intMap[i]
		if ok {
			children += recursiveFindWithMemo(i, target, intMap, memo)
		}
	}
	memo[start] = children
	return children
}
