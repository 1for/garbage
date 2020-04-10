package main

import "strings"
import "fmt"

func main(){
	s := " hello  ly,  world  ";
	
	rs := reverseWords(s)
	fmt.Println(rs);
}

func reverseWords(s string) string {
	var rs []string

	rs = strings.Fields(s);

	for i, j:=0, len(rs) - 1; i<j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i];
	}

	fmt.Println(rs);
	return strings.Join(rs, " ");
}
