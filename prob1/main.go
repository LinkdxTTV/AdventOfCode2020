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
	fmt.Println(processedData)
	for _, num1 := range processedData {
		for _, num2 := range processedData {
			for _, num3 := range processedData {

				if num1 == num2 || num1 == num3 || num2 == num3 {
					continue
				}
				num1Int, _ := strconv.Atoi(num1)
				num2Int, _ := strconv.Atoi(num2)
				num3Int, _ := strconv.Atoi(num3)
				if num1Int+num2Int+num3Int == 2020 {
					fmt.Println(num1Int, num2Int, num3Int)
					fmt.Println(num1Int * num2Int * num3Int)
				}
			}
		}
	}
}
