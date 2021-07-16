package main

import "fmt"

func main(){
	var a = []int{5,7,7,8,8,10}
	var target int = 8

	ret := search(a, target)
	fmt.Println(ret)
}

func search(nums []int, target int) int {
	var total int = 0
	for i:=0 ; i < len(nums) ; i++ {
		if (nums[i] == target){
			total++
		}
	}

	return total
}
