package main

import (
	"math"
)
/**
 *
 */
func findMedianSortedArrays(nums1 []int,  nums2 []int) float64 {
	total := len(nums1) + len(nums2)
	half := 0
	
	if total % 2 == 1{
		half = total/2 + 1
		return float64(findKth(nums1, 0, nums2, 0, half))
	}else{
		half = total/2
		return (float64(findKth(nums1, 0, nums2, 0, half)) + float64(findKth(nums1, 0, nums2, 0, half + 1)) ) / 2.0
	}
}

/**
 * 找到数组中第K大的数
 * find kth number of two sorted array
 */
func findKth(arr1 []int, start1 int, arr2 []int, start2 int, k int) int {
	if start1 >= len(arr1) {
		return arr2[start2 + k - 1];
	}

	if start2 >= len(arr2) {
		return arr1[start1 + k - 1];
	}

	if k == 1 {
		if( arr1[start1] > arr2[start2] ){
			return arr2[start2]
		}else{
			return arr1[start1]
		}
	}

	key1 := math.MaxInt32
	key2 := math.MaxInt32
	halfK := k/2
	
	if (start1 + halfK - 1) < len(arr1){
		key1 = arr1[start1 + halfK - 1]
	}

	if (start2 + halfK - 1) < len(arr2){
		key2 = arr2[start2 + halfK - 1]
	}

	if key1 > key2 {
		return findKth(arr1, start1, arr2, start2 + halfK, k - halfK)
	}else{
		return findKth(arr1, start1 + halfK, arr2, start2, k - halfK)
	}

	return 0
}
