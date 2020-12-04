package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	// Pass processing
	passports := []map[string]string{}
	mapPass := map[string]string{}
	for _, line := range processedData {
		entries := strings.Split(string(line), " ")
		for _, entry := range entries {
			kv := strings.Split(string(entry), ":")
			if len(kv) > 1 {
				mapPass[kv[0]] = kv[1]
			}
		}
		if line == "" {
			passports = append(passports, mapPass)
			mapPass = map[string]string{}
		}
	}
	valid := 0
	fmt.Println(len(passports))
	eyecolors := map[string]string{
		"amb": "",
		"blu": "",
		"brn": "",
		"hzl": "",
		"grn": "",
		"gry": "",
		"oth": "",
	}
	for _, passport := range passports {
		if byr, ok := passport["byr"]; !ok {
			continue
		} else {
			byrInt, _ := strconv.Atoi(byr)
			if byrInt < 1920 || byrInt > 2002 {
				continue
			}
		}
		if iyr, ok := passport["iyr"]; !ok {
			continue
		} else {
			iyrInt, err := strconv.Atoi(iyr)
			if iyrInt < 2010 || iyrInt > 2020 {
				continue
			}
			if err != nil {
				continue
			}
		}
		if eyr, ok := passport["eyr"]; !ok {
			continue
		} else {
			eyrInt, err := strconv.Atoi(eyr)
			if eyrInt < 2020 || eyrInt > 2030 {
				continue
			}
			if err != nil {
				continue
			}
		}
		if hgt, ok := passport["hgt"]; !ok {
			continue
		} else {
			height := false
			if strings.Contains(hgt, "cm") {
				height = true
				num, err := strconv.Atoi(hgt[0:(len(hgt) - 2)])
				if num < 150 || num > 193 {
					continue
				}
				if err != nil {
					continue
				}
			}
			if strings.Contains(hgt, "in") {
				height = true
				num, _ := strconv.Atoi(hgt[0:(len(hgt) - 2)])
				if num < 59 || num > 76 {
					continue
				}
			}
			if !height {
				continue
			}
		}
		if hcl, ok := passport["hcl"]; !ok {
			continue
		} else {
			if string(hcl[0]) != "#" {
				continue
			}
			if len(hcl) != 7 {
				continue
			}
			regex, err := regexp.Match("^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$", []byte(hcl))
			if err != nil {
				fmt.Println(err)
			}
			if !regex {
				continue
			}
		}
		if ecl, ok := passport["ecl"]; !ok {
			continue
		} else {
			if _, ok := eyecolors[ecl]; !ok {
				continue
			}
		}
		if pid, ok := passport["pid"]; !ok {
			continue
		} else {
			if len(pid) != 9 {
				continue
			}
			_, err := strconv.Atoi(pid)
			if err != nil {
				continue
			}
		}
		valid = valid + 1
	}
	fmt.Println(valid)
}
