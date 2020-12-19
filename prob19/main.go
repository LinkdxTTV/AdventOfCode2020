package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ruleNode struct {
	subRules   [][]int
	substrings []string
}

var (
	a string = "a"
	b string = "b"
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

	rulesMap := map[int]ruleNode{}

	ruleEndIndex := 0

	// Rule Processing
	for index, line := range processedData {
		if line == "" {
			ruleEndIndex = index + 1
			break
		}
		splitLine := strings.Split(line, ":")
		// splitLine[0] is rule number, splitLine[1] is the subRules

		ruleNumber, err := strconv.Atoi(splitLine[0])
		if err != nil {
			panic(err)
		}

		if splitLine[1] == " \"a\"" {
			rulesMap[ruleNumber] = ruleNode{nil, []string{a}}
			continue
		}
		if splitLine[1] == " \"b\"" {
			rulesMap[ruleNumber] = ruleNode{nil, []string{b}}
			continue
		}

		node := ruleNode{subRules: [][]int{}}
		rules := strings.Split(splitLine[1], "|")
		for _, ruleSet := range rules {
			rulesArray := []int{}
			individualRules := strings.Split(ruleSet, " ")
			for _, indieRule := range individualRules {
				if indieRule == "" {
					continue
				}
				ruleInt, err := strconv.Atoi(indieRule)
				if err != nil {
					fmt.Println(indieRule)
					panic(err)
				}
				rulesArray = append(rulesArray, ruleInt)
			}
			node.subRules = append(node.subRules, rulesArray)
		}
		rulesMap[ruleNumber] = node
	}

	rules := generateSubstrings(rulesMap[0], 0, rulesMap)
	fmt.Println("31", rulesMap[31], "42", rulesMap[42])

	// Data check
	valid := 0

	map31 := map[string]bool{}
	map42 := map[string]bool{}
	//Process rulesMap31 and rulesMap42
	for _, string1 := range rulesMap[31].substrings {
		map31[string1] = true
	}
	for _, string1 := range rulesMap[42].substrings {
		map42[string1] = true
	}

	for _, line := range processedData[ruleEndIndex:] {

		for _, subString := range rules {
			if line == subString {
				valid++
				break
			}

			// Special Part 2 processing
			checking31 := true
			a, b := 0, 0
			for len(line) > 0 {
				if len(line)%8 == 0 {
					is31, _ := map31[line[len(line)-8:]]
					is42, _ := map42[line[0:8]]
					if is31 && is42 && checking31 {
						a++
						b++
						line = line[8 : len(line)-8]
					}
					if is31 && !is42 && checking31 {
						break
					}
					if !is31 && is42 {
						a++
						line = line[8:]
						checking31 = false
					}
					if !is31 && !is42 {
						break
					}
				} else {
					break
				}
			}
			if len(line) == 0 && b >= 1 && a > b {
				valid++
			}
		}

	}

	fmt.Println(valid)
}

func generateSubstrings(node ruleNode, nodeNum int, rulesMap map[int]ruleNode) []string {
	if len(node.substrings) != 0 {
		return node.substrings
	}
	node.substrings = []string{}
	for _, subRule := range node.subRules { // 2 3 | 3 2

		possibles1 := generateSubstrings(rulesMap[subRule[0]], subRule[0], rulesMap)
		if len(subRule) == 1 {
			for _, string1 := range possibles1 {
				node.substrings = append(node.substrings, string1)
			}
			continue
		}
		possibles2 := generateSubstrings(rulesMap[subRule[1]], subRule[1], rulesMap)
		if len(subRule) == 2 {

			for _, string1 := range possibles1 {
				for _, string2 := range possibles2 {
					node.substrings = append(node.substrings, string1+string2)
				}
			}
		}
		if len(subRule) == 3 {
			possibles3 := generateSubstrings(rulesMap[subRule[2]], subRule[2], rulesMap)
			for _, string1 := range possibles1 {
				for _, string2 := range possibles2 {
					for _, string3 := range possibles3 {
						node.substrings = append(node.substrings, string1+string2+string3)
					}

				}
			}
		}

	}
	rulesMap[nodeNum] = node
	return node.substrings
}
