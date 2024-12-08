package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	url := "https://adventofcode.com/2024/day/5/input"
	file_name := "day_five.txt"

	input := bytes.Split(GetDataForDay(url, file_name), []byte("\n\n"))

	rules := bytes.Split(input[0], []byte("\n"))
	ruleMap :=  make(map[int][]int)
	for i := range rules {
		tmp := bytes.Split(rules[i], []byte("|"))
		key, _ := strconv.Atoi(string(tmp[0]))
		value, _ := strconv.Atoi(string(tmp[1]))
		ruleMap[key] = append(ruleMap[key], value) 
	}
	
	tmp := bytes.Split(input[1], []byte("\n"))
	var list [][]int 
	for i := range len(tmp) - 1 { // -1 for the white space. I should handle this better...
		var pages []int
		page_list := bytes.Split(tmp[i], []byte(","))
		for ii := range len(page_list) {
			page, _ := strconv.Atoi(string(page_list[ii]))
			pages = append(pages, page)  
		}	
		list = append(list, pages)
	}

	input = nil
	rules = nil
	tmp = nil

	total := 0
	for i := range list {
		found := false
		for ii := range list[i] {
			page_rule := ruleMap[list[i][ii]]
			prev_pages := list[i][:ii]
			for _, prev_page := range prev_pages {
				if slices.Contains(page_rule, prev_page) {
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			total += list[i][len(list[i])/2]
		}
	}

	fmt.Println(total)
}
