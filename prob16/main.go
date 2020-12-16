package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type rule struct {
	min int
	max int
}

func (r rule) isIn(input int) bool {
	if input >= r.min && input <= r.max {
		return true
	}
	return false
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

	rulesMap := map[string][]rule{}
	// CHANGE THIS LATER
	linesOfRules := 20
	for _, rulesLine := range processedData[0:linesOfRules] {
		lineSplit := strings.Split(rulesLine, ":")
		ruleName := lineSplit[0]
		numRanges := strings.Split(lineSplit[1], " or ")
		rules := []rule{}
		for _, numRange := range numRanges {
			trimmed := strings.TrimSpace(numRange)
			minMax := strings.Split(trimmed, "-")
			minInt, err := strconv.Atoi(minMax[0])
			if err != nil {
				panic(err)
			}
			maxInt, err := strconv.Atoi(minMax[1])
			if err != nil {
				panic(err)
			}
			rules = append(rules, rule{minInt, maxInt})
		}
		rulesMap[ruleName] = rules
	}

	// fmt.Println(rulesMap)

	tickets := [][]int{}
	for _, inputLines := range processedData[linesOfRules+1:] {
		if inputLines == "your ticket:" || inputLines == "nearby tickets:" || inputLines == "" {
			continue
		}
		ticketNums := strings.Split(inputLines, ",")
		numList := []int{}
		for _, ticketNum := range ticketNums {
			ticketNumInt, err := strconv.Atoi(ticketNum)
			if err != nil {
				panic(err)
			}
			numList = append(numList, ticketNumInt)
		}
		tickets = append(tickets, numList)
	}

	invalidNumbers := []int{}
	ticketsToBeDiscared := map[int]bool{}
	for ticketIndex, ticket := range tickets {
		for _, number := range ticket {
			numberSatisfied := false
			for _, rules := range rulesMap {
				for _, rule := range rules {
					if rule.isIn(number) {
						numberSatisfied = true
					}
				}
			}
			if numberSatisfied == false {
				invalidNumbers = append(invalidNumbers, number)
				ticketsToBeDiscared[ticketIndex] = true
			}
		}
	}
	// fmt.Println(invalidNumbers)
	sum := 0
	for _, num := range invalidNumbers {
		sum += num
	}
	fmt.Println("Part 1:", sum)

	// Part 2

	// Initialize some garbage
	possibleMap := map[string]map[int]int{}
	for name := range rulesMap {
		nameMap := map[int]int{}
		for i := 0; i < len(tickets[0]); i++ {
			nameMap[i] = 0
		}
		possibleMap[name] = nameMap
	}

	for ticketIndex, ticket := range tickets {
		// Trash bad tickets
		if _, ok := ticketsToBeDiscared[ticketIndex]; ok {
			continue
		}
		for numberIndex, number := range ticket {
			for ruleName, rules := range rulesMap {
				fits := false
				for _, rule := range rules {
					if rule.isIn(number) {
						fits = true
						break
					}
				}
				if !fits {
					possibleMap[ruleName][numberIndex]++
				}
			}

		}
	}

	// OKAY..?
	ruleToColumn := map[string]int{}
	columnsTaken := map[int]bool{}
	columnsPossible := 1
	for len(possibleMap) > 0 {
		for rule, ruleMap := range possibleMap {
			possibleColumns := []int{}
			for column, badFits := range ruleMap {
				if badFits == 0 {
					possibleColumns = append(possibleColumns, column)
				}
			}
			if len(possibleColumns) == columnsPossible {
				columnHere := 0
				for _, columnNumber := range possibleColumns {
					_, found := columnsTaken[columnNumber]
					if !found {
						columnHere = columnNumber
						break
					}
				}
				ruleToColumn[rule] = columnHere
				columnsTaken[columnHere] = true
				delete(possibleMap, rule)
				columnsPossible++

				break
			}
		}
	}
	product := 1

	myTicket := tickets[0]
	for k, v := range ruleToColumn {
		if strings.Contains(k, "departure") {
			product = product * myTicket[v]
		}
	}
	fmt.Println("Part 2:", product)
}
