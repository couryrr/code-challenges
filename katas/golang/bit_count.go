package main

import (
	"fmt"
	"math/bits"
)

func main(){
	fmt.Println(CountBits(0))
	fmt.Println(CountBits(4))
	fmt.Println(CountBits(7))

}

func CountBits(n uint) int {
    return bits.OnesCount(n)
}
