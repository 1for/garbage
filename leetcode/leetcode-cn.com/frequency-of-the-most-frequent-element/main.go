package main

import(
	"fmt"
	"sort"
)

func main(){
	var nums = []int{3,9,6}
	var k int = 2
	ret := maxFrequency(nums, k)
	fmt.Println(ret)
}

//排序 + 二分 https://leetcode-cn.com/problems/frequency-of-the-most-frequent-element/solution/1838-zui-gao-pin-yuan-su-de-pin-shu-shua-ub57/
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	length := len(nums)
	var ret = make([]int, length)
	ret[0] = 1
	
	for r := 2; r < length; r++{
		low := 0
		high := r

		for low <= high {
			mid := (low+high)/2
			need := 0

			for i:= mid; i < r; i++ {
				need = need + (nums[r] - nums[i])
			}

			if k > need {
				high = mid - 1
				ret[r] = r - mid + 1
			}

			if k == need {
				ret[r] = r - mid + 1
				break
			}

			if k < need {
				low = mid + 1
			}
		}
	}

	max := 0
	for i := 0; i < length; i++ {
		if ret[i] > max {
			max = ret[i]
		}
	}

	return max
}
