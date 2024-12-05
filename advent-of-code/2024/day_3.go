package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	url := "https://adventofcode.com/2024/day/3/input"
	file_name := "day_three.txt"
	input := string(GetDataForDay(url, file_name))
	//r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	r, _ := regexp.Compile("do\\(\\)|don't\\(\\)|(mul\\((\\d{1,3}),(\\d{1,3})\\))")
	matches := r.FindAllStringSubmatch(input, -1)
	out := 0
	stop := false
	for _, match := range matches{
		if match[0] == "don't()" || match[0] == "do()" {
			stop = (match[0] == "don't()")
			continue
		}

		if !stop {
			m1, err := strconv.Atoi(match[2])
			if err != nil {
				panic(fmt.Sprintf("failed to convert %s", match[2]))
			}
			m2, err := strconv.Atoi(match[3])
			if err != nil {
				panic(fmt.Sprintf("failed to convert %s", match[3]))
			}
			out += m1*m2
		}
	}
	fmt.Println(out)
}
