package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("OwO")
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	processedData := strings.Split(string(data), "\n")

	valid := 0
	for _, password := range processedData {
		split := strings.Split(password, " ")
		nums := strings.Split(split[0], "-")
		num1, num2 := nums[0], nums[1]
		num1int, _ := strconv.Atoi(num1)
		num2int, _ := strconv.Atoi(num2)
		specialLetter := split[1][0]
		pass := split[2]

		letters := map[string]int{}
		for _, letter := range pass {
			letters[string(letter)]++
		}

		if letters[string(specialLetter)] >= num1int && letters[string(specialLetter)] <= num2int {
			valid++
		}
	}
	fmt.Println(valid)
	valid2 := 0
	for _, password := range processedData {
		split := strings.Split(password, " ")
		nums := strings.Split(split[0], "-")
		num1, num2 := nums[0], nums[1]
		num1int, _ := strconv.Atoi(num1)
		num2int, _ := strconv.Atoi(num2)
		specialLetter := split[1][0]
		pass := split[2]

		if string(pass[num1int-1]) == string(specialLetter) && string(pass[num2int-1]) == string(specialLetter) {
			continue
		}
		if string(pass[num1int-1]) == string(specialLetter) || string(pass[num2int-1]) == string(specialLetter) {
			valid2++
		}

	}
	fmt.Println(valid2)
}
