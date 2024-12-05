package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "https://adventofcode.com/2024/day/4/input"
	file_name := "day_four.txt"
	input := strings.Split(string(GetDataForDay(url, file_name)), "\n")

	word := []byte("XMAS")
	count := 0	

	for r := range input {
		for c := range input[r] {
			if dfs(input, word, r, c, 0){
				count++
			}
		}
	}
	fmt.Println(count)
}

func dfs(input []string, word []byte, row, col, word_idx int) bool {
	if word_idx == len(word) {
		return true
	}
	
	row_overflow := row < 0 || row >= len(input)
	col_overflow := col < 0 || col >= len(input[row])

	if row_overflow || col_overflow {
		return false
	}

	letters := []byte(input[row])
	letter := letters[col]
	if letter != word[word_idx] {
		return false	
	}
	
	letters[col] = byte('%')
	if dfs(input, word, row, col - 1, word_idx + 1) || // why can't this find overlaps?
		dfs(input, word, row, col + 1, word_idx + 1) {
		return true
	}
	letters[col] = letter

	return false
}
