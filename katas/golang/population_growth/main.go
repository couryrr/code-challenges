package main

import (
	"fmt"
)

func NbYear(p0 int, percent float64, aug int, p int) int {
	p0 += int(float64(p0)*(percent/100))+aug
	if p0 >= p {
		return 1
	}
	return 1 + NbYear(p0, percent, aug, p)
}


func main(){
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000))
}
