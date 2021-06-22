package main

import "fmt"

func main(){
	//var turnedOn = 1

	result := readBinaryWatch(1)

	fmt.Println(result)
}

func readBinaryWatch(turnedOn int) []string {
	result := make([]string,0)
	//hh::mm
	for i:=0 ; i<12 ; i++ {
		for j:=0 ; j<60; j++ {
			if turnedOn == (DecToBinCountOne(i) + DecToBinCountOne(j)){
				fmt.Println(i,j)
				result = append(result, fmt.Sprintf("%d:%02d", i, j))
			}
		}
	}
	return result
}

func DecToBinCountOne(n int) int {
	var c int = 0	
	for {
		r := n%2
		n = n/2
		if r == 1 {
			c++
		}
		if n == 0 {
			return c
		}
	}
}
