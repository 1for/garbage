package main

import "fmt"

func main(){
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	mid := findMedianSortedArrays(nums1, nums2)

	fmt.Printf("%v", mid)
}
