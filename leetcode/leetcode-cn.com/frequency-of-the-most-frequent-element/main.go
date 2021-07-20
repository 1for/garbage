package main

import(
	"fmt"
	"sort"
)

func main(){
	var nums = []int{1,2,3,4}
	var k int = 91710
	ret := maxFrequency(nums, k)
	fmt.Println(ret)
}

//排序 + 滑动窗口
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	total := 0
	max := 1

	for left,right := 0,0 ; right < len(nums); right++ {
		total = total + nums[right]
		
		for left < right {
			need := (right - left + 1) * nums[right] - total
			if need > k {
				total = total - nums[left]
				left++
			}else {
				break
			}
		}

		if (right - left + 1) > max {
			max = right - left + 1
		}		
	}
	

	return max
}
