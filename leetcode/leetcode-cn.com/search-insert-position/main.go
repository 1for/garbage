package main

import (
	"fmt"
)

func main() {

	var a = []int{1,3,5,6}
	fmt.Println(searchInsert(a, 2))
}

func searchInsert(nums []int, target int) int {
	var l = 0
	var h = len(nums) - 1

	for l<=h {
		mid := (l+h)/2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			h = mid - 1
		}else {
			l = mid + 1
		}
	}
	return l

}
