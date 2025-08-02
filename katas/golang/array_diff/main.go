package main

import "fmt"

func main() {
	diffed := ArrayDiff([]int{1, 2, 2, 2, 3}, []int{2})
	fmt.Printf("diffed: %x", diffed)
}

func ArrayDiff(a, b []int) []int {
	mem := make(map[int]int)
	ret := make([]int, 0, len(a))
	for _, v := range a {
		_, ok := mem[v]
		if !ok {
			for _, vv := range b {
				if vv == v {
					mem[v] = v
                    ok = true
                    break
				}
			}
		}
        if !ok {
            ret = append(ret, v)
        }
	}
	return ret
}
