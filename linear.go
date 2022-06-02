package main

import (
	"fmt"
)

func main() {
	chars := map[int]string{
		1:  "a",
		2:  "b",
		3:  "c",
		4:  "d",
		5:  "e",
		6:  "f",
		7:  "g",
		8:  "h",
		9:  "i",
		10: "j",
		11: "k",
		12: "l",
		13: "m",
		14: "n",
		15: "o",
		16: "p",
		17: "q",
		18: "r",
		19: "s",
		20: "t",
		21: "u",
		22: "v",
		23: "w",
		24: "x",
		25: "y",
		26: "z",
		27: "0",
		28: "1",
		29: "2",
		30: "3",
		31: "4",
		32: "5",
		33: "6",
		34: "7",
		35: "8",
		36: "9",
	}
	fmt.Printf("Enter # of chars: \n")
	var num int
	fmt.Scan(&num)
	first := 1
	second := 1
	third := 1
	fourth := 1
	fifth := 1
	for i := 1; i <= num; i++ {

		fmt.Println(chars[fifth] + chars[fourth] + chars[third] + chars[second] + chars[first])
		if chars[fourth] == "9" {
			fifth++
			fourth = 1
		} else if chars[third] == "9" {
			fourth++
			third = 1
		} else if chars[second] == "9" {
			third++
			second = 1
		} else if chars[first] == "9" {
			second++
			first = 1
		} else {
			first++
		}
	}
}
