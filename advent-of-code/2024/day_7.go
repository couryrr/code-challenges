package main

import (
	"bytes"
	"fmt"
	"strconv"
)

var values []int
func main() {
	url := "https://adventofcode.com/2024/day/7/input"
	file_name := "day_seven.txt"

	input := GetDataForDay(url, file_name)
	
	lines := bytes.Split(input, []byte("\n"))
	total := 0	
	for i := range len(lines)-2 {
		parts := bytes.Split(lines[i], []byte(": "))
		goal, _ := strconv.Atoi(string(parts[0]))
		data := bytes.Split(parts[1], []byte(" "))
		for ii := range len(data) {
 			v, _ := strconv.Atoi(string(data[ii]))
			values = append(values, v)
		}

		if eval(values[0], goal, 1){
			total += goal
		}
		values = nil
	}
	fmt.Println(total)
}


func eval(total, goal, pos int) bool {
	if total == goal && pos == len(values) {
		return true
	}

	if total > goal {
		return false
	}

	if pos + 1 > len(values)  {
		return false
	}

	if eval(total + values[pos], goal, pos+1) {
		return true
	}

	if eval(total * values[pos], goal, pos+1) {
		return true
	}
	concat, _ := strconv.Atoi(strconv.Itoa(total)+strconv.Itoa(values[pos]))
	if eval(concat, goal, pos + 1) {
		return true
	}

	return false

}
