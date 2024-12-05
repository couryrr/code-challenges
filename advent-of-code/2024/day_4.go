package main

import "fmt"

func main() {
	url := "https://adventofcode.com/2024/day/4/input"
	file_name := "day_four.txt"
	input := GetDataForDay(url, file_name)
	
	fmt.Println(string(input))

}
