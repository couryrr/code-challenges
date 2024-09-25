package main

import "fmt"

func main() {}

func Divisors(i int) (res []int) {
    for n := range i {
        adjusted := n+1 
        if i%adjusted == 0 {
            res = append(res, adjusted)
        }
    }

    return res
}

func Add(is []int) (res int){
    res = 0

    for _, i := range is{
        res += i
    }

    return res
}

func Gcf(a, b []int){
    for _, an := range a {
        for _, bn := range b {
            if an == bn {
                fmt.Println(an, bn)
            }
        }
    }
}

func Run(a, b int) (res string){
    res = "Friendly"
    d1 := float32(Add(Divisors(a)))/float32(a)
    d2 := float32(Add(Divisors(b)))/float32(b)
    
    if d1 != d2 {
        res = fmt.Sprintf("%f %f",d1, d2)
    }

    return res
}

func FindHighAndLow(nums []int)(high int, low int){
    high = nums[0]
    low = nums[0]
    for _, i := range nums {
        if i > high {
            high = i
        }

        if i < low {
            low = i
        }
    }
    return high, low
}
