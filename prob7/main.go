package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Bag struct {
	number int
	color  string
}

func main() {
	fmt.Println("~(OwO)~")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	processedData := strings.Split(string(data), "\n")

	bagsMap := map[Bag][]Bag{}

	for _, line := range processedData {
		// I mean maybe I should learn regex...
		bags := strings.Split(line, " bags contain")
		color := bags[0]
		baseBag := Bag{1, color}
		bagsMap[baseBag] = []Bag{}
		contains := bags[1]
		containedBags := strings.Split(contains, ",")
		if strings.Contains(containedBags[0], "no other") {
			continue
		}
		for _, bag := range containedBags {
			if strings.Contains(bag, ".") {
				bag = bag[:len(bag)-1]
			}
			num, err := strconv.Atoi(bag[1:2])
			if err != nil {
				fmt.Println(bag, err)
			}
			subColor := strings.TrimSpace(bag[3 : len(bag)-4])
			bagsMap[baseBag] = append(bagsMap[baseBag], Bag{
				number: num,
				color:  subColor,
			})
		}
	}

	valid := 0
	for bag, _ := range bagsMap {
		lookInside(&valid, bagsMap, bag.color)
	}

	fmt.Println(valid)

	// part 2
	bags := 0
	lookAndCount(&bags, bagsMap, "shiny gold")
	fmt.Println(bags)

}

func lookInside(valid *int, bagsMap map[Bag][]Bag, color string) bool {
	for _, eachBag := range bagsMap[Bag{1, color}] {
		if eachBag.color == "shiny gold" {
			*valid++
			return true
		}
		found := lookInside(valid, bagsMap, eachBag.color)
		if found {
			return true
		}
	}
	return false
}

func lookAndCount(bags *int, bagsMap map[Bag][]Bag, color string) {
	for _, eachBag := range bagsMap[Bag{1, color}] {
		for i := 0; i < eachBag.number; i++ {
			*bags++
			lookAndCount(bags, bagsMap, eachBag.color)
		}

	}
}
