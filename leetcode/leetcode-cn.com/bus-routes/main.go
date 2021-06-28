package main

import(
	"fmt"
	"math"
)

func main(){
	//	[[7,12],[4,5,15],[6],[15,19],[9,12,13]]

	routes := [][]int{{7,12}, {4,5,15}, {6}, {15,19}, {9,12,13}}
	source := 15
	target := 12

	ret := numBusesToDestination(routes, source, target)
	fmt.Println(ret)
	
}

func numBusesToDestination(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}
	
	// 建立不同巴士线路的换乘关系
	// exchange[fromRouteIdx][toRouteIdx]
	numOfRoutes := len(routes)
	exchange := make([][]bool, numOfRoutes)
	for i := range exchange {
		exchange[i] = make([]bool, numOfRoutes)
	}
	//初始化 exchange 完成

	//记录站点所属的公交线路 rec[site] = [RouteIdx1,RouteIdx2]
	rec := make(map[int][]int)

	//遍历公交线路
	for routeIdx, route := range routes {
		//遍历单个线路里面的站点
		for _,site := range route {
			//遍历站点所属的线路
			for _,routeIdxOfSite := range rec[site] {
				exchange[routeIdx][routeIdxOfSite] = true
				exchange[routeIdxOfSite][routeIdx] = true
			}
			rec[site] = append(rec[site], routeIdx)
		}
	}

	//distance 表示 source 到各个公交线路所搭乘的公交车数量
	//distance[routeIdx] = numOfBus
	distance := make([]int, numOfRoutes)
	for i := range distance {
		distance[i] = -1;
	}
	
	//设置广度优先搜索的起始队列
	queue := []int{}
	for _, routeIdx := range rec[source] {
		distance[routeIdx] = 1
		queue = append(queue, routeIdx)
	}

	for len(queue) > 0 {
		currentRoute := queue[0]
		queue = queue[1:]

		for toRoute, canExchange := range exchange[currentRoute] {
			//可以换乘，并且线路之前未经过
			if canExchange && distance[toRoute] == -1 {
				distance[toRoute] = distance[currentRoute] + 1
				queue = append(queue, toRoute)
			}
		}
	}

	ret := math.MaxInt32

	for _,routeIdx := range rec[target]{
		if distance[routeIdx] != -1 && distance[routeIdx] < ret {
			ret = distance[routeIdx]
		}
	}

	if ret < math.MaxInt32	{
		return ret
	}

	return -1
}
