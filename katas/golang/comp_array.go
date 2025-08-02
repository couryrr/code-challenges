package main

import (
	"fmt"
	"slices"
)

// https://www.codewars.com/kata/550498447451fbbd7600041c/train/go
func main() {
	a := []int{121, 144, 19, 161, 19, 144, 19, 11}
	b := []int{121, 14641, 20736, 361, 25921, 361, 20736, 361}

	fmt.Println(Comp(a, b))
}

func Comp(a, b []int) bool {
    if a == nil || b == nil {
        return false
    }
    // sort both arrays
    slices.Sort(a)
    slices.Sort(b)
    for i := range a {
        if a[i]*a[i] != b[i]{
            return false
        }
    }
	return true
}
