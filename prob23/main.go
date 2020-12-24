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

	// Data processing
	processedData := strings.Split(string(data), "\n")

	cupOrder := []int{}
	for _, num := range processedData[0] {
		numInt, err := strconv.Atoi(string(num))
		if err != nil {
			panic(err)
		}
		cupOrder = append(cupOrder, numInt)
	}

	cupsPart2 := deepCopyInt(cupOrder)

	// Part 1
	currentCup := cupOrder[0]
	for i := 0; i < 100; i++ {
		leftOver, pickedUp := pickUpCups(cupOrder, currentCup)
		destination := destinationCup(leftOver, currentCup)
		cupOrder = placeCupsAfterDestination(leftOver, pickedUp, destination)
		currentCup = newCurrentCup(cupOrder, currentCup)
	}

	fmt.Println("part 1 final cup order:", cupOrder)

	// Part 2
	cupsLinked, label2Index := cupLinkedList(cupsPart2)

	current := cupsLinked[0]
	for i := 0; i < 10000000; i++ {
		cup0 := current
		cup1 := cup0.next
		cup2 := cup1.next
		cup3 := cup2.next
		cup4 := cup3.next

		destinationFound := false
		destination := cup0.label - 1
		for !destinationFound {

			if destination <= 0 {
				destination = 1000000
			}
			if destination == cup1.label || destination == cup2.label || destination == cup3.label {
				destination--
				continue
			}
			destinationFound = true
		}

		destCup0 := cupsLinked[label2Index[destination]]
		destCup1 := destCup0.next

		cup0.next = cup4
		destCup0.next = cup1
		cup3.next = destCup1
		current = cup0.next
	}
	fmt.Println("part 2 cups from 1:", *cupsLinked[label2Index[1]], *cupsLinked[label2Index[1]].next, *cupsLinked[label2Index[1]].next.next)
	fmt.Println("part 2 ans:", cupsLinked[label2Index[1]].next.label*cupsLinked[label2Index[1]].next.next.label)
}

// Part 2

// lets make a cyclic linked list structure
type cup struct {
	label int
	next  *cup
}

func cupLinkedList(cups []int) ([]*cup, map[int]int) {
	output := []*cup{}
	outputMap := map[int]int{}

	for _, num := range cups {
		newCup := cup{
			label: num,
			next:  nil,
		}
		output = append(output, &newCup)
	}
	for i := len(cups) + 1; i <= 1000000; i++ {
		newCup := cup{
			label: i,
			next:  nil,
		}
		output = append(output, &newCup)
	}
	// Link it up
	for idx, cup := range output {
		next := idx + 1
		if next >= len(output) {
			next = 0
		}
		cup.next = output[next]
		outputMap[cup.label] = idx
	}
	return output, outputMap
}

// Part 1 functions

func pickUpCups(cups []int, target int) ([]int, []int) {
	targetIdx := 0
	for idx, cup := range cups {
		if cup == target {
			targetIdx = idx
			break
		}
	}
	pickedUp := []int{}
	for i := 0; i < 3; i++ {
		targetIdx++
		if targetIdx >= len(cups) {
			targetIdx = 0
		}
		pickedUp = append(pickedUp, cups[targetIdx])
		cups[targetIdx] = 0
	}
	leftOver := []int{}
	for _, cup := range cups {
		if cup != 0 {
			leftOver = append(leftOver, cup)
		}
	}
	return leftOver, pickedUp
}

func destinationCup(leftover []int, target int) int {
	destination := 0
	_, max := minmax(leftover)
	for destination == 0 {
		target--
		if target <= 0 {
			target = max // Maybe needs to change in the future
		}
		if isIn(target, leftover) {
			destination = target
		}
	}
	return destination
}

func isIn(target int, array []int) bool {
	for _, num := range array {
		if num == target {
			return true
		}
	}
	return false
}

func placeCupsAfterDestination(cups []int, cupsToPlace []int, destination int) []int {
	output := []int{}
	for _, cup := range cups {
		output = append(output, cup)
		if cup == destination {
			output = append(output, cupsToPlace...)
		}
	}
	return output
}

func newCurrentCup(cups []int, currentCup int) int {
	for idx, cup := range cups {
		if cup == currentCup {
			if idx >= len(cups)-1 {
				idx = -1
			}
			return cups[idx+1]
		}
	}
	return 0
}

func deepCopyInt(input []int) []int {
	output := []int{}
	for _, num := range input {
		output = append(output, num)
	}
	return output
}

func returnMillionLongCups(cups []int) []int {
	output := []int{}
	for i := 1; i <= 1000000; i++ {
		if (i - 1) < len(cups) {
			output = append(output, cups[i-1])
		} else {
			output = append(output, i)
		}
	}
	return output
}

func minmax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func findOne(array []int) int {
	for idx, num := range array {
		if num == 1 {
			return idx
		}
	}
	return 0
}
