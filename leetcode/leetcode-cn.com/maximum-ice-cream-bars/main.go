package main

import(
	"sort"
	"fmt"
)

func main(){
	/* return 4
	costs := []int {1,3,2,4,1}
	coins := 7
        */
	/* return 0 */
	costs := []int {10,6,8,7,7,8}
	coins := 5
	
	ret := maxIceCream(costs, coins)
	fmt.Println(ret)
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	cnt := len(costs)
	len := 0;
	for i:=0 ; i< cnt ; i++ {
		if coins < costs[i]{
			return len
		}
		coins = coins - costs[i]
		len ++
	}

	return len
}
