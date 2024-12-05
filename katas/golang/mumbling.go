package main

import (
	"fmt"
	"strings"
)

/**
* accum("abcd") -> "A-Bb-Ccc-Dddd"
* accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
* accum("cwAt") -> "C-Ww-Aaa-Tttt"
**/
func main(){
	Accum("ZpglnRxqenU")
	//Accum("NyffsGeyylB")
	//Accum("MjtkuBovqrU")
	//Accum("EvidjUnokmM")
	//Accum("HbideVbxncC")
}

func Accum(s string) string {
	out := ""
	for k, v := range s{
		out += fmt.Sprintf("-%s", strings.Title(strings.ToLower(strings.Repeat(string(v), k+1))))
	}
	return out[1:]
}

