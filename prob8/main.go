package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type instruction struct {
	command string
	value   int
}

type instructionSet struct {
	set     []instruction
	visited map[int]int
	// Part 2, exchange indices

}

var (
	errorRevisit = errors.New("revisited an instruction")
	errorTerm    = errors.New("sequence terminated")
)

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	processedData := strings.Split(string(data), "\n")

	// Preprocessing
	instructionSet := instructionSet{[]instruction{}, make(map[int]int)}
	flipIndices := []int{}
	for idx, line := range processedData {
		split := strings.Split(line, " ")
		command, num := split[0], split[1]
		numInt, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("strconv err", err)
		}

		instruction := instruction{
			command: command,
			value:   numInt,
		}

		instructionSet.set = append(instructionSet.set, instruction)
		instructionSet.visited[idx] = 0

		// Part 2 preprocessing

		if command == "jmp" || command == "nop" {
			flipIndices = append(flipIndices, idx)
		}
	}
	// Part 1
	// Make a deep copy so we can reuse the original later
	instructionSetCopy := instructionSetDeepCopy(instructionSet)
	accumulator, err := instructionSetCopy.runSet(0)
	if err == errorRevisit {
		fmt.Println(err, "Accumulator value: ", accumulator)
	}

	// Part 2
	for _, idx := range flipIndices {
		// Make a deep copy each time
		instructionSetCopy := instructionSetDeepCopy(instructionSet)
		instructionSetCopy.set[idx].flip()

		accumulator, err := instructionSetCopy.runSet(0)
		if err == errorTerm {
			fmt.Println(err, "Accumulator value: ", accumulator)
		}
	}
}

func (s instructionSet) runSet(start int) (accumulator int, err error) {
	i := start
	// Flip these instructions for part 2
	for {
		if i >= len(s.set) {
			return accumulator, errorTerm
		}
		instruction := s.set[i]

		// Disallow Loops
		if s.visited[i] != 0 {
			return accumulator, errorRevisit
		}
		s.visited[i]++

		// Do
		instruction.do(&i, &accumulator)
	}
}

// Do the instruction
func (i instruction) do(index *int, accumulator *int) {
	switch i.command {
	case "nop":
		*index++

	case "acc":
		*accumulator += i.value
		*index++

	case "jmp":
		*index += i.value

	default:
		fmt.Println("unknown instruction", i.command)
	}
}

// Flip this instruction
func (i *instruction) flip() {
	if i.command == "jmp" {
		i.command = "nop"
	} else if i.command == "nop" {
		i.command = "jmp"
	}
}

// Deep Copying sucks
func instructionSetDeepCopy(instructionSet instructionSet) instructionSet {
	instructionSetCopy := instructionSet
	instructionSetCopy.set = make([]instruction, len(instructionSet.set))
	copy(instructionSetCopy.set, instructionSet.set)
	instructionSetCopy.visited = map[int]int{}
	for k, v := range instructionSet.visited {
		instructionSetCopy.visited[k] = v
	}
	return instructionSetCopy
}
