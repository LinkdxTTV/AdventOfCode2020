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
	bitMaskObj := bitMask("")
	memMap := map[int]int{}

	// Part 1

	for _, instruction := range processedData {
		// Ill never learn regex
		instructionSplit := strings.Split(instruction, " ")
		if instructionSplit[0] == "mask" {
			bitMaskObj = bitMask(instructionSplit[2])
		} else {
			memoryValue, err := strconv.Atoi(instructionSplit[2])
			if err != nil {
				panic(err)
			}
			memoryAddress := strings.Split(instructionSplit[0], "[")
			memoryAddress = strings.Split(memoryAddress[1], "]")
			memoryAddressInt, err := strconv.Atoi(memoryAddress[0])
			if err != nil {
				panic(err)
			}

			memMap[memoryAddressInt] = bitMaskObj.alterIntWithBitmask(memoryValue)

		}
	}

	sum := 0
	for _, value := range memMap {
		sum += value
	}
	fmt.Println(sum)

	// Part 2

	memMap2 := map[int]int{}
	for _, instruction := range processedData {
		// Ill never learn regex
		instructionSplit := strings.Split(instruction, " ")
		if instructionSplit[0] == "mask" {
			bitMaskObj = bitMask(instructionSplit[2])
		} else {
			memoryValue, err := strconv.Atoi(instructionSplit[2])
			if err != nil {
				panic(err)
			}
			memoryAddress := strings.Split(instructionSplit[0], "[")
			memoryAddress = strings.Split(memoryAddress[1], "]")
			memoryAddressInt, err := strconv.Atoi(memoryAddress[0])
			if err != nil {
				panic(err)
			}

			addressWithX := bitMaskObj.fixBitmask(memoryAddressInt)
			addressesToEdit := []int{}
			recursivelyCreateOutputs(&addressesToEdit, addressWithX, 0)
			for _, address := range addressesToEdit {
				memMap2[address] = memoryValue
			}

		}
	}
	sum = 0
	for _, value := range memMap2 {
		sum += value
	}
	fmt.Println(sum)

}

func intToBinaryString(input int) string {
	return strconv.FormatInt(int64(input), 2)
}

type bitMask string

func (b *bitMask) alterIntWithBitmask(input int) int {

	// Get input and pad
	binaryInput := intToBinaryString(input)

	// Pad
	for len(binaryInput) < len(*b) {
		binaryInput = "0" + binaryInput
	}

	outputString := []string{}
	bitMaskLiteral := *b
	for i := 0; i < len(bitMaskLiteral); i++ {
		if string(bitMaskLiteral[i]) != "X" {
			outputString = append(outputString, string(bitMaskLiteral[i]))
		} else {
			outputString = append(outputString, string(binaryInput[i]))
		}
	}

	outputBinary := strings.Join(outputString, "")
	outputInt, err := strconv.ParseInt(outputBinary, 2, 64)
	if err != nil {
		fmt.Println("Alter int error:", err)
	}
	return int(outputInt)
}

func (b *bitMask) fixBitmask(input int) string {
	// Get input and pad
	binaryInput := intToBinaryString(input)

	// Pad
	for len(binaryInput) < len(*b) {
		binaryInput = "0" + binaryInput
	}

	outputString := []string{}
	bitMaskLiteral := *b
	for i := 0; i < len(bitMaskLiteral); i++ {
		switch string(bitMaskLiteral[i]) {
		case "0":
			outputString = append(outputString, string(binaryInput[i]))
		case "1":
			outputString = append(outputString, string(bitMaskLiteral[i]))
		case "X":
			outputString = append(outputString, string(bitMaskLiteral[i]))
		}

	}
	return strings.Join(outputString, "")

}

func recursivelyCreateOutputs(store *[]int, mask string, place int) {
	if place == len(mask) {
		maskInt, err := strconv.ParseInt(mask, 2, 64)
		if err != nil {
			fmt.Println("Alter int error:", err)
		}
		*store = append(*store, int(maskInt))
		return
	}
	if string(mask[place]) == "X" {
		mask0 := mask[:place] + "0" + mask[place+1:]
		recursivelyCreateOutputs(store, mask0, place+1)
		mask1 := mask[:place] + "1" + mask[place+1:]
		recursivelyCreateOutputs(store, mask1, place+1)
	} else {
		recursivelyCreateOutputs(store, mask, place+1)
	}
}
