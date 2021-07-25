package main

import("fmt")

func main(){
	a := [][]int{{2,1},{3,4},{3,2}}

	ret := restoreArray(a)

	fmt.Println(ret)
}

//哈希表
//以每个元素为key，值为与它相邻的元素（只会有1个或者两个）
//相邻元素只有一个的key，可以当作开始元素或者结束元素
func restoreArray(adjacentPairs [][]int) []int {
	//记录哈希表， 并找出开始元素
	neighbor := make(map[int][]int)
	result := make([]int, len(adjacentPairs) + 1)
	
	for i:=0; i < len(adjacentPairs); i++ {
		neighbor[adjacentPairs[i][0]] = append(neighbor[adjacentPairs[i][0]], adjacentPairs[i][1])
		neighbor[adjacentPairs[i][1]] = append(neighbor[adjacentPairs[i][1]], adjacentPairs[i][0])
	}

	for k,v := range neighbor {
		if len(v) == 1 {
			result[0] = k
			break
		}
	}

	next := neighbor[result[0]][0]
	j := 1
	for len(neighbor[next]) > 1 {
		result[j] = next

		//choose next
		if neighbor[next][0] == result[j-1] {
			next = neighbor[next][1]
		}else{
			next = neighbor[next][0]
		}

		j++
	}

	result[j] = next
	return result
}
