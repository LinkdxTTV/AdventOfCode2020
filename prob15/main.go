package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Data processing
	processedData := strings.Split(string(data), ",")
	numbers := []int{}
	for _, num := range processedData {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, numInt)
	}
	fmt.Println("input", numbers)
	turn := 0

	// timer
	start := time.Now()

	spokenMap := map[int][]int{}
	var spoken int
	end := 30000000

	for i := 0; i < end; i++ {
		turn++
		if turn-1 < len(numbers) {
			spoken = numbers[turn-1]
			spokenMap[spoken] = append(spokenMap[spoken], turn)
			continue
		}
		// spoken should still be assigned from last loop
		lastTurn, _ := spokenMap[spoken]
		length := len(lastTurn)
		if length == 1 { // Means it was only spoken last turn, utter a 0
			spoken = 0
		} else {
			spoken = lastTurn[length-1] - lastTurn[length-2]
		}

		spokenMap[spoken] = append(spokenMap[spoken], turn)
	}

	fmt.Println(spoken)
	// End timer
	fmt.Println(time.Now().Sub(start))
}

// Example answer
// 0 3 6 0 3 3 1 0 4 0
