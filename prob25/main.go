package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const divider int = 20201227

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Data processing
	processedData := strings.Split(string(data), "\n")
	key1, err := strconv.Atoi(processedData[0])
	if err != nil {
		panic(err)
	}
	key2, err := strconv.Atoi(processedData[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(key1, key2)

	loopCount1 := 0
	value1 := 1
	for {
		loopCount1++
		value1 = loop(value1, 7)
		if value1 == key1 {
			break
		}
	}
	fmt.Println(loopCount1)
	loopCount2 := 0
	value2 := 1
	for {
		loopCount2++
		value2 = loop(value2, 7)
		if value2 == key2 {
			break
		}
	}
	fmt.Println(loopCount2)

	eValue := 1
	for i := 0; i < loopCount1; i++ {
		eValue = loop(eValue, key2)
	}
	fmt.Println(eValue)
}

func loop(value, subjectNumber int) int {
	value = value * subjectNumber
	value = value % divider
	return value
}
