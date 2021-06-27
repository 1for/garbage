package main

import("fmt")

func main(){
	deadends := []string{"0201","0101","0102","1212","2002"}
	target := "0202"

	ret := openLock(deadends, target)
	fmt.Println(ret)
}

//广度优先搜索
func openLock(deadends []string, target string) int {
	const start = "0000"
	
	dead := make(map[string]bool)
	for _,v :=  range deadends{
		dead[v] = true
	}

	if start == target {
		return 0
	}
	
	var passed = make(map[string]bool)

	type pair struct {
		str string
		num int
	}
	var list = []pair{{str:start, num:0}}

	for len(list) > 0 {
		current := list[0]
		list = list[1:]
		for _, nextString := range next(current.str){
			if !passed[nextString] && !dead[nextString] {
				list = append(list, pair{str:nextString, num:current.num + 1})

				if nextString == target {
					return current.num + 1
				}

				passed[nextString] = true
			}

		}
	}
	
	return -1
}

func next(current string) []string{
	cur := []byte(current)
	ret := []string{}

	for k,v := range cur{
		cur[k] = v + 1
		if cur[k] > '9' {
			cur[k] = '0'
		}
		ret = append(ret, string(cur))

		cur[k] = v - 1
		if cur[k] < '0' {
			cur[k] = '9'
		}
		ret = append(ret, string(cur))
		cur[k] = v
	}
	return ret
}
