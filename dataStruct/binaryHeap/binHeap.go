package main

import(
	"math"
	"fmt"
)

type BinHeap struct{
	size int
	queue []int
}

func NewBinHeap() *BinHeap{
	queue := []int{math.MinInt32}
	return &BinHeap{size: 0, queue: queue}
}

func (this *BinHeap) Insert(el int){
	var i int

	this.queue = append(this.queue, el)
	lastIdx := this.size + 1
	
	for i = lastIdx; el < this.queue[i/2]; i = i/2 {
		this.queue[i] = this.queue[i/2]
	}

	this.queue[i] = el
	this.size ++
}

func (this *BinHeap) DeleteMin() int{
	if(this.size == 0){
		return this.queue[0]
	}

	min := this.queue[1]
	
	var i, tidx, leftChild int
	for i=1; i*2 < this.size; {
		leftChild = i*2
		if this.queue[leftChild] > this.queue[leftChild + 1] {
			tidx = leftChild + 1
		}else{
			tidx = leftChild
		}
		
		this.queue[i] = this.queue[tidx]
		i = tidx
	}

	this.queue[i] = this.queue[this.size]
	this.queue = this.queue[0:this.size]
	this.size --

	return min
}

func (this *BinHeap) Print(){
	fmt.Println(this.queue[1:]);
}
