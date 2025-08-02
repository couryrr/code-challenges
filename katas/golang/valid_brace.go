package main

import "fmt"

// https://www.codewars.com/kata/5277c8a221e209d3f6000b56

func main() {
	fmt.Println(ValidBraces("[({})](]"))
}

var pairs = map[rune]rune{
	'{': '}',
	'(': ')',
	'[': ']',
}

func ValidBraces(str string) bool {
	ret := true
	stack := make([]rune, 0, len(str)/2)
	for _, r := range str {
		v, ok := pairs[r]
		if ok {
			if len(stack) == 0 {
				ret = false
				break
			}
			m := stack[len(stack)-1]
			if v != m {
				ret = false
				break
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}

	}
	return ret && len(stack) == 0
}
