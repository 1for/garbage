package main

import "fmt";

func main(){
	s := "aab";
	p := "c*a*b";

	fmt.Println(isMatch(s, p));
	
}

/**
 * 方法1:
 * 回溯算法
 * 对比第一个字符是否匹配
 */
func isMatch(s string, p string) bool {
	//如果p为空
	if len(p) == 0 {
		return ( len(s) == 0 );
	}

	first_char_bool := false;
	var snext string;
	
	if len(s) == 0 {
		first_char_bool = false;
		snext = s;
	}else{
		if s[0:1] == p[0:1] || p[0:1] == "." {
			first_char_bool = true;
			snext = s[1:];
		}
	}

	if len(p) >= 2 && p[1:2] == "*" {
		return (first_char_bool && isMatch(snext, p)) ||
			isMatch(s, p[2:]);
	}else {
		return first_char_bool && isMatch(snext, p[1:]);
	}
}
