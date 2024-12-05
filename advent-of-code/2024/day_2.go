package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	url := "https://adventofcode.com/2024/day/2/input"
	file_name := "day_two.txt"

	input := GetDataForDay(url, file_name)
	
	lines := strings.Split(string(input), "\n")
	fmt.Println(FindSafeLines(lines))
}

func IsSafe(prev, curr, next int) bool {
	if IntAbs(curr, prev) > 3 {
		return false
	}

	if next > -1 {
		if !(prev < curr && curr < next) && !(prev > curr && curr > next) {
			return false
		}
	}

	return true
}

func FindSafeLines(lines []string) int {
	count := 0
	for i := range len(lines)-1 {
		values := strings.Split(lines[i], " ") 
		for i := 0; i <= len(values)-1; i++ {
			// Append can modify the underlying slice...
			tmp := make([]string, len(values))
			copy(tmp, values)
			if CheckLine(append(tmp[:i], tmp[i+1:]...)) {
				count++
				break
			}
		}
	}
	return count
} 

func CheckLine(values []string) bool {
	for i := 1; i < len(values); i++ {
		prev, err := strconv.Atoi(values[i-1])
		if err != nil {
			panic(fmt.Sprintf("bad string to int conversion for: %s",  values[i-1]))
		}
		
		curr, err := strconv.Atoi(values[i])
		if err != nil {
			panic(fmt.Sprintf("bad string to int conversion for: %s",  values[i]))
		}

		next := -1
		if i+1 <= len(values)-1{
			next, err = strconv.Atoi(values[i+1])
			if err != nil {
				panic(fmt.Sprintf("bad string to int conversion for: %s",  values[i+1]))
			}
		}

		if !IsSafe(prev, curr, next){
			return false
		}
	}
	return true
}
func IntAbs(x, y int) int{ 
	val := 0
	if x < y {
		val = y - x
	} else {
		val = x - y
	}
	return val

}
