package main

import(
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestDivisors(t *testing.T){
    divisors:= Divisors(1)
    assert.Equal(t, []int{1}, divisors, "There should be 1 divisor for 1")

    divisors = Divisors(2)
    assert.Equal(t, []int{1, 2}, divisors, "There should be 2 divisor for 2")

    divisors = Divisors(4)
    assert.Equal(t, []int{1, 2, 4}, divisors, "There should be 3 divisor for 4")
}

func TestAdd(t *testing.T){
    added := Add([]int{1})
    assert.Equal(t, 1, added, "The sum of 1 is 1")

    added = Add([]int{1,2})
    assert.Equal(t, 3, added, "The sum of 1,2 is 3")

    added = Add([]int{1,2,4})
    assert.Equal(t, 7, added, "The sum of 1,2,4 is 7")
}

func TestOutput(t *testing.T){
    o := Run(1, 1)
    assert.Equal(t, "Friendly", o, "1 and 1 are Friendly")

    o = Run(1, 2)
    assert.NotEqual(t, "Friendly", o, "1 and 2 are not Friendly")

    o = Run(6, 28)
    assert.Equal(t, "Friendly", o, "6 and 28 are Friendly")
    
    o = Run(3, 9)
    assert.NotEqual(t, "Friendly", o, "3 and 9 are not Friendly")
}

func TestFindHighAndLow(t *testing.T){
    h,l := FindHighAndLow([]int{0, 0})
    assert.Equal(t,0, h, "I guess it is 0")
    assert.Equal(t,0, l, "I guess it is 0")

    h,l = FindHighAndLow([]int{1, 2, 3, 4, 5})
    assert.Equal(t,5, h, "I guess it is 5")
    assert.Equal(t,1, l, "I guess it is 1")

    h,l = FindHighAndLow([]int{1, 2, -3, 4, 5})
    assert.Equal(t,5, h, "I guess it is 5")
    assert.Equal(t,-3, l, "I guess it is -3")

    h,l = FindHighAndLow([]int{1, 9, 3, 4, -5})
    assert.Equal(t,9, h, "I guess it is 9")
    assert.Equal(t,-5, l, "I guess it is -5")

}
