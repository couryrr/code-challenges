package main

import (
	"bytes"
	"fmt"
	"slices"
	"strconv"
)

func main() {
	url := ""
	file_name := "day_nine.txt"

	input := bytes.TrimSpace(GetDataForDay(url, file_name))
	expanded := make([]int, 0, 10)
	file_id := 0
	isFree := false
	for i, v := range input {
		vv, _ := strconv.Atoi(string(v))
		isFree = (i+1)%2 == 0
		for ii := 0; ii < vv; ii++ {
			if isFree {
				expanded = append(expanded, -1)
			} else {
				expanded = append(expanded, file_id)
			}
		}
		if !isFree {
			file_id++
		}
		isFree = false
	}
	
	free := slices.Index(expanded, -1)
	end := len(expanded) - 1

	for free < end {
		if expanded[end] != -1 {
			expanded[free] = expanded[end]
			expanded[end] = -1
			free = slices.Index(expanded, -1)
		} 
		end--
	}

	s1 := 0
	for i, v := range expanded {
		if v == -1 {
			break
		}
		s1 += i*v
	}
	fmt.Println(s1)
}
