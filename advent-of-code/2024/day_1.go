package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	url := "https://adventofcode.com/2024/day/1/input" 
	file_name := "day_one.txt"

	input := GetDataForDay(url, file_name)	

	sv1, sv2 := SortInput(input)
	diff := CalcDiff(sv1, sv2)
	fmt.Println(diff)
	sim := CalcSim(sv1, sv2)
	fmt.Println(sim)
}

func SortInput(in []byte) ([]int, []int){
	s := string(in)
	w := strings.Split(s, "\n")	
	sv1 := []int{}
	sv2 := []int{}
	for i := range len(w)-1 {
		v := strings.Split(w[i], "   ")
		i1, err := strconv.Atoi(v[0])
		if err != nil {
		}

		i2, err := strconv.Atoi(v[1])
		if err != nil {
		}

		sv1 = append(sv1, i1)
		sv2 = append(sv2, i2)
	}
	slices.Sort(sv1)
	slices.Sort(sv2)
	return sv1, sv2
}

func CalcDiff(sv1, sv2 []int) int{
	dist := 0
	for i := range len(sv1)-1{
		dist += max(sv1[i], sv2[i]) - min(sv1[i], sv2[i])
	}
	return dist
}

func CalcSim(sv1, sv2 []int) int {
	sv2m := map[int]int{}
	for i := range sv2 {
		sv2m[sv2[i]] += 1
	}
	sim := 0
	for i := range sv1 {
		v := sv1[i]
		sim += v * sv2m[v]
	}	
	return sim
}

