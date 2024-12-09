package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	url := "https://adventofcode.com/2024/day/5/input"
	file_name := "day_five.txt"

	input := bytes.Split(GetDataForDay(url, file_name), []byte("\n\n"))
	solution_1(input)
}

func solution_1(input [][]byte){
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
	for i := range len(tmp) {
		var pages []int
		page_list := bytes.Split(tmp[i], []byte(","))
		for ii := range len(page_list) {
			page, _ := strconv.Atoi(string(page_list[ii]))
			pages = append(pages, page)  
		}	
		list = append(list, pages)
	}

	rules = nil
	tmp = nil

	s1 := 0
	s2 := 0
	for i := range list {
		found := false
		for ii := range list[i] {
			page_rule := ruleMap[list[i][ii]]
			prev_pages := list[i][:ii]
			for _, prev_page := range prev_pages {
				if slices.Contains(page_rule, prev_page) {
					found = true
					s2 += solution_2(input, list[i])
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			s1 += list[i][len(list[i])/2]
		}
		fmt.Println(list[i])
	}

	fmt.Println(s1)
	fmt.Println(s2)

}

func solution_2(input [][]byte, pages []int) int {
	// so the map was not needed and probably made this harder
	rules := string(input[0])
	slices.SortStableFunc(pages, func(x, y int) int {
		if strings.Contains(rules, strconv.Itoa(x)+"|"+strconv.Itoa(y)) {
			return -1
		}
		if strings.Contains(rules, strconv.Itoa(x)+"|"+strconv.Itoa(y)) {
			return 1
		}
		return 0
	})
	return pages[len(pages)/2]

}
