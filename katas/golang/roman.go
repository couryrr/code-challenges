package main

import (
	"fmt"
)
var values = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main(){
	Decode("MDCLXVI")
	Decode("XC")
	Decode("CM")
	Decode("III")

}

func Decode(roman string) int {
	res := 0
	prev := 0
	// When I first started I was looking at this rtl
	// I tried ltr for some reason
	// When I got stuck instead of trying my intial rtl I asked gippity for 
	// pseudo code... I feel kind of bad about this :(
	for i := len(roman)-1; i >= 0; i-- {
		curr := values[string(roman[i])]
		if curr < prev {
			res -= curr
		} else {
			res += curr
		}
		prev = curr
	}
	fmt.Println(res)
	return -1
}
