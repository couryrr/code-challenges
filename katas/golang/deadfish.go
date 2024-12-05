package main

import "fmt"

func main() {
	reg := 0
	state := []int{}
	commands := map[rune]func(){
		'i': func(){
			reg += 1
		},
		'd': func(){
			reg -= 1
		},
		's': func() {
			reg *= reg
		},
		'o': func() {
			state = append(state, reg)			
		},
	}

	for _,v := range "k"{
		fn, ok := commands[v]
		if ok {
			fn()
		}
	}
	fmt.Println(state)
}
