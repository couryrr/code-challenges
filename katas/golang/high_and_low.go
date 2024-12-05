package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main(){
	v := HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")	
	fmt.Println(v)
}

func HighAndLow(in string) string {
	bs := strings.Split(in, " ") 
	
	min := math.MaxInt
	max := math.MinInt

	for _, s := range bs { 
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		if i > max {
			max = i
		}

		if i < min {
			min = i
		}
	}

	return fmt.Sprintf("%d %d", max, min)
}

