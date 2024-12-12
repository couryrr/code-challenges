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
	
	expanded, freeSapces := expand(input)
//	s1 := make([]int, len(expanded))   
//	copy(s1, expanded)
//	solution_d9_1(s1)

	end := len(expanded) - 1
	fmt.Printf("%v\n", freeSapces)
	for k, v := range freeSapces {
		fmt.Println(k, v)
	}

	for end >= 0 {
		if expanded[end] != -1 {
			curr_value := expanded[end]
			curr_end := end
			curr_len := 0

			for end >= 0 && expanded[end] == curr_value {
				curr_len++
				end--
			}
			fmt.Printf("for %d the len is %d and end is %d\n",  curr_value, curr_len, curr_end)

				} else {
			end--
		}
	}

	fmt.Printf("%v\n", expanded)
}

func solution_d9_1(expanded []int){
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

	s1 := checksum(expanded)

	fmt.Println(s1)
}

func expand(input []byte) ([]int, map[int]int){
	freeSapces := make(map[int]int, 1024)
	expanded := make([]int, 0, 1024)
	file_id := 0
	isFree := false
	for i, v := range input {
		vv, _ := strconv.Atoi(string(v))
		isFree = (i+1)%2 == 0
		if isFree {
			freeSapces[len(expanded)] = vv
		}
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
	return expanded, freeSapces
}

func checksum(expanded []int) int {
	checksum := 0
	for i, v := range expanded {
		if v == -1 {
			break
		}
		checksum += i*v
	}
	return checksum
}
