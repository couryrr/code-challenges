package main

import (
	"bytes"
	"fmt"
)

var input [][]byte
var word []byte

func main() {
	url := "https://adventofcode.com/2024/day/4/input"
	file_name := "day_four.txt"
	input = bytes.Split(GetDataForDay(url, file_name), []byte("\n"))

	word = []byte("XMAS")
	fmt.Println(search_one())
	
	//not correct yet...
	word = []byte("MAS")
	fmt.Println(search_two())
}

func search_one() int{
	count := 0
	for r := range len(input) {
		for c := range len(input[r]) {
			if dfs(r, c, 0, 1, 0){
				count++
			}

			if dfs(r, c, 0, -1, 0){
				count++
			}

			if dfs(r, c, -1, 0, 0){
				count++
			}
			
			if dfs(r, c, 1, 0, 0){
				count++
			}

			if dfs(r, c, 1, -1, 0) {
				count++
			}

			if dfs(r, c, 1, 1, 0) {
				count++
			}
			
			if dfs(r, c, -1, -1, 0) {
				count++
			}

			if dfs(r, c, -1, 1, 0) {
				count++
			}
		}
	}
	return count
}

func search_two() int {
	count := 0
	for r := range len(input) {
		for c := range len(input[r]) {
			if input[r][c] == 'A' {
				if r - 1 < 0 || r + 1 >= len(input) { 
					break
				}

				if c - 1 < 0 || c + 1 >= len(input[r+1]) {
					break	
				}
				
				c1 := input[r-1][c-1]
				c2 := input[r+1][c+1]
				idk1 := (c1 == 'M' || c1 == 'S') && (c1 != c2)

				c3 := input[r-1][c+1]
				c4 := input[r+1][c-1]
				idk2 := (c3 == 'M' || c3 == 'S') && (c3 != c4)

				if idk1 && idk2 {
					count++
				}
			}
		}
	}
	return count
}

func dfs(row, col, dr, dc, word_idx int) bool {
	if word_idx == len(word) {
		return true
	}
	
	if row < 0 || row >= len(input) { 
		return false 
	}

	if col < 0 || col >= len(input[row]) {
		return false 
	}

	if input[row][col]!= word[word_idx] {
		return false	
	}

	if dfs(row + dr, col + dc, dr, dc, word_idx + 1) { 
		return true
	}

	return false
}

