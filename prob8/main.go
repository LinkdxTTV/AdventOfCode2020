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
	accumulator := 0
	index := 0
	lenArray := len(processedData)
	visited := []bool{}
	nopAndjmp := []int{}
	for i := 0; i < lenArray; i++ {
		// Preprocess the the NOP and JMP commands
		visited = append(visited, false)
		line := processedData[i]
		split := strings.Split(line, " ")
		instruction := split[0]
		if instruction == "nop" || instruction == "jmp" {
			nopAndjmp = append(nopAndjmp, i)
		}
	}

	for _, nopOrjmpIndex := range nopAndjmp {
		// Reset everything
		index = 0
		accumulator = 0
		for i := 0; i < len(visited); i++ {
			visited[i] = false
		}
		for {
			if index >= len(processedData) {
				fmt.Println("Terminated Succesfully, Accumulator value is: ", accumulator)
				break
			}
			if visited[index] {
				break
			}
			visited[index] = true
			line := processedData[index]
			split := strings.Split(line, " ")
			instruction, num := split[0], split[1]
			numSign := string(num[0])
			numInt, _ := strconv.Atoi(num[1:])

			if index == nopOrjmpIndex {
				if instruction == "jmp" {
					instruction = "nop"
				} else if instruction == "nop" {
					instruction = "jmp"
				}
			}

			switch instruction {
			case "nop":
				index++

			case "acc":
				switch numSign {
				case "+":
					accumulator += numInt
				case "-":
					accumulator -= numInt
				}
				index++

			case "jmp":
				switch numSign {
				case "+":
					index += numInt
				case "-":
					index -= numInt
				}
			default:
				fmt.Println("unknown instruction", instruction)
			}

		}
	}

}
