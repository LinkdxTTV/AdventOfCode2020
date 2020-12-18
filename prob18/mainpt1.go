package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const add = "+"
const multiply = "*"
const openParentheses = "("
const closeParentheses = ")"

type operation struct {
	value       int
	operation   string
	parentheses string
}

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Data processing
	processedData := strings.Split(string(data), "\n")

	operationsList := [][]operation{}
	for _, line := range processedData {
		operations := []operation{}
		dataWithoutWhitespace := ""
		for _, char := range line {
			if string(char) != " " {
				dataWithoutWhitespace += string(char)
			}
		}
		for index, char := range dataWithoutWhitespace {
			operation := operation{}
			valueInt, err := strconv.Atoi(string(char))
			if err != nil {
				if string(char) == closeParentheses {
					operation.operation = ""
					operation.value = 0
					operation.parentheses = string(char)
				} else if string(char) == openParentheses {
					operation.operation = ""
					operation.value = 0
					operation.parentheses = string(char)
					for operation.operation == "" {
						index--
						if index < 0 {
							operation.operation = add
						} else if string(dataWithoutWhitespace[index]) == add || string(dataWithoutWhitespace[index]) == multiply {
							operation.operation = string(dataWithoutWhitespace[index])
						}
					}
				} else {
					continue

				}
			} else {
				operation.value = valueInt
				for operation.operation == "" {
					index--
					if index < 0 {
						operation.operation = add
					} else if string(dataWithoutWhitespace[index]) == add || string(dataWithoutWhitespace[index]) == multiply {
						operation.operation = string(dataWithoutWhitespace[index])
					} else if string(dataWithoutWhitespace[index]) == openParentheses || string(dataWithoutWhitespace[index]) == closeParentheses {
						operation.operation = add
					}
				}
			}

			operations = append(operations, operation)
		}

		operationsList = append(operationsList, operations)
	}
	fmt.Println(operationsList)
	endOps := []operation{}
	for _, op := range operationsList {
		endOps = append(endOps, reduceOpsToValue(op, add))
		fmt.Println(endOps[len(endOps)-1])
	}

	fmt.Println(reduceOpsToValue(endOps, add))
}

func reduceOpsToValue(ops []operation, operator string) operation {
	simplified := []operation{}
	parenthesesLayers := 0
	// fmt.Println("ops", ops)
	for index, operation := range ops {

		// fmt.Println(operation, "layer", parenthesesLayers)
		if operation.parentheses == openParentheses {

			if parenthesesLayers == 0 {
				// Find the next end Parentheses
				var endIndex int
				innerPar := 1
				for opsIndex, op := range ops {
					if opsIndex <= index {
						continue
					}
					if op.parentheses == openParentheses {
						innerPar++
					}
					if op.parentheses == closeParentheses {
						innerPar--
					}
					if innerPar == 0 {
						endIndex = opsIndex
						break
					}
				}
				currentOp := reduceOpsToValue(ops[index+1:endIndex+1], operation.operation)
				simplified = append(simplified, currentOp)
				// fmt.Println("I added a simplification", simplified)
			}
			parenthesesLayers++
		}
		if operation.parentheses == closeParentheses {
			parenthesesLayers--
		}
		if parenthesesLayers != 0 {
			continue
		}
		if operation.parentheses == "" {
			simplified = append(simplified, operation)
		}
	}
	// fmt.Println("simplified", simplified)
	answer := 0
	if simplified[0].operation == multiply {
		answer = 1
	}
	for _, op := range simplified {
		switch op.operation {
		case add:
			answer += op.value

		case multiply:
			answer = answer * op.value
		default:
			fmt.Println("You fucked up", op, simplified)
		}
	}

	// fmt.Println("Im returning this: ", answer, operator)
	return operation{
		value:     answer,
		operation: operator,
	}
}
