package main

import (
	"fmt"
	"strconv"
)

func main() {
	DigitalRoot(195)
	fmt.Println(strconv.FormatInt(14, 9))
}

func DigitalRoot(n int) int {
	v := 0
	for _, sn := range strconv.Itoa(n){
		t, err :=  strconv.Atoi(string(sn))
		if err != nil {
			panic(err)
		}
		v += t
	}
	if v >= 10 {
		v = DigitalRoot(v)
	}
	return v
}
