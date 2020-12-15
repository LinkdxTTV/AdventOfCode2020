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
		lastTurn, ok := spokenMap[spoken]
		if ok {
			if len(lastTurn) == 0 || len(lastTurn) == 1 {
				spoken = 0
			} else {
				spoken = lastTurn[len(lastTurn)-1] - lastTurn[len(lastTurn)-2]
			}
		} else {
			spoken = 0
		}
		spokenMap[spoken] = append(spokenMap[spoken], turn)
	}
	fmt.Println(spoken)
}

// 0 3 6 0 3 3 1 0 4 0
