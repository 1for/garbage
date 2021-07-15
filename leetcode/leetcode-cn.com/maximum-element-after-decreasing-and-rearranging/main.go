package main

import (
	"fmt"
	"sort"
)

func main(){
	var a = []int{100,1,1000}

	ret := maximumElementAfterDecrementingAndRearranging(a)
	fmt.Println(ret)
}

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	sort.Ints(arr)

	arr[0] = 1
	length := len(arr)
	for i:=1; i<length; i++ {
		if arr[i] - arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[length - 1]
}
