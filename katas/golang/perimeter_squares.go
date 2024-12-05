package main

import "fmt"

func main() {
	fmt.Println(Perimeter(5))	
}

func Perimeter(n int) int {
	
	if n <= 1 {
		return n
	} 

	return Perimeter(n-1) + Perimeter(n - 2)
}
