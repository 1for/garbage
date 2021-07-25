package main

import(
	"fmt"
)

func main(){
	fmt.Println(firstBadVersion(5))
	
}


func isBadVersion(version int) bool{
	if version < 4 {
		return false
	}else{
		return true
	}
}

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func firstBadVersion(n int) int {
	var l = 1
	var h = n

	for l<=h{

		
		mid := (l+h)/2
		fmt.Println(mid)		

		if isBadVersion(mid) {
			h = mid - 1
		}else {
			if isBadVersion(mid + 1) {
				return mid + 1
			}else {
				l = mid + 1
			}
		}
	}

	//所有都是错误版本，返回第一个
	return 1
}

