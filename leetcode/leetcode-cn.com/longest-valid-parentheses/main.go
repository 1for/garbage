package main

import (
	"fmt"
)

func main() {
	str := "()()(()";

	c := longestValidParentheses(str);
	fmt.Println(c);
}

/** 
 * 题目描述：https://leetcode-cn.com/problems/longest-valid-parentheses/
 * 动态规划解法
 * 
 * 假设 f(i) 表示以第i 位结尾的最长有效括号字串的长度
 * 例如 ()()(()) , f(4)=4, f(5)=0, f(7)=2, f(8)=8
 * s(i) 表示第i位的值，"("或者")"
 * 如果 s(i) = ")"
 *     - 如果 s(i-1) = "(", 则 f(i) = f(i-2) + 2
 *     - 如果 s(i- f(i-1) - 1) = "(", 则 f(i) = f(i - f(i-1) -2) + f(i-1) + 2
 * 其他情况，f(i) = 0
 * 当 i-1 <= 1 时, f(i-1) = 0
 * 最后取 max(f(i))
 */
func longestValidParentheses(s string) int {
	f := make([]int,  len(s))
	max := 0;

	for i:=1; i<len(s); i++ {
		if string(s[i]) == ")" {
			if string(s[i-1]) == "(" {
				if i < 2 {
					f[i] = 2					
				}else{
					f[i] = f[i-2] + 2
				}
			}

			if (i-f[i-1]-1) >= 0 && string(s[i - f[i-1] -1]) == "(" {
				if (i - f[i-1] -2) >= 0 {
					f[i] = f[i - f[i-1] -2] + f[i - 1] + 2
				}else{
					f[i] = f[i - 1] + 2
				}
			}			
		}

		if f[i] > max {
			max = f[i]
		}
	}

	return max;
}
